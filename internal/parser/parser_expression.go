package parser

import golox "golox/internal"

func (p *Parser) expressionAssignment() (golox.Expression, error) {
	// matching: (EXPRESSION_VARIABLE|EXPRESSION_GET) ("=" EXPRESSION)*
	var lhs golox.Expression

	if expr, err := p.expressionLogicOr(); err != nil {
		return nil, err
	} else {
		lhs = expr
	}

	// right-associative: use recursion
	if p.peekType() != golox.TokenTypeEqual {
		return lhs, nil
	} else {
		equalTkn := p.mustConsume()

		// check lvalue type is valid:
		switch lhs.(type) {
		case *golox.ExpressionVariable, *golox.ExpressionGet:
			if rhs, err := p.expressionAssignment(); err != nil {
				return nil, err
			} else {
				switch lhs := lhs.(type) {
				case *golox.ExpressionVariable:
					return &golox.ExpressionAssignment{
						Identifier: lhs.Identifier,
						Value:      rhs,
					}, nil
				case *golox.ExpressionGet:
					return &golox.ExpressionSet{
						Object:     lhs.Object,
						Identifier: lhs.Identifier,
						Value:      rhs,
					}, nil
				}
			}
		}

		return nil, p.newErrorInvalidAssignmentLValue(equalTkn)
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
	for p.peekType() == golox.TokenTypeOr {
		tkn := p.mustConsume()
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
	for p.peekType() == golox.TokenTypeAnd {
		tkn := p.mustConsume()
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
	for p.peekType() == golox.TokenTypeBangEqual ||
		p.peekType() == golox.TokenTypeEqualEqual {
		tkn := p.mustConsume()
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
	for p.peekType() == golox.TokenTypeGreater ||
		p.peekType() == golox.TokenTypeGreaterEqual ||
		p.peekType() == golox.TokenTypeLess ||
		p.peekType() == golox.TokenTypeLessEqual {
		tkn := p.mustConsume()
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
	for p.peekType() == golox.TokenTypeMinus ||
		p.peekType() == golox.TokenTypePlus {
		tkn := p.mustConsume()
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
	for p.peekType() == golox.TokenTypeStar ||
		p.peekType() == golox.TokenTypeSlash {
		tkn := p.mustConsume()
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
	if p.peekType() != golox.TokenTypeBang &&
		p.peekType() != golox.TokenTypeMinus {
		return p.expressionCall()
	} else {
		tkn := p.mustConsume()
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
	// matching: EXPRESSION (("(" (IDENTIFIER ("," IDENTIFIER)*)? ")")|"." IDENTIFIER))*
	var lhs golox.Expression

	if expr, err := p.expressionPrimary(); err != nil {
		return nil, err
	} else {
		lhs = expr
	}

	// left-associative: use while loop
	for {
		switch p.peekType() {
		case golox.TokenTypeDot:
			p.mustConsume()
			if tkn, ok := p.expect(golox.TokenTypeIdentifier); !ok {
				return nil, p.newErrorExpect(tkn, "expect property name after '.'")
			} else {
				lhs = &golox.ExpressionGet{
					Object:     lhs,
					Identifier: tkn,
				}
			}
		case golox.TokenTypeLeftParen:
			p.mustConsume()
			arguments := []golox.Expression{}

			if p.peekType() != golox.TokenTypeRightParen {
				if expr, err := p.parseExpression(); err != nil {
					return nil, err
				} else {
					arguments = append(arguments, expr)
				}
			}

			for p.peekType() == golox.TokenTypeComma {
				// see chapter 10.1.1
				if len(arguments) >= 255 {
					return nil, p.newErrorFunctionCallTooManyArguments(p.mustConsume())
				} else {
					_ = p.mustConsume()
					if expr, err := p.parseExpression(); err != nil {
						return nil, err
					} else {
						arguments = append(arguments, expr)
					}
				}
			}

			if tkn, ok := p.expect(golox.TokenTypeRightParen); !ok {
				return nil, p.newErrorExpect(tkn, "expect ')' after function arguments")
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
	switch p.peekType() {
	case golox.TokenTypeFalse:
		tkn := p.mustConsume()
		return &golox.ExpressionLiteral{
			Location:     tkn.Location,
			LiteralValue: false,
		}, nil
	case golox.TokenTypeTrue:
		tkn := p.mustConsume()
		return &golox.ExpressionLiteral{
			Location:     tkn.Location,
			LiteralValue: true,
		}, nil
	case golox.TokenTypeNil:
		tkn := p.mustConsume()
		return &golox.ExpressionLiteral{
			Location:     tkn.Location,
			LiteralValue: nil,
		}, nil
	case golox.TokenTypeString:
		tkn := p.mustConsume()
		return &golox.ExpressionLiteral{
			Location:     tkn.Location,
			LiteralValue: tkn.LiteralValue,
		}, nil
	case golox.TokenTypeNumber:
		tkn := p.mustConsume()
		return &golox.ExpressionLiteral{
			Location:     tkn.Location,
			LiteralValue: tkn.LiteralValue,
		}, nil
	case golox.TokenTypeLeftParen:
		result := &golox.ExpressionGrouping{
			LeftParenToken: golox.Token{},
			Expression:     nil,
		}
		result.LeftParenToken = p.mustConsume()

		if expr, err := p.parseExpression(); err != nil {
			return nil, err
		} else {
			result.Expression = expr

			if tkn, ok := p.expect(golox.TokenTypeRightParen); !ok {
				return nil, p.newErrorExpect(tkn, "expect closing ')'")
			} else {
				return result, nil
			}
		}
	case golox.TokenTypeIdentifier:
		tkn := p.mustConsume()
		return &golox.ExpressionVariable{Identifier: tkn}, nil
	case golox.TokenTypeThis:
		tkn := p.mustConsume()
		return &golox.ExpressionThis{ThisToken: tkn}, nil
	case golox.TokenTypeSuper:
		result := &golox.ExpressionSuper{
			SuperToken: golox.Token{},
			Method:     golox.Token{},
		}
		result.SuperToken = p.mustConsume()

		if tkn, ok := p.expect(golox.TokenTypeDot); !ok {
			return nil, p.newErrorExpect(tkn, "expect '.' after 'super'")
		} else if tkn, ok := p.expect(golox.TokenTypeIdentifier); !ok {
			return nil, p.newErrorExpect(tkn, "expect superclass method name after 'super.'")
		} else {
			result.Method = tkn
			return result, nil
		}
	default:
		return nil, p.newErrorExpect(p.mustConsume(), "expect expression")
	}
}
