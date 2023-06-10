package lox

import (
	"fmt"
	"strings"
)

type Expression interface {
	isExpression() bool
	GetLocation() Location
	String() string
}

type ExpressionLiteral struct {
	Location     // not requiring a Token, as the expression can be generated
	LiteralValue any
}

func (*ExpressionLiteral) isExpression() bool         { return true }
func (expr *ExpressionLiteral) GetLocation() Location { return expr.Location }
func (expr *ExpressionLiteral) String() string {
	switch val := expr.LiteralValue.(type) {
	case string:
		return fmt.Sprintf("(literal \"%s\")",
			val,
		)
	default:
		return fmt.Sprintf("(literal %v)",
			expr.LiteralValue,
		)
	}
}

type ExpressionGrouping struct {
	LeftParenToken Token
	Expression     Expression
}

func (*ExpressionGrouping) isExpression() bool         { return true }
func (expr *ExpressionGrouping) GetLocation() Location { return expr.LeftParenToken.Location }
func (expr *ExpressionGrouping) String() string {
	return fmt.Sprintf("(group %s)",
		expr.Expression,
	)
}

type ExpressionVariable struct {
	Identifier Token
}

func (*ExpressionVariable) isExpression() bool         { return true }
func (expr *ExpressionVariable) GetLocation() Location { return expr.Identifier.Location }
func (expr *ExpressionVariable) String() string {
	return fmt.Sprintf("(getVar %s)",
		expr.Identifier.Lexeme,
	)
}

type ExpressionCall struct {
	Callee     Expression
	RightParen Token
	Arguments  []Expression
}

func (*ExpressionCall) isExpression() bool         { return true }
func (expr *ExpressionCall) GetLocation() Location { return expr.Callee.GetLocation() }
func (expr *ExpressionCall) String() string {
	var builder strings.Builder
	for i, arg := range expr.Arguments {
		if i != 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(arg.String())
	}
	return fmt.Sprintf("(call %s [%s])",
		expr.Callee, builder.String(),
	)
}

type ExpressionGet struct {
	Object     Expression
	Identifier Token
}

func (*ExpressionGet) isExpression() bool         { return true }
func (expr *ExpressionGet) GetLocation() Location { return expr.Object.GetLocation() }
func (expr *ExpressionGet) String() string {
	return fmt.Sprintf("(getProp %s.%s)",
		expr.Object, expr.Identifier.Lexeme,
	)
}

type ExpressionSet struct {
	Object     Expression
	Identifier Token
	Value      Expression
}

func (*ExpressionSet) isExpression() bool         { return true }
func (expr *ExpressionSet) GetLocation() Location { return expr.Object.GetLocation() }
func (expr *ExpressionSet) String() string {
	return fmt.Sprintf("(setProp %s.%s %s)",
		expr.Object, expr.Identifier.Lexeme, expr.Value,
	)
}

type ExpressionThis struct {
	ThisToken Token
}

func (*ExpressionThis) isExpression() bool         { return true }
func (expr *ExpressionThis) GetLocation() Location { return expr.ThisToken.Location }
func (expr *ExpressionThis) String() string {
	return "(this)"
}

type ExpressionSuper struct {
	SuperToken Token
	Method     Token
}

func (*ExpressionSuper) isExpression() bool         { return true }
func (expr *ExpressionSuper) GetLocation() Location { return expr.SuperToken.Location }
func (expr *ExpressionSuper) String() string {
	return fmt.Sprintf("(super.%s)",
		expr.Method.Lexeme,
	)
}

type ExpressionUnary struct {
	Operator Token
	Right    Expression
}

func (*ExpressionUnary) isExpression() bool         { return true }
func (expr *ExpressionUnary) GetLocation() Location { return expr.Operator.Location }
func (expr *ExpressionUnary) String() string {
	return fmt.Sprintf("(%s %s)",
		expr.Operator.Lexeme, expr.Right,
	)
}

type ExpressionBinary struct {
	Left     Expression
	Operator Token
	Right    Expression
}

func (*ExpressionBinary) isExpression() bool         { return true }
func (expr *ExpressionBinary) GetLocation() Location { return expr.Left.GetLocation() }
func (expr *ExpressionBinary) String() string {
	return fmt.Sprintf("(%s %s %s)",
		expr.Operator.Lexeme, expr.Left, expr.Right,
	)
}

type ExpressionLogical struct {
	Left     Expression
	Operator Token
	Right    Expression
}

func (*ExpressionLogical) isExpression() bool         { return true }
func (expr *ExpressionLogical) GetLocation() Location { return expr.Left.GetLocation() }
func (expr *ExpressionLogical) String() string {
	return fmt.Sprintf("(%s %s %s)",
		expr.Operator.Lexeme, expr.Left, expr.Right,
	)
}

type ExpressionAssignment struct {
	Identifier Token
	Value      Expression
}

func (*ExpressionAssignment) isExpression() bool         { return true }
func (expr *ExpressionAssignment) GetLocation() Location { return expr.Identifier.Location }
func (expr *ExpressionAssignment) String() string {
	return fmt.Sprintf("(assign %s %s)",
		expr.Identifier.Lexeme, expr.Value,
	)
}
