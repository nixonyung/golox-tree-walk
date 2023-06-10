package interpreter

import (
	"fmt"
	golox "golox/internal"
	"golox/internal/helpers"
	"golox/internal/interpreter/builtins"
	"strconv"
)

type Interpreter struct {
	// configs:
	isDebug bool

	// inputs:
	// stmts []golox.Statement
	resolvedLocalVars map[golox.Expression]int

	// states:
	globals *Scope
	scopes  []*Scope
}

func (itp *Interpreter) currScope() *Scope {
	return itp.scopes[len(itp.scopes)-1]
}

func (itp *Interpreter) nthEnclosingScope(n int) *Scope {
	// assume n is from itp.resolvedLocalVars
	scope := itp.currScope()
	for i := 0; i < n; i++ {
		scope = scope.Enclosing
	}
	return scope
}

func (itp *Interpreter) beginBlockScope() {
	itp.scopes[len(itp.scopes)-1] = &Scope{
		NameToValue: map[string]any{},
		Enclosing:   itp.currScope(),
	}
	itp.logBeginBlockScope()
}

func (itp *Interpreter) endBlockScope() {
	itp.logEndBlockScope()
	itp.scopes[len(itp.scopes)-1] = itp.currScope().Enclosing
}

func (itp *Interpreter) beginFunctionScope(closure *Scope) {
	itp.scopes = append(itp.scopes, &Scope{
		NameToValue: map[string]any{},
		Enclosing:   closure,
	})
	itp.logBeginFunctionScope()
}

func (itp *Interpreter) endFunctionScope() {
	itp.logEndFunctionScope()
	itp.scopes = itp.scopes[:len(itp.scopes)-1]
}

func (itp *Interpreter) defineVar(identifier golox.Token, val any) {
	itp.currScope().NameToValue[identifier.Lexeme] = val
	itp.logDefinedVar(identifier, val)
}

func (itp *Interpreter) assignVar(
	identifier golox.Token,
	val any,
	expr golox.Expression,
) (
	any,
	error,
) {
	if dist, ok := itp.resolvedLocalVars[expr]; ok {
		scope := itp.nthEnclosingScope(dist)

		if _, ok := scope.NameToValue[identifier.Lexeme]; !ok {
			return nil, itp.newErrorUndefinedVariable(identifier)
		} else {
			scope.NameToValue[identifier.Lexeme] = val
			itp.logAssignedVar(identifier, val, scope)
			return val, nil
		}
	} else {
		if _, ok := itp.globals.NameToValue[identifier.Lexeme]; !ok {
			return nil, itp.newErrorUndefinedVariable(identifier)
		} else {
			itp.globals.NameToValue[identifier.Lexeme] = val
			itp.logAssignedVar(identifier, val, itp.globals)
			return val, nil
		}
	}
}

func (itp *Interpreter) getVar(
	identifier golox.Token,
	expr golox.Expression,
) (
	any,
	error,
) {
	if dist, ok := itp.resolvedLocalVars[expr]; ok {
		scope := itp.nthEnclosingScope(dist)

		if val, ok := scope.NameToValue[identifier.Lexeme]; !ok {
			return nil, itp.newErrorUndefinedVariable(identifier)
		} else {
			itp.logGotVar(identifier, val, scope)
			return val, nil
		}
	} else {
		if val, ok := itp.globals.NameToValue[identifier.Lexeme]; !ok {
			return nil, itp.newErrorUndefinedVariable(identifier)
		} else {
			itp.logGotVar(identifier, val, itp.globals)
			return val, nil
		}
	}
}

