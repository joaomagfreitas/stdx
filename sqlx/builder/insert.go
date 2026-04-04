package sqlx_builder

import (
	"fmt"
	"strings"

	"github.com/joaomagfreitas/stdx/slicesx"
)

type insertBuilder struct {
	placeholderMapping PlaceholderMapping
	table              string
	columns            []string
	values             []string
	returning          []string
}

// PlaceholderMapping defines a function that maps a parameter index
// (starting at 0) to a SQL placeholder string.
//
// It can be used to adapt the builder to different SQL dialects.
//
// Example:
//
//	// PostgreSQL-style placeholders: $1, $2, ...
//	mapping := func(i int) string {
//	    return fmt.Sprintf("$%d", i+1)
//	}
type PlaceholderMapping func(index int) string

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
func Insert() *insertBuilder {
	return &insertBuilder{
		placeholderMapping: func(index int) string { return "?" },
	}
}

// Into sets the target table name for the INSERT statement. The table name is treated "as-is".
func (b *insertBuilder) Into(table string) *insertBuilder {
	b.table = table
	return b
}

// Columns sets the column names for the INSERT statement. The column names are treated "as-is".
func (b *insertBuilder) Columns(columns ...string) *insertBuilder {
	b.columns = columns
	return b
}

// Values sets explicit values for the INSERT statement.
//
// If not provided, placeholders will be generated automatically
// based on the number of columns and the placeholder mapping function.
func (b *insertBuilder) Values(values ...string) *insertBuilder {
	b.values = values
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
func (b *insertBuilder) PlaceholderMapping(placeholderMapping PlaceholderMapping) *insertBuilder {
	b.placeholderMapping = placeholderMapping
	return b
}

// Returning sets the column names to return inserted values for the INSERT statement.
// The column names are treated "as-is".
func (b *insertBuilder) Returning(returning ...string) *insertBuilder {
	b.returning = returning
	return b
}

// String builds and returns the final SQL INSERT query as a string.
func (b *insertBuilder) String() string {
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
		fmt.Fprintf(&sb, "\nRETURNING (%s)", strings.Join(b.returning, ", "))
	}

	sb.WriteRune(';')

	return sb.String()
}
