package sqlx_builder

import "fmt"

var (
	// DefaultPlaceholderMapping returns the standard "?" placeholder,
	// commonly used by databases like MySQL and SQLite.
	DefaultPlaceholderMapping PlaceholderMapping = func(index int) string {
		return "?"
	}

	// PostgresPlaceholderMapping returns PostgreSQL-style placeholders
	// in the form "$1", "$2", ... Based on the given parameter index.
	//
	// The index is zero-based, so 0 maps to "$1".
	PostgresPlaceholderMapping PlaceholderMapping = func(index int) string {
		return fmt.Sprintf("$%d", index+1)
	}
)

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