func (itp *Interpreter) execute(stmt golox.Statement) error {
	switch stmt := stmt.(type) {
	case nil:
		return nil

	case *golox.StatementBlock:
		itp.beginBlockScope()
		for _, stmt := range stmt.Statements {
			if err := itp.execute(stmt); err != nil {
				return err
			}
		}
		itp.endBlockScope()

	case *golox.StatementExpression:
		if val, err := itp.evaluate(stmt.Expression); err != nil {
			return err
		} else {
			itp.logExecutedStatementExpression(stmt, val)
		}

	case *golox.StatementVar:
		if val, err := itp.evaluate(stmt.Expression); err != nil {
			return err
		} else {
			itp.defineVar(stmt.Identifier, val)
			itp.logExecutedStatementVar(stmt, val)
		}

	case *golox.StatementIf:
		if val, err := itp.evaluate(stmt.Condition); err != nil {
			return err
		} else {
			isTrue := helpers.IsValueTruthy(val)
			itp.logEvaluatedStatementIfCondition(stmt, val, isTrue)
			if isTrue {
				if err := itp.execute(stmt.Then); err != nil {
					return err
				}
			} else {
				if err := itp.execute(stmt.Else); err != nil {
					return err
				}
			}
		}

	case *golox.StatementWhile:
		for {
			if val, err := itp.evaluate(stmt.Condition); err != nil {
				return err
			} else {
				isTrue := helpers.IsValueTruthy(val)
				itp.logEvaluatedStatementWhileCondition(stmt, val, isTrue)
				if isTrue {
					if err := itp.execute(stmt.Body); err != nil {
						return err
					}
				} else {
					break
				}
			}
		}

	case *golox.StatementFun:
		itp.defineVar(stmt.Identifier, &LoxFunction{
			Declaration:   stmt,
			IsInitializer: false,
			Closure:       itp.scopes[len(itp.scopes)-1],
			Interpreter:   itp,
		})
		itp.logExecutedStatementFun(stmt)

	case *golox.StatementReturn:
		if val, err := itp.evaluate(stmt.Expression); err != nil {
			return err
		} else {
			itp.logEvaluatedStatementReturnExpression(stmt, val)
			panic(ReturnValue{Value: val}) // caught in utils:functionScope
		}

	case *golox.StatementClass:
		result := &LoxClass{}
		result.Identifier = stmt.Identifier
		result.Methods = map[string]*LoxFunction{}

		closure := itp.currScope()

		if stmt.Superclass != nil {
			if val, err := itp.evaluate(stmt.Superclass); err != nil {
				return err
			} else if sc, ok := val.(*LoxClass); !ok {
				return itp.newErrorInvalidClass(stmt.Superclass)
			} else {
				result.Superclass = sc
				closure = &Scope{
					NameToValue: map[string]any{
						"super": sc,
					},
					Enclosing: closure,
				}
			}
		}

		for _, method := range stmt.Methods {
			result.Methods[method.Identifier.Lexeme] = &LoxFunction{
				Declaration:   method,
				IsInitializer: method.Identifier.Lexeme == "init",
				Closure:       closure,
				Interpreter:   itp,
			}
		}

		itp.defineVar(stmt.Identifier, result)
		itp.logExecutedStatementClass(stmt, result)

	case *golox.StatementPrint:
		if val, err := itp.evaluate(stmt.Expression); err != nil {
			return err
		} else {
			itp.logEvaluatedStatementPrintExpression(stmt, val)

			switch val := val.(type) {
			case nil:
				fmt.Println("<nil>")
			case string:
				fmt.Println("\"" + string(val) + "\"")
			case float64:
				fmt.Println(strconv.FormatFloat(val, 'f', -1, 64))
			default:
				fmt.Println(val)
			}
		}

	default:
		return itp.newErrorMissingImplementation(stmt)
	}

	return nil
}

