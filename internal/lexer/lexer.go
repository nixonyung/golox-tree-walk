package lexer

import (
	golox "golox/internal"
	"golox/internal/helpers"
	"strconv"
)

type Lexer struct {
	// configs:
	isDebug bool
	srcPath string

	// inputs:
	source []rune

	// outputs:
	tokens []golox.Token

	// states:
	curr, currLine, currCol int
	prev, prevLine, prevCol int
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
	l.currCol += k
}

func (l *Lexer) lexeme(k int) string {
	if l.curr+k > len(l.source) {
		return ""
	} else {
		return string(l.source[l.curr : l.curr+k])
	}
}

func (l *Lexer) newline() {
	l.currLine++
	l.currCol = 1
}

func (l *Lexer) resetStart() {
	l.prev = l.curr
	l.prevLine = l.currLine
	l.prevCol = l.currCol
}

func (l *Lexer) consumeAdvanceReset(k int, tokenType golox.TokenType, literalValue any) {
	newToken := golox.Token{
		Location: golox.Location{
			SrcPath: l.srcPath,
			Line:    l.prevLine,
			Col:     l.prevCol,
		},
		TokenType:    tokenType,
		LiteralValue: literalValue,
		Lexeme:       l.lexeme(k),
	}
	l.tokens = append(l.tokens, newToken)
	l.logAddedToken(newToken)
	l.advance(k)
	l.resetStart()
}

func (l *Lexer) TokensFromSource(source []rune) ([]golox.Token, error) {
	l.source = source
	l.tokens = []golox.Token{}
	l.curr, l.currLine, l.currCol = 0, 1, 1
	l.prev, l.prevLine, l.prevCol = 0, 1, 1

	for {
		if ch, ok := l.lookAhead(0); !ok {
			break
		} else {
			switch ch {
			case '\n':
				l.advance(1)
				l.newline()
				l.resetStart()
			case ' ', '\r', '\t':
				l.advance(1)
				l.resetStart()
			case '(':
				l.consumeAdvanceReset(1, golox.TokenTypeLeftParen, nil)
			case ')':
				l.consumeAdvanceReset(1, golox.TokenTypeRightParen, nil)
			case '{':
				l.consumeAdvanceReset(1, golox.TokenTypeLeftBrace, nil)
			case '}':
				l.consumeAdvanceReset(1, golox.TokenTypeRightBrace, nil)
			case ',':
				l.consumeAdvanceReset(1, golox.TokenTypeComma, nil)
			case '.':
				l.consumeAdvanceReset(1, golox.TokenTypeDot, nil)
			case ';':
				l.consumeAdvanceReset(1, golox.TokenTypeSemicolon, nil)
			case '+':
				l.consumeAdvanceReset(1, golox.TokenTypePlus, nil)
			case '-':
				l.consumeAdvanceReset(1, golox.TokenTypeMinus, nil)
			case '*':
				l.consumeAdvanceReset(1, golox.TokenTypeStar, nil)
			case '/':
				if ch, ok := l.lookAhead(1); ok && ch == '/' {
					k := 2
					for ; ; k++ {
						if ch, ok := l.lookAhead(k); !ok || ch == '\n' {
							break
						}
					}
					l.advance(k + 1)
					l.newline()
					l.resetStart()
				} else {
					l.consumeAdvanceReset(1, golox.TokenTypeSlash, nil)
				}
			case '!':
				if ch, ok := l.lookAhead(1); ok && ch == '=' {
					l.consumeAdvanceReset(2, golox.TokenTypeBangEqual, nil)
				} else {
					l.consumeAdvanceReset(1, golox.TokenTypeBang, nil)
				}
			case '=':
				if ch, ok := l.lookAhead(1); ok && ch == '=' {
					l.consumeAdvanceReset(2, golox.TokenTypeEqualEqual, nil)
				} else {
					l.consumeAdvanceReset(1, golox.TokenTypeEqual, nil)
				}
			case '<':
				if ch, ok := l.lookAhead(1); ok && ch == '=' {
					l.consumeAdvanceReset(2, golox.TokenTypeLessEqual, nil)
				} else {
					l.consumeAdvanceReset(1, golox.TokenTypeLess, nil)
				}
			case '>':
				if ch, ok := l.lookAhead(1); ok && ch == '=' {
					l.consumeAdvanceReset(2, golox.TokenTypeGreaterEqual, nil)
				} else {
					l.consumeAdvanceReset(1, golox.TokenTypeGreater, nil)
				}
			case '"':
				k := 1
				for ; ; k++ {
					if ch, ok := l.lookAhead(k); !ok {
						return nil, l.newErrorUnterminatedString()
					} else if ch == '"' {
						break
					} else if ch == '\n' {
						l.newline()
					}
				}
				l.consumeAdvanceReset(k+1, golox.TokenTypeString, string(l.lexeme(k)[1:]))
			default:
				switch {
				case helpers.IsCharNumeric(ch):
					k := 1
					for ; ; k++ {
						if ch, ok := l.lookAhead(k); !ok {
							break
						} else if !helpers.IsCharNumeric(ch) {
							if ch != '.' {
								// end of number if not numeric and not == '.'
								break
							} else {
								if ch2, ok2 := l.lookAhead(k + 1); !ok2 || !helpers.IsCharNumeric(ch2) {
									// char after '.' is not numeric, so the '.' is a method call, should break
									break
								}
								// else ch is the decimal point, should continue
							}
						}
					}
					if val, err := strconv.ParseFloat(l.lexeme(k), 64); err != nil {
						return nil, l.newErrorInvalidNumericString(err)
					} else {
						l.consumeAdvanceReset(k, golox.TokenTypeNumber, val)
					}
				case helpers.IsCharAlphabet(ch) || ch == '_':
					k := 1
					for ; ; k++ {
						if ch, ok := l.lookAhead(k); !ok ||
							(ch != '_' &&
								!helpers.IsCharAlphabet(ch) &&
								!helpers.IsCharNumeric(ch)) {
							break
						}
					}
					if tokenType, isKeyword := golox.Keywords[l.lexeme(k)]; isKeyword {
						l.consumeAdvanceReset(k, tokenType, nil)
					} else {
						l.consumeAdvanceReset(k, golox.TokenTypeIdentifier, nil)
					}
				default:
					return nil, l.newErrorUnexpectedCharacter(ch)
				}
			}
		}
	}
	l.consumeAdvanceReset(0, golox.TokenTypeEOF, nil)
	return l.tokens, nil
}

func NewLexer(
	isDebug bool,
	srcPath string,
) *Lexer {
	return &Lexer{
		isDebug: isDebug,
		srcPath: srcPath,
	}
}
