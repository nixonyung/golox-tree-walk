package parser

import (
	"fmt"
	golox "golox/internal"
	"strconv"
	"strings"
)

type Parser struct {
	tokens []golox.Token     // input
	stmts  []golox.Statement // output
	curr   int               // index to current token
	errors []error
}

func (p *Parser) peekTokenType() golox.TokenType {
	if p.curr >= len(p.tokens) {
		return golox.TokenTypeEOF
	} else {
		return p.tokens[p.curr].TokenType
	}
}

func (p *Parser) skipToken() golox.Token {
	tkn := p.tokens[p.curr]
	p.curr++
	return tkn
}

func (p *Parser) expectTokenType(tokenType golox.TokenType) (golox.Token, bool) {
	tkn := p.tokens[p.curr]
	if tkn.TokenType == tokenType {
		p.curr++
		return tkn, true
	} else {
		// return tkn on fail to provide location info
		return tkn, false
	}
}

func (p *Parser) parseDeclaration() (golox.Statement, error) {
	var stmt golox.Statement
	var err error
	switch p.peekTokenType() {
	case golox.TokenTypeVar:
		stmt, err = p.statementVar()
	case golox.TokenTypeFun:
		stmt, err = p.statementFun(FunctionTypeFunction)
	case golox.TokenTypeClass:
		stmt, err = p.statementClass()
	default:
		stmt, err = p.parseStatement()
	}

	if err != nil {
		// "synchronize": skip tokens until we find the next declaration/statement
		for {
			switch p.peekTokenType() {
			case golox.TokenTypeEOF:
				goto L_SYNCHRONIZE_END
			case golox.TokenTypeSemicolon:
				_ = p.skipToken()
				goto L_SYNCHRONIZE_END
			case golox.TokenTypeVar,
				golox.TokenTypeIf,
				golox.TokenTypeWhile,
				golox.TokenTypeFor,
				golox.TokenTypeFun,
				golox.TokenTypeReturn,
				golox.TokenTypeClass,
				golox.TokenTypePrint:
				goto L_SYNCHRONIZE_END
			default:
				_ = p.skipToken()
			}
		}
	}
L_SYNCHRONIZE_END:

	return stmt, err
}

func (p *Parser) parseStatement() (golox.Statement, error) {
	switch p.peekTokenType() {
	case golox.TokenTypeVar,
		golox.TokenTypeFun,
		golox.TokenTypeClass:
		tkn := p.skipToken()
		return nil, golox.NewErrorf(
			tkn.Location,
			"expect statement but not declaration",
		)
	case golox.TokenTypeEOF:
		return nil, nil
	case golox.TokenTypeLeftBrace:
		return p.statementBlock()
	case golox.TokenTypeIf:
		return p.statementIf()
	case golox.TokenTypeWhile:
		return p.statementWhile()
	case golox.TokenTypeFor:
		return p.statementFor()
	case golox.TokenTypeReturn:
		return p.statementReturn()
	case golox.TokenTypePrint:
		return p.statementPrint()
	default:
		return p.statementExpression()
	}
}

func (p *Parser) statementBlock() (*golox.StatementBlock, error) {
	// matching: "{" STATEMENT* "}"
	result := &golox.StatementBlock{
		Location:   golox.Location{},
		Statements: []golox.Statement{},
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeLeftBrace); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect block statement")
	} else {
		result.Location = tkn.Location
	}

	for {
		switch p.peekTokenType() {
		case golox.TokenTypeEOF:
			tkn := p.skipToken()
			return nil, golox.NewErrorf(
				tkn.Location,
				"missing closing '}', started at line %d:%d",
				result.Location.Line, result.Location.Col,
			)
		case golox.TokenTypeRightBrace:
			_ = p.skipToken()
			return result, nil
		default:
			if stmt, err := p.parseDeclaration(); err != nil {
				return nil, err
			} else {
				result.Statements = append(result.Statements, stmt)
			}
		}
	}
}

func (p *Parser) statementExpression() (*golox.StatementExpression, error) {
	// matching: EXPRESSION? ";"
	result := &golox.StatementExpression{
		Expression: nil,
	}

	if expr, err := p.parseExpression(); err != nil {
		return nil, err
	} else {
		result.Expression = expr
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeSemicolon); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect ';' after expression")
	}

	return result, nil
}

