package sqlx_expression_test

import (
	"testing"

	sqlx_expression "github.com/joaomagfreitas/stdx/sqlx/expression"
)

func TestExpression(t *testing.T) {
	testCases := []struct {
		desc string
		expr sqlx_expression.Expression
		sql  string
	}{
		{
			desc: "equals",
			expr: sqlx_expression.Equals("foo", 0),
			sql:  "foo = ?",
		},
		{
			desc: "equals to",
			expr: sqlx_expression.EqualsTo("foo", "0"),
			sql:  "foo = 0",
		},
		{
			desc: "not equals",
			expr: sqlx_expression.NotEquals("foo", 0),
			sql:  "foo <> ?",
		},
		{
			desc: "not equals to",
			expr: sqlx_expression.NotEqualsTo("foo", "0"),
			sql:  "foo <> 0",
		},
		{
			desc: "greater",
			expr: sqlx_expression.Greater("foo", 0),
			sql:  "foo > ?",
		},
		{
			desc: "greater than",
			expr: sqlx_expression.GreaterThan("foo", "0"),
			sql:  "foo > 0",
		},
		{
			desc: "greater or equals",
			expr: sqlx_expression.GreaterOrEquals("foo", 0),
			sql:  "foo >= ?",
		},
		{
			desc: "greater or equals than",
			expr: sqlx_expression.GreaterOrEqualsThan("foo", "0"),
			sql:  "foo >= 0",
		},
		{
			desc: "less",
			expr: sqlx_expression.Less("foo", 0),
			sql:  "foo < ?",
		},
		{
			desc: "greater than",
			expr: sqlx_expression.LessThan("foo", "0"),
			sql:  "foo < 0",
		},
		{
			desc: "less or equals",
			expr: sqlx_expression.LessOrEquals("foo", 0),
			sql:  "foo <= ?",
		},
		{
			desc: "less or equals than",
			expr: sqlx_expression.LessOrEqualsThan("foo", "0"),
			sql:  "foo <= 0",
		},
		{
			desc: "like",
			expr: sqlx_expression.Like("foo", 0),
			sql:  "foo LIKE ?",
		},
		{
			desc: "like with",
			expr: sqlx_expression.LikeWith("foo", "0"),
			sql:  "foo LIKE 0",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if sql := tC.expr.String(sqlx_expression.DefaultPlaceholderMapping); sql != tC.sql {
				t.Fatalf("\nexp: %s\ngot: %s\n", tC.sql, sql)
			}
		})
	}
}

func TestCombineExpression(t *testing.T) {
	testCases := []struct {
		desc string
		expr sqlx_expression.Expression
		sql  string
	}{
		{
			desc: "equals",
			expr: sqlx_expression.Equals("foo", 0).And().Equals("bar", 1),
			sql:  "(foo = ? AND bar = ?)",
		},
		{
			desc: "equals to",
			expr: sqlx_expression.EqualsTo("foo", "0").And().EqualsTo("bar", "1"),
			sql:  "(foo = 0 AND bar = 1)",
		},
		{
			desc: "not equals",
			expr: sqlx_expression.NotEquals("foo", 0).And().NotEquals("bar", 0),
			sql:  "(foo <> ? AND bar <> ?)",
		},
		{
			desc: "not equals to",
			expr: sqlx_expression.NotEqualsTo("foo", "0").And().NotEqualsTo("bar", "1"),
			sql:  "(foo <> 0 AND bar <> 1)",
		},
		{
			desc: "greater",
			expr: sqlx_expression.Greater("foo", 0).And().Greater("bar", 1),
			sql:  "(foo > ? AND bar > ?)",
		},
		{
			desc: "greater than",
			expr: sqlx_expression.GreaterThan("foo", "0").And().GreaterThan("bar", "1"),
			sql:  "(foo > 0 AND bar > 1)",
		},
		{
			desc: "greater or equals",
			expr: sqlx_expression.GreaterOrEquals("foo", 0).And().GreaterOrEquals("bar", 1),
			sql:  "(foo >= ? AND bar >= ?)",
		},
		{
			desc: "greater or equals than",
			expr: sqlx_expression.GreaterOrEqualsThan("foo", "0").And().GreaterOrEqualsThan("bar", "1"),
			sql:  "(foo >= 0 AND bar >= 1)",
		},
		{
			desc: "less",
			expr: sqlx_expression.Less("foo", 0).And().Less("bar", 1),
			sql:  "(foo < ? AND bar < ?)",
		},
		{
			desc: "greater than",
			expr: sqlx_expression.LessThan("foo", "0").And().LessThan("bar", "1"),
			sql:  "(foo < 0 AND bar < 1)",
		},
		{
			desc: "less or equals",
			expr: sqlx_expression.LessOrEquals("foo", 0).And().LessOrEquals("bar", 1),
			sql:  "(foo <= ? AND bar <= ?)",
		},
		{
			desc: "less or equals than",
			expr: sqlx_expression.LessOrEqualsThan("foo", "0").And().LessOrEqualsThan("bar", "1"),
			sql:  "(foo <= 0 AND bar <= 1)",
		},
		{
			desc: "like",
			expr: sqlx_expression.Like("foo", 0).And().Like("bar", 1),
			sql:  "(foo LIKE ? AND bar LIKE ?)",
		},
		{
			desc: "like with",
			expr: sqlx_expression.LikeWith("foo", "0").And().LikeWith("bar", "1"),
			sql:  "(foo LIKE 0 AND bar LIKE 1)",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if sql := tC.expr.String(sqlx_expression.DefaultPlaceholderMapping); sql != tC.sql {
				t.Fatalf("\nexp: %s\ngot: %s\n", tC.sql, sql)
			}
		})
	}
}

func TestCombineMultipleExpression(t *testing.T) {
	sql := `((foo = 1 OR foo = 2) OR foo > ?)`
	expr := sqlx_expression.EqualsTo("foo", "1").Or().EqualsTo("foo", "2").Or().Greater("foo", 0)

	assertExpression(t, expr, sql)
}

func assertExpression(t *testing.T, expr sqlx_expression.Expression, sql string) {
	if exprSql := expr.String(sqlx_expression.DefaultPlaceholderMapping); exprSql != sql {
		t.Fatalf("\nexp: %s\ngot: %s\n", exprSql, sql)
	}
}
