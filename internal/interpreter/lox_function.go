package interpreter

import golox "golox/internal"

type LoxFunction struct {
	Declaration   *golox.StatementFun
	IsInitializer bool
	Closure       *Scope
	Interpreter   *Interpreter
}

type ReturnValue struct {
	Value any
}

func (fn *LoxFunction) String() string {
	return "<fn: " + fn.Declaration.Identifier.Lexeme + ">"
}

func (fn *LoxFunction) Arity() int {
	return len(fn.Declaration.Parameters)
}

func (fn *LoxFunction) Call(args []any) (returnValue any, returnErr error) {
	defer func() {
		// implement returning from a function using exception
		if rv, ok := recover().(ReturnValue); ok {
			fn.Interpreter.endFunctionScope()

			// force init() to return 'this'
			// note that resolver should have returned error if the return statement has a value
			if fn.IsInitializer {
				returnValue = fn.Closure.NameToValue["this"]
			} else {
				returnValue = rv.Value
			}
			returnErr = nil
		}
	}()

	fn.Interpreter.beginFunctionScope(fn.Closure)
	for i, param := range fn.Declaration.Parameters {
		fn.Interpreter.defineVar(param, args[i])
	}

	// panic when there is a return statement with value
	for _, stmt := range fn.Declaration.Body {
		if err := fn.Interpreter.execute(stmt); err != nil {
			return nil, err
		}
	}
	fn.Interpreter.endFunctionScope()

	// force init() to return 'this'
	if fn.IsInitializer {
		return fn.Closure.NameToValue["this"], nil
	} else {
		return nil, nil
	}
}

func (fn *LoxFunction) WithThisBoundTo(ins *LoxInstance) *LoxFunction {
	return &LoxFunction{
		Declaration:   fn.Declaration,
		IsInitializer: fn.IsInitializer,
		Closure: &Scope{
			Enclosing: fn.Closure,
			NameToValue: map[string]any{
				"this": ins,
			},
		},
		Interpreter: fn.Interpreter,
	}
}
