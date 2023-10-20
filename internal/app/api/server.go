package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tehrelt/url-shortener/internal/app/api/middleware"
	"github.com/tehrelt/url-shortener/internal/app/model"
	"github.com/tehrelt/url-shortener/internal/app/store"
)

const (
	ENV_LOCAL = "local"
	ENV_DEV   = "dev"
	ENV_PROD  = "prod"
)

type server struct {
	store  store.Store
	logger *slog.Logger
	router *mux.Router
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case ENV_LOCAL:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case ENV_DEV:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case ENV_PROD:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func newServer(store store.Store, env string) *server {
	s := &server{
		store:  store,
		logger: setupLogger(env),
		router: mux.NewRouter(),
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(middleware.CommonMiddleware)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	s.router.HandleFunc("/urls", s.handleAliases()).Methods("GET")
	s.router.HandleFunc("/{alias}", s.handleAlias()).Methods("GET")

	url := s.router.PathPrefix("/alias").Subrouter()
	url.HandleFunc("/", s.handleCreateAlias()).Methods("POST")
	url.HandleFunc("/{alias}", s.handleGetAlias()).Methods("GET")
	url.HandleFunc("/{alias}", s.handleDeleteAlias()).Methods("DELETE")
}

func (s *server) handleAliases() http.HandlerFunc {
	type response struct {
		Count int           `json:"count"`
		Data  []model.Alias `json:"data"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.Alias().GetAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		re := response{
			Data:  data,
			Count: len(data),
		}

		s.respond(w, r, http.StatusOK, re)
	}
}

func (s *server) handleCreateAlias() http.HandlerFunc {
	type request struct {
		Url   string `json:"url"`
		Alias string `json:"alias,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		fmt.Println("creating alias:", req)

		a := &model.Alias{
			Alias: req.Alias,
			URL:   req.Url,
		}

		if err := s.store.Alias().Create(a); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, a)
	}
}

func (s *server) handleGetAlias() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alias := mux.Vars(r)["alias"]

		a, err := s.store.Alias().Find(alias)
		if err != nil {
			s.error(w, r, http.StatusNotFound, ErrorAliasNotFound)
			return
		}

		s.respond(w, r, http.StatusOK, a)
	}
}

func (s *server) handleAlias() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		alias := mux.Vars(r)["alias"]

		a, err := s.store.Alias().Find(alias)
		if err != nil {
			s.error(w, r, http.StatusNotFound, ErrorAliasNotFound)
			return
		}

		http.Redirect(w, r, a.URL, http.StatusSeeOther)
	}
}

func (s *server) handleDeleteAlias() http.HandlerFunc {
	type request struct {
		Url   string `json:"url"`
		Alias string `json:"alias,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		alias := mux.Vars(r)["alias"]

		_, err := s.store.Alias().Find(alias)
		if err != nil {
			s.error(w, r, http.StatusNotFound, ErrorAliasNotFound)
			return
		}

		if err := s.store.Alias().Delete(alias); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, map[string]string{
			"alias": alias,
		})
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
