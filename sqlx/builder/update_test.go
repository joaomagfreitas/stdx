package sqlx_builder_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	sqlx_builder "github.com/joaomagfreitas/stdx/sqlx/builder"
)

func TestUpdate(t *testing.T) {
	testCases := []struct {
		builder func() string
		desc    string
		sql     string
	}{
		{
			desc: "joins values following the order they were passed",
			builder: func() string {
				return sqlx_builder.
					Update().
					Table("foo").
					Set("bar", "lorem", "ipsum").
					Where("bar").
					Values("1", "2", "3", "4").
					String()
			},
			sql: `
			UPDATE foo
			SET bar = 1, lorem = 2, ipsum = 3
			WHERE bar = 4;`,
		},
		{
			desc: "uses placeholder parameters if values are not passed",
			builder: func() string {
				return sqlx_builder.
					Update().
					Table("foo").
					Set("bar", "lorem", "ipsum").
					Where("bar").
					String()
			},
			sql: `
			UPDATE foo
			SET bar = ?, lorem = ?, ipsum = ?
			WHERE bar = ?;`,
		},
		{
			desc: "fills missing values using placeholder parameters",
			builder: func() string {
				return sqlx_builder.
					Update().
					Table("foo").
					Set("bar", "lorem", "ipsum").
					Where("bar").
					Values("1", "2").
					String()
			},
			sql: `
			UPDATE foo
			SET bar = 1, lorem = 2, ipsum = ?
			WHERE bar = ?;`,
		},
		{
			desc: "index is updated to match column number, when filling missing values using placeholder parameters",
			builder: func() string {
				return sqlx_builder.
					Update().
					Table("foo").
					Set("bar", "lorem", "ipsum").
					Where("bar").
					Values("1", "2").
					PlaceholderMapping(func(index int) string { return fmt.Sprintf("$%d", index+1) }).
					String()
			},
			sql: `
			UPDATE foo
			SET bar = 1, lorem = 2, ipsum = $1
			WHERE bar = $2;`,
		},
		{
			desc: "includes returning columns if passed",
			builder: func() string {
				return sqlx_builder.
					Update().
					Table("foo").
					Set("bar", "lorem", "ipsum").
					Where("bar").
					Values("1", "2").
					Returning("lorem", "ipsum").
					PlaceholderMapping(func(index int) string { return fmt.Sprintf("$%d", index+1) }).
					String()
			},
			sql: `
			UPDATE foo
			SET bar = 1, lorem = 2, ipsum = $1
			WHERE bar = $2
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
