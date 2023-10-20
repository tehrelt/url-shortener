package api

import (
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tehrelt/url-shortener/internal/app/store/dbstore"
)

func Start(config *Config) error {

	db, err := setupDataBase(config.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	store := dbstore.New(db)
	srv := newServer(store, config.Env)
	return http.ListenAndServe(config.Port, srv)
}

func setupDataBase(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS aliases (
			id integer primary key autoincrement,
			alias text unique not null,
			url text not null,
			visits integer
		);
		CREATE INDEX IF NOT EXISTS idx_alias ON aliases(alias);`,
	)
	if err != nil {
		return nil, err
	}

	if _, err := stmt.Exec(); err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
