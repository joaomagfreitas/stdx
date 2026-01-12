package errorsx

import (
	"errors"
	"fmt"
	"strings"
)

// StructuredError represents an error with structured context.
// It can wrap an underlying error while adding origin, category,
// operation, and a human-readable message.
type StructuredError struct {
	Wrapped   error
	Origin    string
	Category  string
	Operation string
	Message   string
}

func (err StructuredError) Error() string {
	var sb strings.Builder
	args := []any{}

	if len(err.Origin) > 0 {
		sb.WriteString("origin: %s, ")
		args = append(args, err.Origin)
	}

	if len(err.Category) > 0 {
		sb.WriteString("category: %s, ")
		args = append(args, err.Category)
	}

	if len(err.Operation) > 0 {
		sb.WriteString("operation: %s, ")
		args = append(args, err.Operation)
	}

	if len(err.Message) > 0 {
		sb.WriteString("message: %s, ")
		args = append(args, err.Message)
	}

	if err.Wrapped != nil {
		sb.WriteString("err: %v")
		args = append(args, err.Wrapped)
	}

	return fmt.Sprintf(strings.TrimSuffix(sb.String(), ", "), args...)
}

// Wrap wraps an error in a [StructuredError].
// The original error can still be accessed using `errors.As`.
func Wrap(
	err error,
	origin string,
	category string,
	operation string,
	message string,
) error {
	return errors.Join(
		StructuredError{
			Wrapped:   err,
			Origin:    origin,
			Category:  category,
			Operation: operation,
			Message:   message,
		},
		err,
	)
}
