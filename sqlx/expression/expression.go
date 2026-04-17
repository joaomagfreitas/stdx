package sqlx_expression

import "fmt"

// Expression represents a composable SQL expression.
type Expression interface {
	// And begins a logical AND combination with another expression.
	And() *combineExpression

	// Or begins a logical OR combination with another expression.
	Or() *combineExpression

	// String builds the expression in SQL using [placeholder] for mapping placeholders.
	String(placeholder PlaceholderMapping) string
}

// placeholderExpression represents an expression which the left hand is a column and the
// right hand is a placeholder.
type placeholderExpression struct {
	column   string
	index    int
	operator operator
}

// valueExpression represents an expression which the left hand is a column and the
// right hand is a raw value.
type valueExpression struct {
	column   string
	value    string
	operator operator
}

// combineExpression represents a logical combination of two expressions.
type combineExpression struct {
	left     Expression
	right    Expression
	operator operator
}

// And begins a logical AND with another expression.
func (exp placeholderExpression) And() *combineExpression {
	return &combineExpression{left: exp, operator: and}
}

// Or begins a logical OR with another expression.
func (exp placeholderExpression) Or() *combineExpression {
	return &combineExpression{left: exp, operator: or}
}

// String builds the expression in SQL using [placeholder] for mapping placeholders.
func (exp placeholderExpression) String(placeholder PlaceholderMapping) string {
	return fmt.Sprintf("%s %s %s", exp.column, exp.operator.String(), placeholder(exp.index))
}

// And begins a logical AND with another expression.
func (exp valueExpression) And() *combineExpression {
	return &combineExpression{left: exp, operator: and}
}

// Or begins a logical OR with another expression.
func (exp valueExpression) Or() *combineExpression {
	return &combineExpression{left: exp, operator: or}
}

// String builds the expression in SQL using [placeholder] for mapping placeholders.
func (exp valueExpression) String(_ PlaceholderMapping) string {
	return fmt.Sprintf("%s %s %s", exp.column, exp.operator.String(), exp.value)
}

// And begins a logical AND with another expression.
func (exp combineExpression) And() *combineExpression {
	return &combineExpression{left: exp, operator: and}
}

// Or begins a logical OR with another expression.
func (exp combineExpression) Or() *combineExpression {
	return &combineExpression{left: exp, operator: or}
}

// String builds the expression in SQL using [placeholder] for mapping placeholders.
func (exp combineExpression) String(placeholder PlaceholderMapping) string {
	return fmt.Sprintf("(%s %s %s)", exp.left.String(placeholder), exp.operator.String(), exp.right.String(placeholder))
}

// Equals sets the right-hand side to an equality comparison with a placeholder.
func (exp *combineExpression) Equals(column string, placeholderIndex int) Expression {
	exp.right = Equals(column, placeholderIndex)
	return exp
}

// EqualsTo sets the right-hand side to an equality comparison with a value.
func (exp *combineExpression) EqualsTo(column string, value string) Expression {
	exp.right = EqualsTo(column, value)
	return exp
}

// NotEquals sets the right-hand side to an inequality comparison with a placeholder.
func (exp *combineExpression) NotEquals(column string, placeholderIndex int) Expression {
	exp.right = NotEquals(column, placeholderIndex)
	return exp
}

// NotEqualsTo sets the right-hand side to an inequality comparison with a value.
func (exp *combineExpression) NotEqualsTo(column string, value string) Expression {
	exp.right = NotEqualsTo(column, value)
	return exp
}

// Greater sets the right-hand side to a greater-than comparison with a placeholder.
func (exp *combineExpression) Greater(column string, placeholderIndex int) Expression {
	exp.right = Greater(column, placeholderIndex)
	return exp
}

// GreaterThan sets the right-hand side to a greater-than comparison with a value.
func (exp *combineExpression) GreaterThan(column string, value string) Expression {
	exp.right = GreaterThan(column, value)
	return exp
}

// Less sets the right-hand side to a less-than comparison with a placeholder.
func (exp *combineExpression) Less(column string, placeholderIndex int) Expression {
	exp.right = Less(column, placeholderIndex)
	return exp
}

