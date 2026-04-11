package sqlx_builder_test

import (
	"testing"

	sqlx_builder "github.com/joaomagfreitas/stdx/sqlx/builder"
)

func TestPlaceholderMapping(t *testing.T) {
	testCases := []struct {
		desc               string
		placeholderMapping sqlx_builder.PlaceholderMapping
		placeholder        string
	}{
		{
			desc:               "default placeholder includes '?' for parameters",
			placeholderMapping: sqlx_builder.DefaultPlaceholderMapping,
			placeholder:        "?",
		},
		{
			desc:               "postgres placeholder includes '$' followed by col index +1 for parameters",
			placeholderMapping: sqlx_builder.PostgresPlaceholderMapping,
			placeholder:        "$1",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			p := tC.placeholderMapping(0)
			if p != tC.placeholder {
				t.Fatal(p)
			}
		})
	}
}