func (itp *Interpreter) evaluate(expr golox.Expression) (any, error) {
	switch expr := expr.(type) {
	case nil:
		return nil, nil

	case *golox.ExpressionLiteral:
		return expr.LiteralValue, nil

	case *golox.ExpressionGrouping:
		return itp.evaluate(expr.Expression)

	case *golox.ExpressionVariable:
		return itp.getVar(expr.Identifier, expr)

	case *golox.ExpressionCall:
		if val, err := itp.evaluate(expr.Callee); err != nil {
			return nil, err
		} else if callee, ok := val.(LoxCallable); !ok {
			return nil, itp.newErrorInvalidFunctionCallee(expr.Callee)
		} else if len(expr.Arguments) != callee.Arity() {
			return nil, itp.newErrorFunctionArityMismatch(expr, callee.Arity(), len(expr.Arguments))
		} else {
			args := []any{}
			for _, arg := range expr.Arguments {
				if val, err := itp.evaluate(arg); err != nil {
					return nil, err
				} else {
					args = append(args, val)
				}
			}
			return callee.Call(args)
		}

	case *golox.ExpressionGet:
		if val, err := itp.evaluate(expr.Object); err != nil {
			return nil, err
		} else if obj, ok := val.(*LoxInstance); !ok {
			return nil, itp.newErrorInvalidObjectInstance(expr.Object)
		} else {
			return obj.Get(expr.Identifier)
		}

	case *golox.ExpressionSet:
		if objVal, err := itp.evaluate(expr.Object); err != nil {
			return nil, err
		} else if obj, ok := objVal.(*LoxInstance); !ok {
			return nil, itp.newErrorInvalidObjectInstance(expr.Object)
		} else if val, err := itp.evaluate(expr.Value); err != nil {
			return nil, err
		} else {
			obj.Set(expr.Identifier, val)
			return val, nil
		}

	case *golox.ExpressionThis:
		return itp.getVar(expr.ThisToken, expr)

	case *golox.ExpressionSuper:
		if dist, ok := itp.resolvedLocalVars[expr]; !ok {
			return nil, itp.newErrorUndefinedVariable(expr.SuperToken)
		} else {
			superVal := itp.nthEnclosingScope(dist).NameToValue["super"]

			if sc, ok := superVal.(*LoxClass); !ok {
				return nil, itp.newErrorInvalidSuperclassValue(expr, superVal)
			} else if method, err := sc.FindMethod(expr.Method); err != nil {
				return nil, err
			} else {
				objVal := itp.nthEnclosingScope(dist - 1).NameToValue["this"]

				if obj, ok := objVal.(*LoxInstance); !ok {
					return nil, itp.newErrorInvalidThisValue(expr, objVal)
				} else {
					return method.WithThisBoundTo(obj), nil
				}
			}
		}

	case *golox.ExpressionUnary:
		if rhs, err := itp.evaluate(expr.Right); err != nil {
			return nil, err
		} else {
			switch expr.Operator.TokenType {
			case golox.TokenTypeMinus:
				if rhs, ok := rhs.(float64); ok {
					return -rhs, nil
				} else {
					return nil, itp.newErrorOperandMustBe("a number", expr.Operator)
				}
			case golox.TokenTypeBang:
				return !helpers.IsValueTruthy(rhs), nil
			}
		}

	case *golox.ExpressionBinary:
		if lhs, err := itp.evaluate(expr.Left); err != nil {
			return nil, err
		} else if rhs, err := itp.evaluate(expr.Right); err != nil {
			return nil, err
		} else {
			switch expr.Operator.TokenType {
			case golox.TokenTypeGreater:
				if lhs, rhs, ok := helpers.AssertTwo[float64, float64](lhs, rhs); ok {
					return lhs > rhs, nil
				} else {
					return nil, itp.newErrorOperandsMustBe("both numbers", expr.Operator)
				}
			case golox.TokenTypeGreaterEqual:
				if lhs, rhs, ok := helpers.AssertTwo[float64, float64](lhs, rhs); ok {
					return lhs >= rhs, nil
				} else {
					return nil, itp.newErrorOperandsMustBe("both numbers", expr.Operator)
				}
			case golox.TokenTypeLess:
				if lhs, rhs, ok := helpers.AssertTwo[float64, float64](lhs, rhs); ok {
					return lhs < rhs, nil
				} else {
					return nil, itp.newErrorOperandsMustBe("both numbers", expr.Operator)
				}
			case golox.TokenTypeLessEqual:
				if lhs, rhs, ok := helpers.AssertTwo[float64, float64](lhs, rhs); ok {
					return lhs <= rhs, nil
				} else {
					return nil, itp.newErrorOperandsMustBe("both numbers", expr.Operator)
				}
			case golox.TokenTypeBangEqual:
				return lhs != rhs, nil
			case golox.TokenTypeEqualEqual:
				return lhs == rhs, nil
			case golox.TokenTypeMinus:
				if lhs, rhs, ok := helpers.AssertTwo[float64, float64](lhs, rhs); ok {
					return lhs - rhs, nil
				} else {
					return nil, itp.newErrorOperandsMustBe("both numbers", expr.Operator)
				}
			case golox.TokenTypePlus:
				if lhs, rhs, ok := helpers.AssertTwo[float64, float64](lhs, rhs); ok {
					return lhs + rhs, nil
				}
				if lhs, rhs, ok := helpers.AssertTwo[string, string](lhs, rhs); ok {
					return lhs + rhs, nil
				}
				return nil, itp.newErrorOperandsMustBe("both numbers or both strings", expr.Operator)

			case golox.TokenTypeSlash:
				if lhs, rhs, ok := helpers.AssertTwo[float64, float64](lhs, rhs); ok {
					return lhs / rhs, nil
				} else {
					return nil, itp.newErrorOperandsMustBe("both numbers", expr.Operator)
				}
			case golox.TokenTypeStar:
				if lhs, rhs, ok := helpers.AssertTwo[float64, float64](lhs, rhs); ok {
					return lhs * rhs, nil
				} else {
					return nil, itp.newErrorOperandsMustBe("both numbers", expr.Operator)
				}
			}
		}

	case *golox.ExpressionLogical:
		if lhs, err := itp.evaluate(expr.Left); err != nil {
			return nil, err
		} else if expr.Operator.TokenType == golox.TokenTypeOr && helpers.IsValueTruthy(lhs) {
			return lhs, nil
		} else if expr.Operator.TokenType == golox.TokenTypeAnd && !helpers.IsValueTruthy(lhs) {
			return lhs, nil
		} else if rhs, err := itp.evaluate(expr.Right); err != nil {
			return nil, err
		} else {
			return rhs, nil
		}

	case *golox.ExpressionAssignment:
		if val, err := itp.evaluate(expr.Value); err != nil {
			return nil, err
		} else {
			return itp.assignVar(expr.Identifier, val, expr)
		}
	}

	return nil, itp.newErrorMissingImplementation(expr)
}

func (itp *Interpreter) InterpretStatements(
	stmts []golox.Statement,
	resolvedLocalVars map[golox.Expression]int,
) error {
	itp.resolvedLocalVars = resolvedLocalVars
	itp.logGlobalScope()

	for _, stmt := range stmts {
		if err := itp.execute(stmt); err != nil {
			return err
		}
	}
	return nil
}

func NewInterpreter(
	isDebug bool,
) *Interpreter {
	globals := &Scope{
		NameToValue: map[string]any{
			"clock": &builtins.Clock{},
		},
		Enclosing: &Scope{},
	}

	return &Interpreter{
		isDebug:           isDebug,
		resolvedLocalVars: nil,
		globals:           globals,
		scopes:            []*Scope{globals},
	}
}