func (p *Parser) statementVar() (*golox.StatementVar, error) {
	// matching: "var" IDENTIFIER ("=" EXPRESSION)? ";"
	result := &golox.StatementVar{
		VarToken:   golox.Token{},
		Identifier: golox.Token{},
		Expression: nil,
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeVar); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect 'var' keyword")
	} else {
		result.VarToken = tkn
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeIdentifier); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect identifier after 'var'")
	} else {
		result.Identifier = tkn
	}

	if p.peekTokenType() == golox.TokenTypeEqual {
		_ = p.skipToken()

		if expr, err := p.parseExpression(); err != nil {
			return nil, err
		} else {
			result.Expression = expr
		}
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeSemicolon); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect ';' after var statement")
	}

	return result, nil
}

func (p *Parser) statementIf() (*golox.StatementIf, error) {
	// matching: "if" "(" EXPRESSION ")" STATEMENT ("else" STATEMENT)?
	result := &golox.StatementIf{
		IfToken:   golox.Token{},
		Condition: nil,
		Then:      nil,
		Else:      nil,
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeIf); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect 'if' keyword")
	} else {
		result.IfToken = tkn
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeLeftParen); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect '(' after 'if'")
	}

	if expr, err := p.parseExpression(); err != nil {
		return nil, err
	} else {
		result.Condition = expr
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeRightParen); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect ')' after if condition")
	}

	if stmt, err := p.parseStatement(); err != nil {
		return nil, err
	} else {
		result.Then = stmt
	}

	if p.peekTokenType() == golox.TokenTypeElse {
		_ = p.skipToken()

		if stmt, err := p.parseStatement(); err != nil {
			return nil, err
		} else {
			result.Else = stmt
		}
	}

	return result, nil
}

func (p *Parser) statementWhile() (*golox.StatementWhile, error) {
	// matching: "while" "(" EXPRESSION ")" STATEMENT
	result := &golox.StatementWhile{
		WhileToken: golox.Token{},
		Condition:  nil,
		Body:       nil,
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeWhile); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect 'while' keyword")
	} else {
		result.WhileToken = tkn
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeLeftParen); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect '(' after 'while'")
	}

	if expr, err := p.parseExpression(); err != nil {
		return nil, err
	} else {
		result.Condition = expr
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeRightParen); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect ')' after while condition")
	}

	if stmt, err := p.parseStatement(); err != nil {
		return nil, err
	} else {
		result.Body = stmt
	}

	return result, nil
}

func (p *Parser) statementFor() (golox.Statement, error) {
	// matching: "for" "(" (STATEMENT_VAR|STATEMENT_EXPRESSION|";") EXPRESSION? ";" EXPRESSION? ")" STATEMENT

	var forToken golox.Token
	if tkn, ok := p.expectTokenType(golox.TokenTypeFor); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect 'for' keyword")
	} else {
		forToken = tkn
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeLeftParen); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect '(' after 'for'")
	}

	var initializer golox.Statement
	hasInitializer := false
	switch p.peekTokenType() {
	case golox.TokenTypeSemicolon:
		_ = p.skipToken()
	case golox.TokenTypeVar:
		if stmt, err := p.statementVar(); err != nil {
			return nil, err
		} else {
			initializer = stmt
			hasInitializer = true
		}
	default:
		if stmt, err := p.statementExpression(); err != nil {
			return nil, err
		} else {
			initializer = stmt
			hasInitializer = true
		}
	}

	var condition golox.Expression
	switch p.peekTokenType() {
	case golox.TokenTypeSemicolon:
		tkn := p.skipToken()
		// artificial true condition
		condition = &golox.ExpressionLiteral{
			Location:     tkn.Location,
			LiteralValue: true,
		}
	default:
		if expr, err := p.parseExpression(); err != nil {
			return nil, err
		} else {
			condition = expr
		}

		if tkn, ok := p.expectTokenType(golox.TokenTypeSemicolon); !ok {
			return nil, golox.NewErrorf(tkn.Location, "expect ';' after loop condition")
		}
	}

	var increment golox.Expression
	hasIncrement := false
	switch p.peekTokenType() {
	case golox.TokenTypeRightParen:
		_ = p.skipToken()
	default:
		if expr, err := p.parseExpression(); err != nil {
			return nil, err
		} else {
			increment = expr
			hasIncrement = true
		}

		if tkn, ok := p.expectTokenType(golox.TokenTypeRightParen); !ok {
			return nil, golox.NewErrorf(tkn.Location, "expect ')' after for clause")
		}
	}

	var body golox.Statement
	if stmt, err := p.parseStatement(); err != nil {
		return nil, err
	} else {
		body = stmt
	}

	// statementFor is a syntactic sugar
	//
	// {
	//     INITIALIZER;
	//     while(CONDITION) {
	//         BODY;  // BODY could be a single statement or a block statement
	//         INCREMENT;
	//     }
	// }
	var result golox.Statement

	if hasIncrement {
		body = &golox.StatementBlock{
			Location: body.GetLocation(),
			Statements: []golox.Statement{
				body,
				&golox.StatementExpression{Expression: increment},
			},
		}
	}

	result = &golox.StatementWhile{
		WhileToken: forToken,
		Condition:  condition,
		Body:       body,
	}

	if hasInitializer {
		result = &golox.StatementBlock{
			Location:   result.GetLocation(),
			Statements: []golox.Statement{initializer, result},
		}
	}

	return result, nil
}

