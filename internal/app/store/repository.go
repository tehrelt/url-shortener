package store

import "github.com/tehrelt/url-shortener/internal/app/model"

type AliasRepository interface {
	Create(alias *model.Alias) error
	Find(alias string) (*model.Alias, error)
	Delete(alias string) error

	GetAll() ([]model.Alias, error)
}
