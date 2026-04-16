package sqlx_builder

import "fmt"

/*
WHERE A = $1
WHERE A != 2
WHERE A LIKE 'add'
WHERE A IN (1, 2, 3)
*/

type expr interface {
	And(expression expr)
	Or(expression expr)
	String(placeholder PlaceholderMapping, placeholderIndex int)
}

type placeholderExpression struct {
	column   string
	operator operator
}

func (exp placeholderExpression) String() {}

type valueExpression struct {
	column   string
	operator operator
	value    string
}

type expression struct {
	column           string
	value            string
	operator         operator
	placeholder      PlaceholderMapping
	placeholderIndex int
}

func (exp expression) String() string {
	var op string
	switch exp.operator {
	case equals:
		op = "="
	case notEquals:
		op = "<>"
	case greater:
		op = ">"
	case greaterOrEquals:
		op = ">="
	case less:
		op = "<"
	case lessOrEquals:
		op = "<="
	case like:
		op = "LIKE"
	case and:
		op = "AND"
	case or:
		op = "OR"
	}

	if exp.placeholder != nil {
		return fmt.Sprintf("%s %s %s", exp.column, op, exp.placeholder(exp.placeholderIndex))
	}

	return fmt.Sprintf("%s %s %s", exp.column, op, exp.value)
}

func Equals(column string) expression {
	return expression{column: column, operator: equals, placeholder: DefaultPlaceholderMapping}
}

func EqualsTo(column string, value string) expression {
	return expression{column: column, value: value, operator: equals}
}

func NotEquals(column string) expression {
	return expression{column: column, operator: notEquals, placeholder: DefaultPlaceholderMapping}
}

func NotEqualsTo(column string, value string) expression {
	return expression{column: column, value: value, operator: notEquals}
}

func Greater(column string) expression {
	return expression{column: column, operator: greater, placeholder: DefaultPlaceholderMapping}
}

func GreaterThan(column string, value string) expression {
	return expression{column: column, value: value, operator: greater}
}

func Less(column string) expression {
	return expression{column: column, operator: less, placeholder: DefaultPlaceholderMapping}
}

func LessThan(column string, value string) expression {
	return expression{column: column, value: value, operator: less}
}

func GreaterOrEquals(column string) expression {
	return expression{column: column, operator: greaterOrEquals, placeholder: DefaultPlaceholderMapping}
}

func GreaterOrEqualsThan(column string, value string) expression {
	return expression{column: column, value: value, operator: greaterOrEquals, placeholder: DefaultPlaceholderMapping}
}

func LessOrEquals(column string) expression {
	return expression{column: column, operator: lessOrEquals, placeholder: DefaultPlaceholderMapping}
}

func LessOrEqualsThan(column string, value string) expression {
	return expression{column: column, value: value, operator: lessOrEquals}
}

func Like(column string) expression {
	return expression{column: column, operator: like}
}

func LikeWith(column string, value string) expression {
	return expression{column: column, value: value, operator: like}
}

func And(column string) expression {
	return expression{column: column, operator: and}
}

func AndAs(column string, value string) expression {
	return expression{column: column, value: value, operator: and}
}

func Or(column string) expression {
	return expression{column: column, operator: or}
}

func OrAs(column string, value string) expression {
	return expression{column: column, value: value, operator: or}
}