func (p *Parser) statementFun(fnType FunctionType) (*golox.StatementFun, error) {
	// matching: "fun"? IDENTIFIER "(" (PARAMETER ("," PARAMETER))? ")" STATEMENT_BLOCK
	result := &golox.StatementFun{
		FunToken:   golox.Token{},
		Identifier: golox.Token{},
		Parameters: []golox.Token{},
		Body:       nil,
	}

	switch fnType {
	case FunctionTypeFunction:
		if tkn, ok := p.expectTokenType(golox.TokenTypeFun); !ok {
			return nil, golox.NewErrorf(tkn.Location, "expect 'fun' keyword")
		} else {
			result.FunToken = tkn
		}
	case FunctionTypeMethod:
		break
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeIdentifier); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect function name")
	} else {
		result.Identifier = tkn
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeLeftParen); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect '(' after function name")
	}

	if p.peekTokenType() != golox.TokenTypeRightParen {
		for {
			if tkn, ok := p.expectTokenType(golox.TokenTypeIdentifier); !ok {
				return nil, golox.NewErrorf(tkn.Location, "expect parameter name after ','")
			} else {
				result.Parameters = append(result.Parameters, tkn)
			}

			if p.peekTokenType() != golox.TokenTypeComma {
				break
			} else {
				tkn := p.skipToken()
				if len(result.Parameters) >= 255 {
					return nil, golox.NewErrorf(
						tkn.Location,
						"function '%s' cannot have more than 255 parameters",
						result.Identifier.Lexeme,
					)
				}
			}
		}
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeRightParen); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect ')' after function parameters")
	}

	if stmt, err := p.statementBlock(); err != nil {
		return nil, err
	} else {
		result.Body = stmt.Statements
	}

	return result, nil
}

func (p *Parser) statementReturn() (*golox.StatementReturn, error) {
	// matching: "return" EXPRESSION? ";"
	result := &golox.StatementReturn{
		ReturnToken: golox.Token{},
		Expression:  nil,
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeReturn); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect 'return' keyword")
	} else {
		result.ReturnToken = tkn
	}

	if p.peekTokenType() != golox.TokenTypeSemicolon {
		if expr, err := p.parseExpression(); err != nil {
			return nil, err
		} else {
			result.Expression = expr
		}
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeSemicolon); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect ';' after return value")
	}

	return result, nil
}

func (p *Parser) statementClass() (*golox.StatementClass, error) {
	// matching: "class" IDENTIFIER ("<" IDENTIFIER)? "{" STATEMENT_FUNCTION* "}"
	result := &golox.StatementClass{
		ClassToken: golox.Token{},
		Identifier: golox.Token{},
		Superclass: nil,
		Methods:    []*golox.StatementFun{},
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeClass); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect 'class' keyword")
	} else {
		result.ClassToken = tkn
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeIdentifier); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect class name")
	} else {
		result.Identifier = tkn
	}

	if p.peekTokenType() == golox.TokenTypeLess {
		_ = p.skipToken()
		if tkn, ok := p.expectTokenType(golox.TokenTypeIdentifier); !ok {
			return nil, golox.NewErrorf(tkn.Location, "expect superclass name after '<'")
		} else {
			result.Superclass = &golox.ExpressionVariable{
				Identifier: tkn,
			}
		}
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeLeftBrace); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect '{' before class body")
	}

	for p.peekTokenType() != golox.TokenTypeRightBrace {
		if stmt, err := p.statementFun(FunctionTypeMethod); err != nil {
			return nil, err
		} else {
			result.Methods = append(result.Methods, stmt)
		}
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeRightBrace); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect '}' after class body")
	}

	return result, nil
}

