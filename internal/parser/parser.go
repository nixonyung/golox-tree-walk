package parser

import (
	"fmt"
	golox "golox/internal"
	"strconv"
	"strings"
)

type Parser struct {
	// configs:
	isDebug bool

	// inputs:
	tokens []golox.Token

	// outputs:
	stmts []golox.Statement

	// states:
	curr        int
	errors      []error
	indentLevel int // debug use
}

// only requires 1-lookahead

func (p *Parser) peekType() golox.TokenType {
	if p.curr >= len(p.tokens) {
		return golox.TokenTypeEOF
	} else {
		return p.tokens[p.curr].TokenType
	}
}

func (p *Parser) expect(tokenTypes ...golox.TokenType) (golox.Token, bool) {
	tkn := p.tokens[p.curr]

	if len(tokenTypes) == 0 {
		p.curr++
		return tkn, true
	} else {
		for _, tknType := range tokenTypes {
			if tkn.TokenType == tknType {
				p.curr++
				return tkn, true
			}
		}
		// return tkn on fail to provide location info
		return tkn, false
	}
}

func (p *Parser) mustConsume() golox.Token {
	tkn := p.tokens[p.curr]
	p.curr++
	return tkn
}

func (p *Parser) synchronize() {
	for {
		switch p.peekType() {
		case golox.TokenTypeEOF:
			return
		case golox.TokenTypeSemicolon:
			_ = p.mustConsume()
			return
		case golox.TokenTypeClass,
			golox.TokenTypeFun,
			golox.TokenTypeVar,
			golox.TokenTypeFor,
			golox.TokenTypeIf,
			golox.TokenTypeWhile,
			golox.TokenTypePrint,
			golox.TokenTypeReturn:
			return
		default:
			_ = p.mustConsume()
		}
	}
}

func (p *Parser) parseDeclaration() (golox.Statement, error) {
	switch p.peekType() {
	case golox.TokenTypeVar:
		return p.statementVar()
	case golox.TokenTypeFun:
		return p.statementFun(FunctionTypeFunction)
	case golox.TokenTypeClass:
		return p.statementClass()
	default:
		return p.parseStatement()
	}
}

func (p *Parser) parseStatement() (golox.Statement, error) {
	switch p.peekType() {
	case golox.TokenTypeVar,
		golox.TokenTypeFun,
		golox.TokenTypeClass:
		return nil, p.newErrorUnexpectedDeclaration(p.mustConsume())
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

func (p *Parser) parseExpression() (golox.Expression, error) {
	// recursive descent: start from the lowest precedence
	return p.expressionAssignment()
}

func (p *Parser) StatementsFromTokens(tokens []golox.Token) ([]golox.Statement, error) {
	p.tokens = tokens
	p.stmts = []golox.Statement{}
	p.curr = 0
	p.indentLevel = 0

	for {
		if stmt, err := p.parseDeclaration(); err != nil {
			p.errors = append(p.errors, err)
			p.synchronize()
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

func NewParser(
	isDebug bool,
) *Parser {
	return &Parser{
		isDebug: isDebug,
	}
}
