package dbstore

import (
	"database/sql"

	"github.com/tehrelt/url-shortener/internal/app/store"

)

type Store struct {
	db              *sql.DB
	aliasRepository *AliasRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Alias() store.AliasRepository {
	if s.aliasRepository != nil {
		return s.aliasRepository
	}

	s.aliasRepository = &AliasRepository{
		store: s,
	}

	return s.aliasRepository
}