func (p *Parser) statementPrint() (*golox.StatementPrint, error) {
	// matching: "print" EXPRESSION? ";"
	result := &golox.StatementPrint{
		PrintToken: golox.Token{},
		Expression: nil,
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypePrint); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect 'print' keyword")
	} else {
		result.PrintToken = tkn
	}

	if expr, err := p.parseExpression(); err != nil {
		return nil, err
	} else {
		result.Expression = expr
	}

	if tkn, ok := p.expectTokenType(golox.TokenTypeSemicolon); !ok {
		return nil, golox.NewErrorf(tkn.Location, "expect ';' after print")
	}

	return result, nil
}

func (p *Parser) parseExpression() (golox.Expression, error) {
	// recursive descent: start from the lowest precedence
	return p.expressionAssignment()
}

func (p *Parser) expressionAssignment() (golox.Expression, error) {
	// matching: (EXPRESSION_VARIABLE|EXPRESSION_GET) ("=" EXPRESSION)*
	var lhs golox.Expression

	if expr, err := p.expressionLogicOr(); err != nil {
		return nil, err
	} else {
		lhs = expr
	}

	// right-associative: use recursion
	if p.peekTokenType() != golox.TokenTypeEqual {
		return lhs, nil
	} else {
		equalTkn := p.skipToken()

		// check lvalue type is valid:
		switch lhs := lhs.(type) {
		case *golox.ExpressionVariable:
			if rhs, err := p.expressionAssignment(); err != nil {
				return nil, err
			} else {
				return &golox.ExpressionAssignment{
					Identifier: lhs.Identifier,
					Value:      rhs,
				}, nil
			}
		case *golox.ExpressionGet:
			if rhs, err := p.expressionAssignment(); err != nil {
				return nil, err
			} else {
				return &golox.ExpressionSet{
					Object:     lhs.Object,
					Identifier: lhs.Identifier,
					Value:      rhs,
				}, nil
			}
		default:
			return nil, golox.NewErrorf(
				equalTkn.Location,
				"invalid assignment lvalue",
			)
		}
	}
}

func (p *Parser) expressionLogicOr() (golox.Expression, error) {
	// matching: EXPRESSION ("or" EXPRESSION)*
	var lhs golox.Expression

	if expr, err := p.expressionLogicAnd(); err != nil {
		return nil, err
	} else {
		lhs = expr
	}

	// left-associative: use while loop
	for p.peekTokenType() == golox.TokenTypeOr {
		tkn := p.skipToken()
		if rhs, err := p.expressionLogicOr(); err != nil {
			return nil, err
		} else {
			lhs = &golox.ExpressionLogical{
				Left:     lhs,
				Operator: tkn,
				Right:    rhs,
			}
		}
	}

	return lhs, nil
}

func (p *Parser) expressionLogicAnd() (golox.Expression, error) {
	// matching: EXPRESSION ("and" EXPRESSION)*
	var lhs golox.Expression

	if expr, err := p.expressionEquality(); err != nil {
		return nil, err
	} else {
		lhs = expr
	}

	// left-associative: use while loop
	for p.peekTokenType() == golox.TokenTypeAnd {
		tkn := p.skipToken()
		if rhs, err := p.expressionEquality(); err != nil {
			return nil, err
		} else {
			lhs = &golox.ExpressionLogical{
				Left:     lhs,
				Operator: tkn,
				Right:    rhs,
			}
		}
	}

	return lhs, nil
}

func (p *Parser) expressionEquality() (golox.Expression, error) {
	// matching: EXPRESSION (("!="|"==") EXPRESSION)*
	var lhs golox.Expression

	if expr, err := p.expressionComparison(); err != nil {
		return nil, err
	} else {
		lhs = expr
	}

	// left-associative: use while loop
	for p.peekTokenType() == golox.TokenTypeBangEqual ||
		p.peekTokenType() == golox.TokenTypeEqualEqual {
		tkn := p.skipToken()
		if rhs, err := p.expressionComparison(); err != nil {
			return nil, err
		} else {
			lhs = &golox.ExpressionBinary{
				Left:     lhs,
				Operator: tkn,
				Right:    rhs,
			}
		}
	}

	return lhs, nil
}

