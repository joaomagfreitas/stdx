package sqlx_builder

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

type operator byte

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
