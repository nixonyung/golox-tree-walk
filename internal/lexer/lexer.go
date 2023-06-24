package lexer

import (
	golox "golox/internal"
	"strconv"
)

type Lexer struct {
	source          []rune        // input
	srcPath         string        // input
	tokens          []golox.Token // output
	curr, line, col int
}

func (l *Lexer) lookAhead(k int) (rune, bool) {
	if l.curr+k >= len(l.source) {
		return 0, false
	} else {
		return l.source[l.curr+k], true
	}
}

func (l *Lexer) advance(k int) {
	l.curr += k
	l.col += k
}

func (l *Lexer) lexeme(k int) string {
	if l.curr+k > len(l.source) {
		return ""
	} else {
		return string(l.source[l.curr : l.curr+k])
	}
}

func (l *Lexer) newline() {
	l.line++
	l.col = 1
}

func (l *Lexer) consumeAsToken(k int, tokenType golox.TokenType, literalValue any) {
	tkn := golox.Token{
		Location: golox.Location{
			SrcPath: l.srcPath,
			Line:    l.line,
			Col:     l.col,
		},
		TokenType:    tokenType,
		LiteralValue: literalValue,
		Lexeme:       l.lexeme(k),
	}
	l.tokens = append(l.tokens, tkn)
	golox.Logf(
		golox.ModuleLexer,
		"%s:%d:%d: '%s' (%s)",
		l.srcPath, l.line, l.col, tkn.Lexeme, tkn.TokenType,
	)
	l.advance(k)
}

func isCharAlphabet(ch rune) bool {
	return ('A' <= ch && ch <= 'Z') || ('a' <= ch && ch <= 'z') || ch == '_'
}

func isCharNumeric(ch rune) bool {
	return ('0' <= ch && ch <= '9')
}

func (l *Lexer) TokensFromSource(source []rune, srcPath string) ([]golox.Token, error) {
	l.source = source
	l.srcPath = srcPath
	l.tokens = []golox.Token{}
	l.curr, l.line, l.col = 0, 1, 1

	for {
		if ch, ok := l.lookAhead(0); !ok {
			break
		} else {
			switch ch {
			case '\n':
				l.advance(1)
				l.newline()
			case ' ', '\r', '\t':
				l.advance(1)
			case '(':
				l.consumeAsToken(1, golox.TokenTypeLeftParen, nil)
			case ')':
				l.consumeAsToken(1, golox.TokenTypeRightParen, nil)
			case '{':
				l.consumeAsToken(1, golox.TokenTypeLeftBrace, nil)
			case '}':
				l.consumeAsToken(1, golox.TokenTypeRightBrace, nil)
			case ',':
				l.consumeAsToken(1, golox.TokenTypeComma, nil)
			case '.':
				l.consumeAsToken(1, golox.TokenTypeDot, nil)
			case ';':
				l.consumeAsToken(1, golox.TokenTypeSemicolon, nil)
			case '+':
				l.consumeAsToken(1, golox.TokenTypePlus, nil)
			case '-':
				l.consumeAsToken(1, golox.TokenTypeMinus, nil)
			case '*':
				l.consumeAsToken(1, golox.TokenTypeStar, nil)
			case '/':
				if ch, ok := l.lookAhead(1); ok && ch == '/' {
					l.advance(2)
					for {
						if ch, ok := l.lookAhead(0); !ok || ch == '\n' {
							break
						} else {
							l.advance(1)
						}
					}
					l.advance(1)
					l.newline()
				} else {
					l.consumeAsToken(1, golox.TokenTypeSlash, nil)
				}
			case '!':
				if ch, ok := l.lookAhead(1); ok && ch == '=' {
					l.consumeAsToken(2, golox.TokenTypeBangEqual, nil)
				} else {
					l.consumeAsToken(1, golox.TokenTypeBang, nil)
				}
			case '=':
				if ch, ok := l.lookAhead(1); ok && ch == '=' {
					l.consumeAsToken(2, golox.TokenTypeEqualEqual, nil)
				} else {
					l.consumeAsToken(1, golox.TokenTypeEqual, nil)
				}
			case '<':
				if ch, ok := l.lookAhead(1); ok && ch == '=' {
					l.consumeAsToken(2, golox.TokenTypeLessEqual, nil)
				} else {
					l.consumeAsToken(1, golox.TokenTypeLess, nil)
				}
			case '>':
				if ch, ok := l.lookAhead(1); ok && ch == '=' {
					l.consumeAsToken(2, golox.TokenTypeGreaterEqual, nil)
				} else {
					l.consumeAsToken(1, golox.TokenTypeGreater, nil)
				}
			case '"':
				leftQuoteLine, leftQuoteCol := l.line, l.col
				k := 1
				for ; ; k++ {
					if ch, ok := l.lookAhead(k); !ok {
						return nil, golox.NewErrorf(
							golox.Location{
								SrcPath: l.srcPath,
								Line:    l.line,
								Col:     l.col,
							},
							"unterminated string, started at line %d:%d",
							leftQuoteLine, leftQuoteCol)
					} else if ch == '"' {
						break
					} else if ch == '\n' {
						l.newline()
					}
				}
				l.consumeAsToken(k+1, golox.TokenTypeString, string(l.lexeme(k)[1:]))
			default:
				switch {
				case isCharNumeric(ch):
					k := 1
					for ; ; k++ {
						if ch, ok := l.lookAhead(k); !ok {
							break
						} else if !isCharNumeric(ch) {
							if ch != '.' {
								// end of number if not numeric and not == '.'
								break
							} else {
								if ch2, ok2 := l.lookAhead(k + 1); !ok2 || !isCharNumeric(ch2) {
									// char after '.' is not numeric, so the '.' is a method call, should break
									break
								}
								// else ch is the decimal point, should continue
							}
						}
					}
					if val, err := strconv.ParseFloat(l.lexeme(k), 64); err != nil {
						return nil, golox.NewErrorf(
							golox.Location{
								SrcPath: l.srcPath,
								Line:    l.line,
								Col:     l.col,
							},
							"invalid numeric string: %s", err.Error(),
						)
					} else {
						l.consumeAsToken(k, golox.TokenTypeNumber, val)
					}
				case isCharAlphabet(ch):
					k := 1
					for ; ; k++ {
						if ch, ok := l.lookAhead(k); !ok || (!isCharAlphabet(ch) && !isCharNumeric(ch)) {
							break
						}
					}
					if tokenType, isKeyword := golox.Keywords[l.lexeme(k)]; isKeyword {
						l.consumeAsToken(k, tokenType, nil)
					} else {
						l.consumeAsToken(k, golox.TokenTypeIdentifier, nil)
					}
				default:
					return nil, golox.NewErrorf(
						golox.Location{
							SrcPath: l.srcPath,
							Line:    l.line,
							Col:     l.col,
						},
						"unexpected character '%c'", ch,
					)
				}
			}
		}
	}
	l.consumeAsToken(0, golox.TokenTypeEOF, nil)
	return l.tokens, nil
}

func NewLexer() *Lexer {
	return &Lexer{}
}