func (p *Parser) expressionComparison() (golox.Expression, error) {
	// matching: EXPRESSION ((">"|">="|"<"|"<=") EXPRESSION)*
	var lhs golox.Expression

	if expr, err := p.expressionTerm(); err != nil {
		return nil, err
	} else {
		lhs = expr
	}

	// left-associative: use while loop
	for p.peekTokenType() == golox.TokenTypeGreater ||
		p.peekTokenType() == golox.TokenTypeGreaterEqual ||
		p.peekTokenType() == golox.TokenTypeLess ||
		p.peekTokenType() == golox.TokenTypeLessEqual {
		tkn := p.skipToken()
		if rhs, err := p.expressionTerm(); err != nil {
			return nil, err
		} else {
			lhs = &golox.ExpressionBinary{
				Left:     lhs,
				Operator: tkn,
				Right:    rhs,
			}
		}
	}

	return lhs, nil
}

func (p *Parser) expressionTerm() (golox.Expression, error) {
	// matching: EXPRESSION (("-"|"+") EXPRESSION)*
	var lhs golox.Expression

	if expr, err := p.expressionFactor(); err != nil {
		return nil, err
	} else {
		lhs = expr
	}

	// left-associative
	for p.peekTokenType() == golox.TokenTypeMinus ||
		p.peekTokenType() == golox.TokenTypePlus {
		tkn := p.skipToken()
		if rhs, err := p.expressionFactor(); err != nil {
			return nil, err
		} else {
			lhs = &golox.ExpressionBinary{
				Left:     lhs,
				Operator: tkn,
				Right:    rhs,
			}
		}
	}

	return lhs, nil
}

func (p *Parser) expressionFactor() (golox.Expression, error) {
	// matching: EXPRESSION (("*"|"/") EXPRESSION)*
	var lhs golox.Expression

	if expr, err := p.expressionUnary(); err != nil {
		return nil, err
	} else {
		lhs = expr
	}

	// left-associative
	for p.peekTokenType() == golox.TokenTypeStar ||
		p.peekTokenType() == golox.TokenTypeSlash {
		tkn := p.skipToken()
		if rhs, err := p.expressionUnary(); err != nil {
			return nil, err
		} else {
			lhs = &golox.ExpressionBinary{
				Left:     lhs,
				Operator: tkn,
				Right:    rhs,
			}
		}
	}

	return lhs, nil
}

func (p *Parser) expressionUnary() (golox.Expression, error) {
	// matching: ("!"|"-")* EXPRESSION

	// right-associative
	if p.peekTokenType() != golox.TokenTypeBang &&
		p.peekTokenType() != golox.TokenTypeMinus {
		return p.expressionCall()
	} else {
		tkn := p.skipToken()
		if rhs, err := p.expressionUnary(); err != nil {
			return nil, err
		} else {
			return &golox.ExpressionUnary{
				Operator: tkn,
				Right:    rhs,
			}, nil
		}
	}
}

func (p *Parser) expressionCall() (golox.Expression, error) {
	// matching: EXPRESSION ("." IDENTIFIER | ("(" (IDENTIFIER ("," IDENTIFIER)*)? ")"))*
	var lhs golox.Expression

	if expr, err := p.expressionPrimary(); err != nil {
		return nil, err
	} else {
		lhs = expr
	}

	// left-associative: use while loop
	for {
		switch p.peekTokenType() {
		case golox.TokenTypeDot:
			_ = p.skipToken()

			if tkn, ok := p.expectTokenType(golox.TokenTypeIdentifier); !ok {
				return nil, golox.NewErrorf(tkn.Location, "expect property name after '.'")
			} else {
				lhs = &golox.ExpressionGet{
					Object:     lhs,
					Identifier: tkn,
				}
			}
		case golox.TokenTypeLeftParen:
			_ = p.skipToken()

			arguments := []golox.Expression{}
			if p.peekTokenType() != golox.TokenTypeRightParen {
				for {
					if expr, err := p.parseExpression(); err != nil {
						return nil, err
					} else {
						arguments = append(arguments, expr)
					}

					if p.peekTokenType() != golox.TokenTypeComma {
						break
					} else {
						tkn := p.skipToken()
						if len(arguments) >= 255 {
							return nil, golox.NewErrorf(
								tkn.Location,
								"function call cannot have more than 255 parameters",
							)
						}
					}
				}
			}

			if tkn, ok := p.expectTokenType(golox.TokenTypeRightParen); !ok {
				return nil, golox.NewErrorf(tkn.Location, "expect ')' after function arguments")
			} else {
				lhs = &golox.ExpressionCall{
					Callee:     lhs,
					RightParen: tkn,
					Arguments:  arguments,
				}
			}
		default:
			return lhs, nil
		}
	}
}

