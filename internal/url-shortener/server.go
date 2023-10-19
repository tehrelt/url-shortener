package urlshortener

import (
	"log/slog"
	"os"
)

const (
	ENV_LOCAL = "local"
	ENV_DEV   = "dev"
	ENV_PROD  = "prod"
)

type Server struct {
	Config *Config
	Logger slog.Logger
}

func New(config *Config) *Server {
	return &Server{
		Config: config,
		Logger: *setupLogger(config.Env),
	}
}

func (s *Server) Start() error {

	s.Logger.Info("starting server", slog.String("PORT", s.Config.Port))

	return nil // http.ListenAndServe("localhost:"+s.config.Port)
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
