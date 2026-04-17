package sqlx_expression

const (
	equals operator = iota
	notEquals
	greater
	less
	greaterOrEquals
	lessOrEquals
	like
	and
	or
)

// operator represents a SQL operator used in expressions.
type operator byte

// String returns the SQL string representation of the operator.
func (op operator) String() string {
	switch op {
	case equals:
		return "="
	case notEquals:
		return "<>"
	case greater:
		return ">"
	case greaterOrEquals:
		return ">="
	case less:
		return "<"
	case lessOrEquals:
		return "<="
	case like:
		return "LIKE"
	case and:
		return "AND"
	case or:
		return "OR"
	default:
		return "??"
	}
}