func (p *Parser) expressionPrimary() (golox.Expression, error) {
	switch p.peekTokenType() {
	case golox.TokenTypeFalse:
		tkn := p.skipToken()
		return &golox.ExpressionLiteral{
			Location:     tkn.Location,
			LiteralValue: false,
		}, nil
	case golox.TokenTypeTrue:
		tkn := p.skipToken()
		return &golox.ExpressionLiteral{
			Location:     tkn.Location,
			LiteralValue: true,
		}, nil
	case golox.TokenTypeNil:
		tkn := p.skipToken()
		return &golox.ExpressionLiteral{
			Location:     tkn.Location,
			LiteralValue: nil,
		}, nil
	case golox.TokenTypeString:
		tkn := p.skipToken()
		return &golox.ExpressionLiteral{
			Location:     tkn.Location,
			LiteralValue: tkn.LiteralValue,
		}, nil
	case golox.TokenTypeNumber:
		tkn := p.skipToken()
		return &golox.ExpressionLiteral{
			Location:     tkn.Location,
			LiteralValue: tkn.LiteralValue,
		}, nil
	case golox.TokenTypeLeftParen:
		result := &golox.ExpressionGrouping{
			LeftParenToken: golox.Token{},
			Expression:     nil,
		}
		result.LeftParenToken = p.skipToken()

		if expr, err := p.parseExpression(); err != nil {
			return nil, err
		} else {
			result.Expression = expr

			if tkn, ok := p.expectTokenType(golox.TokenTypeRightParen); !ok {
				return nil, golox.NewErrorf(tkn.Location, "expect closing ')'")
			}

			return result, nil
		}
	case golox.TokenTypeIdentifier:
		tkn := p.skipToken()
		return &golox.ExpressionVariable{Identifier: tkn}, nil
	case golox.TokenTypeThis:
		tkn := p.skipToken()
		return &golox.ExpressionThis{ThisToken: tkn}, nil
	case golox.TokenTypeSuper:
		result := &golox.ExpressionSuper{
			SuperToken: golox.Token{},
			Method:     golox.Token{},
		}
		result.SuperToken = p.skipToken()

		if tkn, ok := p.expectTokenType(golox.TokenTypeDot); !ok {
			return nil, golox.NewErrorf(tkn.Location, "expect '.' after 'super'")
		}

		if tkn, ok := p.expectTokenType(golox.TokenTypeIdentifier); !ok {
			return nil, golox.NewErrorf(tkn.Location, "expect superclass method name after 'super.'")
		} else {
			result.Method = tkn
		}

		return result, nil
	default:
		tkn := p.skipToken()
		return nil, golox.NewErrorf(tkn.Location, "expect expression")
	}
}

func (p *Parser) logParsedStatement(stmt golox.Statement) {
	golox.Logf(golox.ModuleParser, "")
	golox.Logf(golox.ModuleParser, "%s", stmt.GetLocation())
	golox.Logf(golox.ModuleParser, "|")
	indentLevel := 0
	for _, line := range strings.Split(stmt.String(), "\n") {
		if strings.HasPrefix(line, "}") {
			indentLevel--
		}

		golox.Logf(golox.ModuleParser, "|%s%s",
			strings.Repeat(golox.STMT_INDENT, indentLevel), line,
		)

		if strings.HasSuffix(line, "{") {
			indentLevel++
		}
	}
	golox.Logf(golox.ModuleParser, "|")
}

func (p *Parser) StatementsFromTokens(tokens []golox.Token) ([]golox.Statement, error) {
	p.tokens = tokens
	p.stmts = []golox.Statement{}
	p.curr = 0

	for {
		if stmt, err := p.parseDeclaration(); err != nil {
			p.errors = append(p.errors, err)
		} else if stmt == nil {
			break
		} else {
			p.stmts = append(p.stmts, stmt)
			p.logParsedStatement(stmt)
		}
	}

	if len(p.errors) == 0 {
		return p.stmts, nil
	} else {
		var builder strings.Builder

		builder.WriteString("------\nsyntax errors:\n------\n")
		for _, err := range p.errors {
			builder.WriteString(err.Error())
			builder.WriteByte('\n')
		}
		builder.WriteString("------\n")
		builder.WriteString("total ")
		builder.WriteString(strconv.Itoa(len(p.errors)))
		builder.WriteString(" errors")

		return nil, fmt.Errorf(builder.String())
	}
}

func NewParser() *Parser {
	return &Parser{}
}
