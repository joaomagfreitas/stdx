package sqlx_expression_test

import (
	"testing"

	sqlx_expression "github.com/joaomagfreitas/stdx/sqlx/expression"
)

func TestPlaceholderMapping(t *testing.T) {
	testCases := []struct {
		desc               string
		placeholderMapping sqlx_expression.PlaceholderMapping
		placeholder        string
	}{
		{
			desc:               "default placeholder includes '?' for parameters",
			placeholderMapping: sqlx_expression.DefaultPlaceholderMapping,
			placeholder:        "?",
		},
		{
			desc:               "postgres placeholder includes '$' followed by col index +1 for parameters",
			placeholderMapping: sqlx_expression.PostgresPlaceholderMapping,
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
