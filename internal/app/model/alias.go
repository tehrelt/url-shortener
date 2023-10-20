package model

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Alias struct {
	ID     int    `json:"id"`
	URL    string `json:"url"`
	Alias  string `json:"alias"`
	Visits int    `json:"visits,omitempty"`
}

func (a *Alias) Validate() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.URL, validation.Required, is.URL),
		validation.Field(&a.Alias, validation.Required),
	)
}

func (a *Alias) BeforeCreate() error {
	if len(a.Alias) > 0 {
		a.Alias = strings.ReplaceAll(a.Alias, " ", "-")
	}

	return nil
}
