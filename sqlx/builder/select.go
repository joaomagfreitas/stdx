package sqlx_builder

import (
	"fmt"
	"strings"

	"github.com/joaomagfreitas/stdx/slicesx"
)

type selectBuilder struct {
	placeholderMapping PlaceholderMapping
	table              string
	columns            []string
	values             []string
	where              []expression
	limit              int64
	sortAscending      bool
	sortColumn         string
}

/*
SELECT * FROM ... WHERE ... OR.. AND ... LIMIT
*/
// Insert allows to create an INSERT SQL query in a fluent manner,
// by providing a builder API for each supported clause.
// //
// Example:
//
//	query := Insert().
//	    Into("users").
//	    Columns("name", "age").
//	    String()
//
//	// Output:
//	// INSERT INTO users (name, age)
//	// VALUES (?, ?);
func Select() *selectBuilder {
	return &selectBuilder{
		placeholderMapping: DefaultPlaceholderMapping,
	}
}

// Columns sets the column names for the INSERT statement. The column names are treated "as-is".
func (b *selectBuilder) Columns(columns ...string) *selectBuilder {
	b.columns = columns
	return b
}

// From sets the target table name for the SELECT statement. The table name is treated "as-is".
func (b *selectBuilder) From(table string) *selectBuilder {
	b.table = table
	return b
}

// Where sets the filter expressions for the SELECT statement.
func (b *selectBuilder) Where(expressions ...expression) *selectBuilder {
	b.where = expressions
	return b
}

// Limit sets how many rows to return for the SELECT statement.
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
//	query := Insert().
//	    Into("users").
//	    Columns("name", "age").
//	    PlaceholderMapping(func(i int) string {
//	        return fmt.Sprintf("$%d", i+1)
//	    }).
//	    String()
//
//	// Output:
//	// INSERT INTO users (name, age)
//	// VALUES ($1, $2);
func (b *selectBuilder) PlaceholderMapping(placeholderMapping PlaceholderMapping) *selectBuilder {
	b.placeholderMapping = placeholderMapping
	return b
}

// String builds and returns the final SQL INSERT query as a string.
func (b *selectBuilder) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "INSERT INTO %s", b.table)

	if len(b.columns) > 0 {
		fmt.Fprintf(&sb, " (%s)", strings.Join(b.columns, ", "))
	}

	var vs []string
	if len(b.values) > 0 {
		vs = b.values
	} else {
		vs = slicesx.Gen(len(b.columns), b.placeholderMapping)
	}
	fmt.Fprintf(&sb, "\nVALUES (%s)", strings.Join(vs, ", "))

	if len(b.returning) > 0 {
		fmt.Fprintf(&sb, "\nRETURNING %s", strings.Join(b.returning, ", "))
	}

	sb.WriteRune(';')

	return sb.String()
}
