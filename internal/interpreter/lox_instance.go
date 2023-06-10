package interpreter

import golox "golox/internal"

type LoxInstance struct {
	Class  *LoxClass
	Fields map[string]any
}

func (ins *LoxInstance) String() string {
	return "<instance of " + ins.Class.String() + ">"
}

func (ins *LoxInstance) Get(identifier golox.Token) (any, error) {
	if val, ok := ins.Fields[identifier.Lexeme]; ok {
		return val, nil
	} else if method, err := ins.Class.FindMethod(identifier); err != nil {
		return nil, err
	} else {
		return method.WithThisBoundTo(ins), nil
	}
}

func (ins *LoxInstance) Set(identifier golox.Token, value any) {
	ins.Fields[identifier.Lexeme] = value
}
