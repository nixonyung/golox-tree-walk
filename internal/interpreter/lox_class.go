package interpreter

import golox "golox/internal"

type LoxClass struct {
	Identifier golox.Token
	Superclass *LoxClass
	Methods    map[string]*LoxFunction
}

// implements Callable
func (c *LoxClass) String() string {
	return "<class: " + c.Identifier.Lexeme + ">"
}

// implements Callable
func (c *LoxClass) Arity() int {
	if initMethod, ok := c.FindInit(); ok {
		return initMethod.Arity()
	} else {
		return 0
	}
}

// implements Callable
func (c *LoxClass) Call(args []any) (any, error) {
	ins := &LoxInstance{
		Class:  c,
		Fields: map[string]any{},
	}

	if initMethod, ok := c.FindInit(); ok {
		initMethod.WithThisBoundTo(ins).Call(args)
	}
	return ins, nil
}

func (c *LoxClass) FindMethod(identifier golox.Token) (*LoxFunction, error) {
	if method, ok := c.Methods[identifier.Lexeme]; ok {
		return method, nil
	} else if c.Superclass != nil {
		return c.Superclass.FindMethod(identifier)
	} else {
		return nil, c.newErrorUndefinedProperty(identifier)
	}
}

func (c *LoxClass) FindInit() (*LoxFunction, bool) {
	if initMethod, ok := c.Methods["init"]; ok {
		return initMethod, true
	} else if c.Superclass != nil {
		return c.Superclass.FindInit()
	} else {
		return nil, false
	}
}
