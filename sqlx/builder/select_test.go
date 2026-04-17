package sqlx_builder_test

import (
	"slices"
	"strings"
	"testing"

	sqlx_builder "github.com/joaomagfreitas/stdx/sqlx/builder"
	sqlx_expression "github.com/joaomagfreitas/stdx/sqlx/expression"
)

func TestSelect(t *testing.T) {
	testCases := []struct {
		builder func() string
		desc    string
		sql     string
	}{
		{
			desc: "sets rows limit if value is passed",
			builder: func() string {
				return sqlx_builder.
					Select().
					Columns("*").
					From("foo").
					Limit(5).
					String()
			},
			sql: `
			SELECT * FROM foo
			LIMIT 5;
			`,
		},
		{
			desc: "sets sort order (asc) if requested",
			builder: func() string {
				return sqlx_builder.
					Select().
					Columns("*").
					From("foo").
					SortAsc("bar").
					String()
			},
			sql: `
			SELECT * FROM foo
			ORDER BY bar ASC;
			`,
		},
		{
			desc: "sets sort order (desc) if requested",
			builder: func() string {
				return sqlx_builder.
					Select().
					Columns("*").
					From("foo").
					SortDesc("bar").
					String()
			},
			sql: `
			SELECT * FROM foo
			ORDER BY bar DESC;
			`,
		},
		{
			desc: "sets filter expression if provided",
			builder: func() string {
				return sqlx_builder.
					Select().
					Columns("*").
					From("foo").
					Where(
						sqlx_expression.
							Like("bar", 0).
							Or().
							LikeWith("bar", "'lorem'"),
					).
					PlaceholderMapping(sqlx_expression.PostgresPlaceholderMapping).
					String()
			},
			sql: `
			SELECT * FROM foo
			WHERE (bar LIKE $1 OR bar LIKE 'lorem');
			`,
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
