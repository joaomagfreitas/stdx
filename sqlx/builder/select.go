package sqlx_builder

import (
	"fmt"
	"strings"

	sqlx_expression "github.com/joaomagfreitas/stdx/sqlx/expression"
)

type selectBuilder struct {
	where              sqlx_expression.Expression
	placeholderMapping sqlx_expression.PlaceholderMapping
	table              string
	sortColumn         string
	columns            []string
	limit              int64
	sortAscending      bool
}

// Select allows to create a SELECT SQL query in a fluent manner,
// by providing a builder API for each supported clause.
// //
// Example:
//
//	query := Select().
//	    Columns("*")
//		From("users").
//		Where(
//			sql_expression.
//				Equals("user_id", "1").
//				Or()
//				Equals("name", 0),
//		).
//		Limit(10).
//	    String()
//
//	// Output:
//	// SELECT * FROM users
//	// WHERE (user_id = 1 OR name = ?)
//	// LIMIT 5;
func Select() *selectBuilder {
	return &selectBuilder{
		placeholderMapping: sqlx_expression.DefaultPlaceholderMapping,
	}
}

// Columns sets the column names to return. The column names are treated "as-is".
func (b *selectBuilder) Columns(columns ...string) *selectBuilder {
	b.columns = columns
	return b
}

// From sets the target table name. The table name is treated "as-is".
func (b *selectBuilder) From(table string) *selectBuilder {
	b.table = table
	return b
}

// Where sets the filter expression.
func (b *selectBuilder) Where(expression sqlx_expression.Expression) *selectBuilder {
	b.where = expression
	return b
}

// Limit sets how many rows to return.
func (b *selectBuilder) Limit(count int64) *selectBuilder {
	b.limit = count
	return b
}

// SortAsc sets query to sort results in ascending order by a column. The column name is treated "as-is".
func (b *selectBuilder) SortAsc(column string) *selectBuilder {
	b.sortAscending = true
	b.sortColumn = column
	return b
}

// SortDesc sets query to sort results in descending order by a column. The column name is treated "as-is".
func (b *selectBuilder) SortDesc(column string) *selectBuilder {
	b.sortAscending = false
	b.sortColumn = column
	return b
}

// PlaceholderMapping overrides the default placeholder generation logic.
//
// Useful for supporting different SQL dialects (e.g. PostgreSQL).
//
// Example:
//
//	query := Select().
//	    Columns("*")
//		From("users").
//		Where(
//			sql_expression.
//				Equals("user_id", "1").
//				Or()
//				Equals("name", 0),
//		).
//		Limit(10).
//		PlaceholderMapping(sqlx_expression.PostgresPlaceholderMapping).
//	    String()
//
//	// Output:
//	// SELECT * FROM users
//	// WHERE (user_id = 1 OR name = $1)
//	// LIMIT 5;
func (b *selectBuilder) PlaceholderMapping(placeholderMapping sqlx_expression.PlaceholderMapping) *selectBuilder {
	b.placeholderMapping = placeholderMapping
	return b
}

// String builds and returns the final SQL SELECT query as a string.
func (b *selectBuilder) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "SELECT %s FROM %s", strings.Join(b.columns, ", "), b.table)

	if b.where != nil {
		fmt.Fprintf(&sb, "\nWHERE %s", b.where.String(b.placeholderMapping))
	}

	if len(b.sortColumn) > 0 {
		sort := "ASC"
		if !b.sortAscending {
			sort = "DESC"
		}

		fmt.Fprintf(&sb, "\nORDER BY %s %s", b.sortColumn, sort)
	}

	if b.limit > 0 {
		fmt.Fprintf(&sb, "\nLIMIT %d", b.limit)
	}

	sb.WriteRune(';')

	return sb.String()
}
