package dbstore

import "github.com/tehrelt/url-shortener/internal/app/model"

type AliasRepository struct {
	store *Store
}

func (r *AliasRepository) Create(alias *model.Alias) error {

	if err := alias.Validate(); err != nil {
		return err
	}

	if err := alias.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO aliases (alias, url) VALUES ($1, $2) RETURNING id",
		alias.Alias,
		alias.URL,
	).Scan(&alias.ID)
}

func (r *AliasRepository) Find(alias string) (*model.Alias, error) {
	a := &model.Alias{}

	if err := r.store.db.QueryRow(
		"SELECT id, alias, url FROM aliases WHERE alias = $1",
		alias,
	).Scan(&a.ID, &a.Alias, &a.URL); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *AliasRepository) Delete(alias string) error {
	return r.store.db.QueryRow(
		"DELETE FROM aliases WHERE alias = $1",
		alias,
	).Err()
}