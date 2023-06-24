package golox

import (
	"strings"
)

const (
	STMT_INDENT = "    "
)

type Statement interface {
	implStatement()
	GetLocation() Location
	String() string // does not append newline for all statement types
}

func (*StatementBlock) implStatement()      {}
func (*StatementExpression) implStatement() {}
func (*StatementVar) implStatement()        {}
func (*StatementIf) implStatement()         {}
func (*StatementWhile) implStatement()      {}
func (*StatementFun) implStatement()        {}
func (*StatementReturn) implStatement()     {}
func (*StatementClass) implStatement()      {}
func (*StatementPrint) implStatement()      {}

type StatementBlock struct {
	Location   // not requiring a Token, as the statement can be generated
	Statements []Statement
}

func (stmt *StatementBlock) GetLocation() Location {
	return stmt.Location
}

func (stmt *StatementBlock) String() string {
	var b strings.Builder
	b.WriteString("{\n")
	for _, stmt := range stmt.Statements {
		b.WriteString(stmt.String())
		b.WriteString("\n")
	}
	b.WriteString("}")
	return b.String()
}

type StatementExpression struct {
	Expression
}

func (stmt *StatementExpression) GetLocation() Location {
	return stmt.Expression.GetLocation()
}

func (stmt *StatementExpression) String() string {
	var b strings.Builder
	if stmt.Expression != nil {
		b.WriteString(stmt.Expression.String())
	}
	b.WriteString(";")
	return b.String()
}

type StatementVar struct {
	VarToken   Token
	Identifier Token
	Expression
}

func (stmt *StatementVar) GetLocation() Location {
	return stmt.VarToken.Location
}

func (stmt *StatementVar) String() string {
	var b strings.Builder
	b.WriteString("var ")
	b.WriteString(stmt.Identifier.Lexeme)
	b.WriteString(" = ")
	if stmt.Expression != nil {
		b.WriteString(stmt.Expression.String())
		b.WriteString(";")
	} else {
		b.WriteString("<nil>;")
	}
	return b.String()
}

type StatementIf struct {
	IfToken   Token
	Condition Expression
	Then      Statement
	Else      Statement
}

func (stmt *StatementIf) GetLocation() Location {
	return stmt.IfToken.Location
}

func (stmt *StatementIf) String() string {
	var b strings.Builder
	b.WriteString("if ")
	b.WriteString(stmt.Condition.String())
	if stmtBlockThen, ok := stmt.Then.(*StatementBlock); ok {
		b.WriteString(" ")
		b.WriteString(stmtBlockThen.String())
	} else {
		b.WriteString(" {\n")
		b.WriteString(stmt.Then.String())
		b.WriteString("\n}")
	}
	if stmt.Else != nil {
		if stmtBlockElse, ok := stmt.Else.(*StatementBlock); ok {
			b.WriteString(" else ")
			b.WriteString(stmtBlockElse.String())
		} else {
			b.WriteString(" else {\n")
			b.WriteString(stmt.Else.String())
			b.WriteString("\n}")
		}
	}
	return b.String()
}

type StatementWhile struct {
	WhileToken Token
	Condition  Expression
	Body       Statement
}

func (stmt *StatementWhile) GetLocation() Location {
	return stmt.WhileToken.Location
}

func (stmt *StatementWhile) String() string {
	var b strings.Builder
	b.WriteString("while ")
	b.WriteString(stmt.Condition.String())

	if stmtBlockBody, ok := stmt.Body.(*StatementBlock); ok {
		b.WriteString(" ")
		b.WriteString(stmtBlockBody.String())
	} else {
		b.WriteString(" {\n")
		b.WriteString(stmt.Body.String())
		b.WriteString(" \n}")
	}
	return b.String()
}

type StatementFun struct {
	FunToken   Token
	Identifier Token
	Parameters []Token
	Body       []Statement
}

func (stmt *StatementFun) GetLocation() Location {
	if stmt.FunToken != (Token{}) {
		// is a function
		return stmt.FunToken.Location
	} else {
		// is a method
		return stmt.Identifier.Location
	}
}

func (stmt *StatementFun) String() string {
	var b strings.Builder
	b.WriteString("fun ")
	b.WriteString(stmt.Identifier.Lexeme)
	b.WriteString("(")
	for i, param := range stmt.Parameters {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(param.Lexeme)
	}
	b.WriteString(") {\n")
	for _, stmt := range stmt.Body {
		b.WriteString(stmt.String())
		b.WriteString("\n")
	}
	b.WriteString("}")
	return b.String()
}

type StatementReturn struct {
	ReturnToken Token
	Expression
}

func (stmt *StatementReturn) GetLocation() Location {
	return stmt.ReturnToken.Location
}

func (stmt *StatementReturn) String() string {
	var b strings.Builder
	b.WriteString("return")
	if stmt.Expression != nil {
		b.WriteString(" ")
		b.WriteString(stmt.Expression.String())
	}
	b.WriteString(";")
	return b.String()
}

type StatementClass struct {
	ClassToken Token
	Identifier Token
	Superclass *ExpressionVariable
	Methods    []*StatementFun
}

func (stmt *StatementClass) GetLocation() Location {
	return stmt.ClassToken.Location
}

func (stmt *StatementClass) String() string {
	var b strings.Builder
	b.WriteString("class ")
	b.WriteString(stmt.Identifier.Lexeme)
	if stmt.Superclass != nil {
		b.WriteString(" < ")
		b.WriteString(stmt.Superclass.Identifier.Lexeme)
	}
	b.WriteString(" {\n")
	for _, stmtMethod := range stmt.Methods {
		b.WriteString(stmtMethod.String())
		b.WriteString("\n")
	}
	b.WriteString("}")
	return b.String()
}

type StatementPrint struct {
	PrintToken Token
	Expression
}

func (stmt *StatementPrint) GetLocation() Location {
	return stmt.PrintToken.Location
}

func (stmt *StatementPrint) String() string {
	var b strings.Builder
	b.WriteString("print")
	if stmt.Expression != nil {
		b.WriteString(" ")
		b.WriteString(stmt.Expression.String())
	}
	b.WriteString(";")
	return b.String()
}
