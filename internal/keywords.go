package golox

var (
	Keywords = map[string]TokenType{
		"var":    TokenTypeVar,
		"nil":    TokenTypeNil,
		"true":   TokenTypeTrue,
		"false":  TokenTypeFalse,
		"and":    TokenTypeAnd,
		"or":     TokenTypeOr,
		"if":     TokenTypeIf,
		"else":   TokenTypeElse,
		"for":    TokenTypeFor,
		"while":  TokenTypeWhile,
		"fun":    TokenTypeFun,
		"return": TokenTypeReturn,
		"class":  TokenTypeClass,
		"super":  TokenTypeSuper,
		"this":   TokenTypeThis,
		"print":  TokenTypePrint,
	}
)
