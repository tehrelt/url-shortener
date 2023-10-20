package model

import "testing"

func TestAlias(t *testing.T) *Alias {
	t.Helper()

	return &Alias{
		URL: "https://github.com/tehrelt/url-shortener",
		Alias: "this-repo",
	}
}
