package sqlx_builder

import (
	"fmt"
	"strings"

	sqlx_expression "github.com/joaomagfreitas/stdx/sqlx/expression"
)

type updateBuilder struct {
	placeholderMapping sqlx_expression.PlaceholderMapping
	table              string
	setColumns         []string
	whereColumns       []string
	values             []string
	returning          []string
}

// Update allows to create an UPDATE SQL query in a fluent manner,
// by providing a builder API for each supported clause.
// //
// Example:
//
//	query := Update().
//	    Table("users").
//	    SET("name", "age").
//	    WHERE("id").
//	    String()
//
//	// Output:
//	// UPDATE users
//	// SET name = ?, age = ?
//	// WHERE id = ?
func Update() *updateBuilder {
	return &updateBuilder{
		placeholderMapping: sqlx_expression.DefaultPlaceholderMapping,
	}
}

// Table sets the target table name for the UPDATE statement. The table name is treated "as-is".
func (b *updateBuilder) Table(table string) *updateBuilder {
	b.table = table
	return b
}

// Set sets the column names to update for the UPDATE statement. The column names are treated "as-is".
func (b *updateBuilder) Set(columns ...string) *updateBuilder {
	b.setColumns = columns
	return b
}

// Where sets the column names of the matching rows for the UPDATE statement. The column names are treated "as-is".
func (b *updateBuilder) Where(columns ...string) *updateBuilder {
	b.whereColumns = columns
	return b
}

// Values sets explicit values for the UPDATE statement.
//
// If not provided, placeholders will be generated automatically
// based on the number of columns and the placeholder mapping function.
func (b *updateBuilder) Values(values ...string) *updateBuilder {
	b.values = values
	return b
}

// PlaceholderMapping overrides the default placeholder generation logic.
//
// Useful for supporting different SQL dialects (e.g. PostgreSQL).
//
// Example:
//
//	query := Update().
//	    Table("users").
//	    SET("name", "age").
//	    WHERE("id").
//	    PlaceholderMapping(func(i int) string {
//	        return fmt.Sprintf("$%d", i+1)
//	    })
//	    String()
//
//	// Output:
//	// UPDATE users
//	// SET name = $1, age = $2
//	// WHERE id = $3
func (b *updateBuilder) PlaceholderMapping(placeholderMapping sqlx_expression.PlaceholderMapping) *updateBuilder {
	b.placeholderMapping = placeholderMapping
	return b
}

// Returning sets the column names to return updated values for the UPDATE statement.
// The column names are treated "as-is".
func (b *updateBuilder) Returning(returning ...string) *updateBuilder {
	b.returning = returning
	return b
}

// String builds and returns the final SQL UPDATE query as a string.
func (b *updateBuilder) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "UPDATE %s", b.table)

	lsc := len(b.setColumns)
	lv := len(b.values)
	idx := 0

	set := []string{}
	for i, col := range b.setColumns {
		if i < lv {
			set = append(set, fmt.Sprintf("%s = %s", col, b.values[i]))
		} else {
			set = append(set, fmt.Sprintf("%s = %s", col, b.placeholderMapping(idx)))
			idx++
		}
	}

	fmt.Fprintf(&sb, "\nSET %s", strings.Join(set, ", "))

	where := []string{}
	for i, col := range b.whereColumns {
		j := i + lsc
		if j < lv {
			where = append(where, fmt.Sprintf("%s = %s", col, b.values[j]))
		} else {
			where = append(where, fmt.Sprintf("%s = %s", col, b.placeholderMapping(idx)))
			idx++
		}
	}

	fmt.Fprintf(&sb, "\nWHERE %s", strings.Join(where, "AND "))

	if len(b.returning) > 0 {
		fmt.Fprintf(&sb, "\nRETURNING %s", strings.Join(b.returning, ", "))
	}

	sb.WriteRune(';')

	return sb.String()
}