// LessThan sets the right-hand side to a less-than comparison with a value.
func (exp *combineExpression) LessThan(column string, value string) Expression {
	exp.right = LessThan(column, value)
	return exp
}

// GreaterOrEquals sets the right-hand side to a >= comparison with a placeholder.
func (exp *combineExpression) GreaterOrEquals(column string, placeholderIndex int) Expression {
	exp.right = GreaterOrEquals(column, placeholderIndex)
	return exp
}

// GreaterOrEqualsThan sets the right-hand side to a >= comparison with a value.
func (exp *combineExpression) GreaterOrEqualsThan(column string, value string) Expression {
	exp.right = GreaterOrEqualsThan(column, value)
	return exp
}

// LessOrEquals sets the right-hand side to a <= comparison with a placeholder.
func (exp *combineExpression) LessOrEquals(column string, placeholderIndex int) Expression {
	exp.right = LessOrEquals(column, placeholderIndex)
	return exp
}

// LessOrEqualsThan sets the right-hand side to a <= comparison with a value.
func (exp *combineExpression) LessOrEqualsThan(column string, value string) Expression {
	exp.right = LessOrEqualsThan(column, value)
	return exp
}

// Like sets the right-hand side to a LIKE comparison with a placeholder.
func (exp *combineExpression) Like(column string, placeholderIndex int) Expression {
	exp.right = Like(column, placeholderIndex)
	return exp
}

// LikeWith sets the right-hand side to a LIKE comparison with a value.
func (exp *combineExpression) LikeWith(column string, value string) Expression {
	exp.right = LikeWith(column, value)
	return exp
}

// Equals returns an equality expression using a placeholder.
func Equals(column string, placeholderIndex int) Expression {
	return placeholderExpression{column: column, operator: equals, index: placeholderIndex}
}

// EqualsTo returns an equality expression using a literal value.
func EqualsTo(column string, value string) Expression {
	return valueExpression{column: column, value: value, operator: equals}
}

// NotEquals returns an inequality expression using a placeholder.
func NotEquals(column string, placeholderIndex int) Expression {
	return placeholderExpression{column: column, operator: notEquals, index: placeholderIndex}
}

// NotEqualsTo returns an inequality expression using a literal value.
func NotEqualsTo(column string, value string) Expression {
	return valueExpression{column: column, value: value, operator: notEquals}
}

// Greater returns a greater-than expression using a placeholder.
func Greater(column string, placeholderIndex int) Expression {
	return placeholderExpression{column: column, operator: greater, index: placeholderIndex}
}

// GreaterThan returns a greater-than expression using a literal value.
func GreaterThan(column string, value string) Expression {
	return valueExpression{column: column, value: value, operator: greater}
}

// Less returns a less-than expression using a placeholder.
func Less(column string, placeholderIndex int) Expression {
	return placeholderExpression{column: column, operator: less, index: placeholderIndex}
}

// LessThan returns a less-than expression using a literal value.
func LessThan(column string, value string) Expression {
	return valueExpression{column: column, value: value, operator: less}
}

// GreaterOrEquals returns a >= expression using a placeholder.
func GreaterOrEquals(column string, placeholderIndex int) Expression {
	return placeholderExpression{column: column, operator: greaterOrEquals, index: placeholderIndex}
}

// GreaterOrEqualsThan returns a >= expression using a literal value.
func GreaterOrEqualsThan(column string, value string) Expression {
	return valueExpression{column: column, value: value, operator: greaterOrEquals}
}

// LessOrEquals returns a <= expression using a placeholder.
func LessOrEquals(column string, placeholderIndex int) Expression {
	return placeholderExpression{column: column, operator: lessOrEquals, index: placeholderIndex}
}

// LessOrEqualsThan returns a <= expression using a literal value.
func LessOrEqualsThan(column string, value string) Expression {
	return valueExpression{column: column, value: value, operator: lessOrEquals}
}

// Like returns a LIKE expression using a placeholder.
func Like(column string, placeholderIndex int) Expression {
	return placeholderExpression{column: column, operator: like, index: placeholderIndex}
}

// LikeWith returns a LIKE expression using a literal value.
func LikeWith(column string, value string) Expression {
	return valueExpression{column: column, value: value, operator: like}
}
