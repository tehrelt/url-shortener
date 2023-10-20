package dbstore

import (
	"fmt"

	"github.com/tehrelt/url-shortener/internal/app/model"
)

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
		"INSERT INTO aliases (alias, url, visits) VALUES ($1, $2, $3) RETURNING id",
		alias.Alias,
		alias.URL,
		0,
	).Scan(&alias.ID)
}

func (r *AliasRepository) Find(alias string) (*model.Alias, error) {
	a := &model.Alias{}

	if err := r.store.db.QueryRow(
		"SELECT id, alias, url, visits FROM aliases WHERE alias = $1",
		alias,
	).Scan(&a.ID, &a.Alias, &a.URL, &a.Visits); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *AliasRepository) Delete(alias string) error {

	fmt.Printf("delete an alias: %s", alias)

	stmt, err := r.store.db.Prepare("DELETE FROM aliases WHERE alias = ?")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(alias); err != nil {
		return err
	}

	return nil
}

func (r *AliasRepository) GetAll() ([]model.Alias, error) {
	var aliases []model.Alias

	rows, err := r.store.db.Query("SELECT id, alias, url, visits FROM aliases")
	if err != nil {
		return nil, err
	}

	var t model.Alias
	for rows.Next() {
		rows.Scan(&t.ID, &t.Alias, &t.URL, &t.Visits)

		aliases = append(aliases, model.Alias{
			ID:     t.ID,
			Alias:  t.Alias,
			URL:    t.URL,
			Visits: t.Visits,
		})
	}

	return aliases, nil
}

func (r *AliasRepository) IncrementVisit(alias string) error {
	a, err := r.Find(alias)
	if err != nil {
		return err
	}

	stmt, err := r.store.db.Prepare(`
		UPDATE aliases 
		SET
			visits = ?
		WHERE alias = ?
	`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(a.Visits+1, a.Alias); err != nil {
		return err
	}

	return nil
}
