package sqlx_builder_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	sqlx_builder "github.com/joaomagfreitas/stdx/sqlx/builder"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		builder func() string
		desc    string
		sql     string
	}{
		{
			desc: "joins values following the order they were passed",
			builder: func() string {
				return sqlx_builder.
					Insert().
					Into("foo").
					Values("1", "2", "3").
					String()
			},
			sql: `
			INSERT INTO foo
			VALUES (1, 2, 3);`,
		},
		{
			desc: "joins placeholder parameters for each column, if values are not passed",
			builder: func() string {
				return sqlx_builder.
					Insert().
					Into("foo").
					Columns("bar", "lorem", "ipsum").
					String()
			},
			sql: `
			INSERT INTO foo (bar, lorem, ipsum)
			VALUES (?, ?, ?);`,
		},
		{
			desc: "uses placeholder mapping function to map parameters for each column, if values are not passed",
			builder: func() string {
				return sqlx_builder.
					Insert().
					Into("foo").
					Columns("bar", "lorem", "ipsum").
					PlaceholderMapping(func(index int) string { return fmt.Sprintf("$%d", index+1) }).
					String()
			},
			sql: `
			INSERT INTO foo (bar, lorem, ipsum)
			VALUES ($1, $2, $3);`,
		},
		{
			desc: "does not join placeholder parameters for each column, if values are passed",
			builder: func() string {
				return sqlx_builder.
					Insert().
					Into("foo").
					Columns("bar", "lorem", "ipsum").
					Values("1", "2", "3").
					String()
			},
			sql: `
			INSERT INTO foo (bar, lorem, ipsum)
			VALUES (1, 2, 3);`,
		},
		{
			desc: "includes passed column names in returning statement",
			builder: func() string {
				return sqlx_builder.
					Insert().
					Into("foo").
					Columns("bar", "lorem", "ipsum").
					Values("1", "2", "3").
					Returning("lorem", "ipsum").
					String()
			},
			sql: `
			INSERT INTO foo (bar, lorem, ipsum)
			VALUES (1, 2, 3)
			RETURNING lorem, ipsum;`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tkb := strings.Fields(tC.builder())
			tke := strings.Fields(tC.sql)

			if !slices.Equal(tkb, tke) {
				t.Fatalf("\nexp: %v\ngot: %v", tke, tkb)
			}
		})
	}
}
