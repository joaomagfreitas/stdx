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
