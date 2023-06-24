package golox

type Token struct {
	Location
	TokenType
	LiteralValue any
	Lexeme       string
}

//go:generate stringer -type=TokenType
type TokenType int

const (
	TokenTypeUndefined TokenType = iota

	// single-character tokens:
	TokenTypeLeftParen
	TokenTypeRightParen
	TokenTypeLeftBrace
	TokenTypeRightBrace
	TokenTypeComma
	TokenTypeDot
	TokenTypeSemicolon
	TokenTypePlus
	TokenTypeMinus
	TokenTypeStar
	TokenTypeSlash

	// one or two character tokens:
	TokenTypeBang
	TokenTypeBangEqual
	TokenTypeEqual
	TokenTypeEqualEqual
	TokenTypeLess
	TokenTypeLessEqual
	TokenTypeGreater
	TokenTypeGreaterEqual

	// literals:
	TokenTypeString
	TokenTypeNumber

	// keywords:
	TokenTypeVar
	TokenTypeNil
	TokenTypeTrue
	TokenTypeFalse
	TokenTypeAnd
	TokenTypeOr
	TokenTypeIf
	TokenTypeElse
	TokenTypeFor
	TokenTypeWhile
	TokenTypeFun
	TokenTypeReturn
	TokenTypeClass
	TokenTypeSuper
	TokenTypeThis
	TokenTypePrint

	TokenTypeIdentifier

	TokenTypeEOF
)
