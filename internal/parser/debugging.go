package parser

import (
	"fmt"
	golox "golox/internal"
	"log"
	"strings"
)

const (
	log_prefix = "DEBUG :: parser"
)

func (p *Parser) logParsedStatement(
	stmt golox.Statement,
) {
	if p.isDebug {
		log.Printf("%s:",
			log_prefix,
		)
		log.Printf("%s: %s",
			log_prefix, stmt.GetLocation(),
		)
		log.Printf("%s: |",
			log_prefix,
		)

		indentLevel := 0

		for _, line := range strings.Split(stmt.String(), "\n") {
			if strings.HasPrefix(line, "}") {
				indentLevel--
			}

			log.Printf("%s: |%s%s",
				log_prefix, strings.Repeat(golox.INDENT, indentLevel), line,
			)

			if strings.HasSuffix(line, "{") {
				indentLevel++
			}
		}

		log.Printf("%s: |",
			log_prefix,
		)
	}
}

func (p *Parser) newErrorUnexpectedDeclaration(
	gotToken golox.Token,
) error {
	return fmt.Errorf("%s: expect statement but not declaration",
		gotToken.Location,
	)
}

func (p *Parser) newErrorExpect(
	gotToken golox.Token,
	message string,
) error {
	return fmt.Errorf("%s: %s",
		gotToken.Location, message,
	)
}

func (p *Parser) newErrorMissingRightBrace(
	gotToken golox.Token,
	leftBraceLocation golox.Location,
) error {
	return fmt.Errorf("%s: missing closing '}', started at line %d:%d",
		gotToken.Location, leftBraceLocation.Line, leftBraceLocation.Col,
	)
}

func (p *Parser) newErrorInvalidAssignmentLValue(
	equalToken golox.Token,
) error {
	return fmt.Errorf("%s: invalid assignment lvalue",
		equalToken.Location,
	)
}

func (p *Parser) newErrorFunctionTooManyParameters(
	commaToken golox.Token,
	functionIdentifier golox.Token,
) error {
	return fmt.Errorf("%s: function '%s' cannot have more than 255 parameters",
		commaToken.Location, functionIdentifier.Lexeme,
	)
}

func (p *Parser) newErrorFunctionCallTooManyArguments(
	commaToken golox.Token,
) error {
	return fmt.Errorf("%s: function call cannot have more than 255 parameters",
		commaToken.Location,
	)
}
