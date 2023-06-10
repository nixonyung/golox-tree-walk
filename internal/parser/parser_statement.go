package parser

import golox "golox/internal"

func (p *Parser) statementBlock() (*golox.StatementBlock, error) {
	// matching: "{" STATEMENT* "}"
	result := &golox.StatementBlock{
		Location:   golox.Location{},
		Statements: []golox.Statement{},
	}

	if tkn, ok := p.expect(golox.TokenTypeLeftBrace); !ok {
		return nil, p.newErrorExpect(tkn, "expect block statement")
	} else {
		result.Location = tkn.Location
	}

	for {
		switch p.peekType() {
		case golox.TokenTypeEOF:
			return nil, p.newErrorMissingRightBrace(p.mustConsume(), result.Location)
		case golox.TokenTypeRightBrace:
			_ = p.mustConsume()
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

func (p *Parser) statementVar() (*golox.StatementVar, error) {
	// matching: "var" IDENTIFIER ("=" EXPRESSION)? ";"
	result := &golox.StatementVar{
		VarToken:   golox.Token{},
		Identifier: golox.Token{},
		Expression: nil,
	}

	if tkn, ok := p.expect(golox.TokenTypeVar); !ok {
		return nil, p.newErrorExpect(tkn, "expect 'var' keyword")
	} else {
		result.VarToken = tkn
	}

	if tkn, ok := p.expect(golox.TokenTypeIdentifier); !ok {
		return nil, p.newErrorExpect(tkn, "expect identifier after 'var'")
	} else {
		result.Identifier = tkn
	}

	if p.peekType() == golox.TokenTypeEqual {
		_ = p.mustConsume()

		if expr, err := p.parseExpression(); err != nil {
			return nil, err
		} else {
			result.Expression = expr
		}
	}

	if tkn, ok := p.expect(golox.TokenTypeSemicolon); !ok {
		return nil, p.newErrorExpect(tkn, "expect ';' after var statement")
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

	if tkn, ok := p.expect(golox.TokenTypeIf); !ok {
		return nil, p.newErrorExpect(tkn, "expect 'if' keyword")
	} else {
		result.IfToken = tkn
	}

	if tkn, ok := p.expect(golox.TokenTypeLeftParen); !ok {
		return nil, p.newErrorExpect(tkn, "expect '(' after 'if'")
	}

	if expr, err := p.parseExpression(); err != nil {
		return nil, err
	} else {
		result.Condition = expr
	}

	if tkn, ok := p.expect(golox.TokenTypeRightParen); !ok {
		return nil, p.newErrorExpect(tkn, "expect ')' after if condition")
	}

	if stmt, err := p.parseStatement(); err != nil {
		return nil, err
	} else {
		result.Then = stmt
	}

	if p.peekType() == golox.TokenTypeElse {
		_ = p.mustConsume()

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

	if tkn, ok := p.expect(golox.TokenTypeWhile); !ok {
		return nil, p.newErrorExpect(tkn, "expect 'while' keyword")
	} else {
		result.WhileToken = tkn
	}

	if tkn, ok := p.expect(golox.TokenTypeLeftParen); !ok {
		return nil, p.newErrorExpect(tkn, "expect '(' after 'while'")
	}

	if expr, err := p.parseExpression(); err != nil {
		return nil, err
	} else {
		result.Condition = expr
	}

	if tkn, ok := p.expect(golox.TokenTypeRightParen); !ok {
		return nil, p.newErrorExpect(tkn, "expect ')' after while condition")
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
	var result golox.Statement

	var forToken golox.Token
	if tkn, ok := p.expect(golox.TokenTypeFor); !ok {
		return nil, p.newErrorExpect(tkn, "expect 'for' keyword")
	} else {
		forToken = tkn
	}

	if tkn, ok := p.expect(golox.TokenTypeLeftParen); !ok {
		return nil, p.newErrorExpect(tkn, "expect '(' after 'for'")
	}

	var initializer golox.Statement
	hasInitializer := false
	switch p.peekType() {
	case golox.TokenTypeSemicolon:
		_ = p.mustConsume()
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
	switch p.peekType() {
	case golox.TokenTypeSemicolon:
		tkn := p.mustConsume()
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

		if tkn, ok := p.expect(golox.TokenTypeSemicolon); !ok {
			return nil, p.newErrorExpect(tkn, "expect ';' after loop condition")
		}
	}

	var increment golox.Expression
	hasIncrement := false
	switch p.peekType() {
	case golox.TokenTypeRightParen:
		_ = p.mustConsume()
	default:
		if expr, err := p.parseExpression(); err != nil {
			return nil, err
		} else {
			increment = expr
			hasIncrement = true
		}

		if tkn, ok := p.expect(golox.TokenTypeRightParen); !ok {
			return nil, p.newErrorExpect(tkn, "expect ')' after for clause")
		}
	}

	var body golox.Statement
	if stmt, err := p.parseStatement(); err != nil {
		return nil, err
	} else {
		body = stmt
	}

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
		if tkn, ok := p.expect(golox.TokenTypeFun); !ok {
			return nil, p.newErrorExpect(tkn, "expect 'fun' keyword")
		} else {
			result.FunToken = tkn
		}
	case FunctionTypeMethod:
		break
	}

	if tkn, ok := p.expect(golox.TokenTypeIdentifier); !ok {
		return nil, p.newErrorExpect(tkn, "expect function name")
	} else {
		result.Identifier = tkn
	}

	if tkn, ok := p.expect(golox.TokenTypeLeftParen); !ok {
		return nil, p.newErrorExpect(tkn, "expect '(' after function name")
	}

	if p.peekType() == golox.TokenTypeIdentifier {
		result.Parameters = append(result.Parameters, p.mustConsume())

		for p.peekType() == golox.TokenTypeComma {
			// see chapter 10.1.1
			if len(result.Parameters) >= 255 {
				return nil, p.newErrorFunctionTooManyParameters(p.mustConsume(), result.Identifier)
			} else {
				_ = p.mustConsume()

				if tkn, ok := p.expect(golox.TokenTypeIdentifier); !ok {
					return nil, p.newErrorExpect(tkn, "expect parameter name after ','")
				} else {
					result.Parameters = append(result.Parameters, tkn)
				}
			}
		}
	}

	if tkn, ok := p.expect(golox.TokenTypeRightParen); !ok {
		return nil, p.newErrorExpect(tkn, "expect ')' after function parameters")
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

	if tkn, ok := p.expect(golox.TokenTypeReturn); !ok {
		return nil, p.newErrorExpect(tkn, "expect 'return' keyword")
	} else {
		result.ReturnToken = tkn
	}

	if p.peekType() != golox.TokenTypeSemicolon {
		if expr, err := p.parseExpression(); err != nil {
			return nil, err
		} else {
			result.Expression = expr
		}
	}

	if tkn, ok := p.expect(golox.TokenTypeSemicolon); !ok {
		return nil, p.newErrorExpect(tkn, "expect ';' after return value")
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

	if tkn, ok := p.expect(golox.TokenTypeClass); !ok {
		return nil, p.newErrorExpect(tkn, "expect 'class' keyword")
	} else {
		result.ClassToken = tkn
	}

	if tkn, ok := p.expect(golox.TokenTypeIdentifier); !ok {
		return nil, p.newErrorExpect(tkn, "expect class name")
	} else {
		result.Identifier = tkn
	}

	if p.peekType() == golox.TokenTypeLess {
		_ = p.mustConsume()
		if tkn, ok := p.expect(golox.TokenTypeIdentifier); !ok {
			return nil, p.newErrorExpect(tkn, "expect superclass name after '<'")
		} else {
			result.Superclass = &golox.ExpressionVariable{
				Identifier: tkn,
			}
		}
	}

	if tkn, ok := p.expect(golox.TokenTypeLeftBrace); !ok {
		return nil, p.newErrorExpect(tkn, "expect '{' before class body")
	}

	for p.peekType() != golox.TokenTypeRightBrace {
		if stmt, err := p.statementFun(FunctionTypeMethod); err != nil {
			return nil, err
		} else {
			result.Methods = append(result.Methods, stmt)
		}
	}

	if tkn, ok := p.expect(golox.TokenTypeRightBrace); !ok {
		return nil, p.newErrorExpect(tkn, "expect '}' after class body")
	}

	return result, nil
}

func (p *Parser) statementPrint() (*golox.StatementPrint, error) {
	// matching: "print" EXPRESSION? ";"
	result := &golox.StatementPrint{
		PrintToken: golox.Token{},
		Expression: nil,
	}

	if tkn, ok := p.expect(golox.TokenTypePrint); !ok {
		return nil, p.newErrorExpect(tkn, "expect 'print' keyword")
	} else {
		result.PrintToken = tkn
	}

	if expr, err := p.parseExpression(); err != nil {
		return nil, err
	} else {
		result.Expression = expr
	}

	if tkn, ok := p.expect(golox.TokenTypeSemicolon); !ok {
		return nil, p.newErrorExpect(tkn, "expect ';' after print")
	}

	return result, nil
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

	if tkn, ok := p.expect(golox.TokenTypeSemicolon); !ok {
		return nil, p.newErrorExpect(tkn, "expect ';' after expression")
	}

	return result, nil
}
