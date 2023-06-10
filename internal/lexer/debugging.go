package lexer

import (
	"fmt"
	golox "golox/internal"
	"log"
)

const (
	log_prefix = "DEBUG :: lexer"
)

func (l *Lexer) lineInfo() string {
	return fmt.Sprintf("%s:%d:%d",
		l.srcPath, l.currLine, l.currCol,
	)
}

func (l *Lexer) logAddedToken(
	tkn golox.Token,
) {
	if l.isDebug {
		log.Printf("%s: %s: added %s('%s')",
			log_prefix, l.lineInfo(), tkn.TokenType, tkn.Lexeme,
		)
	}
}

func (l *Lexer) newErrorUnterminatedString() error {
	return fmt.Errorf("%s: unterminated string, started at line %d:%d",
		l.lineInfo(), l.prevLine, l.prevCol,
	)
}

func (l *Lexer) newErrorInvalidNumericString(
	parsingErr error,
) error {
	return fmt.Errorf("%s: invalid numeric string: %w",
		l.lineInfo(), parsingErr,
	)
}

func (l *Lexer) newErrorUnexpectedCharacter(
	ch rune,
) error {
	return fmt.Errorf("%s: unexpected character '%c'",
		l.lineInfo(), ch,
	)
}
