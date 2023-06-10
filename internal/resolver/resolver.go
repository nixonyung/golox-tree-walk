package resolver

import golox "golox/internal"

type Resolver struct {
	// configs:
	isDebug bool

	// inputs:
	// stmts []golox.Statement

	// outputs:
	resolvedLocalVars map[golox.Expression]int

	// states:
	scopes           []map[string]bool
	currFunctionType FunctionType
	currClassType    ClassType
}

func (r *Resolver) currScope() (map[string]bool, bool) {
	if len(r.scopes) > 0 {
		return r.scopes[len(r.scopes)-1], true
	} else {
		return nil, false
	}
}

func (r *Resolver) beginScope() {
	r.scopes = append(r.scopes, map[string]bool{})
}

func (r *Resolver) endScope() {
	r.scopes = r.scopes[:len(r.scopes)-1]
}

func (r *Resolver) declareVarInCurrScope(identifier golox.Token) error {
	if currScope, ok := r.currScope(); !ok {
		return nil
	} else if _, ok := currScope[identifier.Lexeme]; ok {
		return r.newErrorVariableIsAlreadyDefined(identifier)
	} else {
		currScope[identifier.Lexeme] = false
		r.logDeclaredVariableInCurrScope(identifier)
		return nil
	}
}

func (r *Resolver) defineVarInCurrScope(identifier golox.Token) {
	if currScope, ok := r.currScope(); ok {
		currScope[identifier.Lexeme] = true
		r.logDefinedVariableInCurrScope(identifier)
	}
}

func (r *Resolver) defineThisInCurrScope(classIdentifier golox.Token) {
	if currScope, ok := r.currScope(); ok {
		currScope["this"] = true
		r.logDefinedThisInCurrScope(classIdentifier)
	}
}

func (r *Resolver) defineSuperInCurrScope(classIdentifier golox.Token) {
	if currScope, ok := r.currScope(); ok {
		currScope["super"] = true
		r.logDefinedSuperInCurrScope(classIdentifier)
	}
}

func (r *Resolver) isVarDeclaredInScope(
	identifier golox.Token,
	scope map[string]bool,
) bool {
	_, ok := scope[identifier.Lexeme]
	return ok
}

func (r *Resolver) isVarDefinedInScope(
	identifier golox.Token,
	scope map[string]bool,
) bool {
	return scope[identifier.Lexeme]
}

func (r *Resolver) resolveVariable(
	identifier golox.Token,
	expr golox.Expression,
) {
	for i := len(r.scopes) - 1; i >= 0; i-- {
		if r.isVarDeclaredInScope(identifier, r.scopes[i]) {
			dist := len(r.scopes) - 1 - i
			r.resolvedLocalVars[expr] = dist
			r.logResolvedVariable(identifier, expr, dist)
			return
		}
	}
}

func (r *Resolver) resolveFunction(
	functionType FunctionType,
	stmt *golox.StatementFun,
) error {
	lastFunctionType := r.currFunctionType
	r.currFunctionType = functionType

	r.beginScope()
	for _, param := range stmt.Parameters {
		if err := r.declareVarInCurrScope(param); err != nil {
			return err
		}
		r.defineVarInCurrScope(param)
	}
	for _, stmt := range stmt.Body {
		if err := r.resolveStatement(stmt); err != nil {
			return err
		}
	}
	r.endScope()
	r.currFunctionType = lastFunctionType
	return nil
}

func (r *Resolver) resolveExpression(expr golox.Expression) error {
	switch expr := expr.(type) {
	case nil:
		break
	case *golox.ExpressionLiteral:
		break
	case *golox.ExpressionGrouping:
		return r.resolveExpression(expr.Expression)
	case *golox.ExpressionVariable:
		if currScope, ok := r.currScope(); ok &&
			r.isVarDeclaredInScope(expr.Identifier, currScope) &&
			!r.isVarDefinedInScope(expr.Identifier, currScope) {
			return r.newErrorVariableInItsOwnInitializer(expr.Identifier)
		} else {
			r.resolveVariable(expr.Identifier, expr)
			return nil
		}
	case *golox.ExpressionCall:
		if err := r.resolveExpression(expr.Callee); err != nil {
			return err
		}
		for _, arg := range expr.Arguments {
			if err := r.resolveExpression(arg); err != nil {
				return err
			}
		}
	case *golox.ExpressionGet:
		return r.resolveExpression(expr.Object)
	case *golox.ExpressionSet:
		if err := r.resolveExpression(expr.Object); err != nil {
			return err
		}
		if err := r.resolveExpression(expr.Value); err != nil {
			return err
		}
	case *golox.ExpressionThis:
		switch r.currClassType {
		case ClassTypeNone:
			return r.newErrorTopLevelThis(expr.ThisToken)
		default:
			r.resolveVariable(expr.ThisToken, expr)
			return nil
		}
	case *golox.ExpressionSuper:
		switch r.currClassType {
		case ClassTypeNone:
			return r.newErrorSuperOutsideClass(expr.SuperToken)
		case ClassTypeSubclass:
			r.resolveVariable(expr.SuperToken, expr)
			return nil
		default:
			return r.newErrorSuperWithoutSuperclass(expr.SuperToken)
		}
	case *golox.ExpressionUnary:
		return r.resolveExpression(expr.Right)
	case *golox.ExpressionBinary:
		if err := r.resolveExpression(expr.Left); err != nil {
			return err
		}
		if err := r.resolveExpression(expr.Right); err != nil {
			return err
		}
	case *golox.ExpressionLogical:
		if err := r.resolveExpression(expr.Left); err != nil {
			return err
		}
		if err := r.resolveExpression(expr.Right); err != nil {
			return err
		}
	case *golox.ExpressionAssignment:
		if err := r.resolveExpression(expr.Value); err != nil {
			return err
		}
		r.resolveVariable(expr.Identifier, expr)
		return nil
	default:
		return r.newErrorMissingImplementation(expr)
	}
	return nil
}

func (r *Resolver) resolveStatement(stmt golox.Statement) error {
	switch stmt := stmt.(type) {
	case nil:
		break
	case *golox.StatementExpression:
		return r.resolveExpression(stmt.Expression)
	case *golox.StatementBlock:
		r.beginScope()
		for _, stmt := range stmt.Statements {
			if err := r.resolveStatement(stmt); err != nil {
				return err
			}
		}
		r.endScope()
	case *golox.StatementVar:
		if err := r.declareVarInCurrScope(stmt.Identifier); err != nil {
			return err
		}
		if err := r.resolveExpression(stmt.Expression); err != nil {
			return err
		}
		// definition of the variable should not be considered before resolving
		// the expression
		r.defineVarInCurrScope(stmt.Identifier)
		return nil
	case *golox.StatementFun:
		if err := r.declareVarInCurrScope(stmt.Identifier); err != nil {
			return err
		}

		r.defineVarInCurrScope(stmt.Identifier)

		if err := r.resolveFunction(FunctionTypeFunction, stmt); err != nil {
			return err
		} else {
			return nil
		}
	case *golox.StatementReturn:
		switch r.currFunctionType {
		case FunctionTypeNone:
			return r.newErrorTopLevelReturn(stmt.ReturnToken)
		case FunctionTypeInitializer:
			if stmt.Expression != nil {
				return r.newErrorReturnWithValueInInitializer(stmt.ReturnToken)
			} else {
				return nil
			}
		default:
			return r.resolveExpression(stmt.Expression)
		}
	case *golox.StatementClass:
		if err := r.declareVarInCurrScope(stmt.Identifier); err != nil {
			return err
		} else {
			r.defineVarInCurrScope(stmt.Identifier)

			lastClassType := r.currClassType
			if stmt.Superclass != nil {
				if stmt.Superclass.Identifier.Lexeme == stmt.Identifier.Lexeme {
					return r.newErrorClassInheritFromItself(stmt.Identifier, stmt.Superclass.Identifier)
				} else {
					r.currClassType = ClassTypeSubclass
					if err := r.resolveExpression(stmt.Superclass); err != nil {
						return err
					}
					r.beginScope()
					r.defineSuperInCurrScope(stmt.Identifier)
				}
			} else {
				r.currClassType = ClassTypeClass
			}

			r.beginScope()
			r.defineThisInCurrScope(stmt.Identifier)
			for _, method := range stmt.Methods {
				if method.Identifier.Lexeme == "init" {
					if err := r.resolveFunction(FunctionTypeInitializer, method); err != nil {
						return err
					}
				} else {
					if err := r.resolveFunction(FunctionTypeMethod, method); err != nil {
						return err
					}
				}
			}
			r.endScope()

			if stmt.Superclass != nil {
				r.endScope()
			}
			r.currClassType = lastClassType
		}
	case *golox.StatementPrint:
		return r.resolveExpression(stmt.Expression)
	case *golox.StatementIf:
		if err := r.resolveExpression(stmt.Condition); err != nil {
			return err
		}
		if err := r.resolveStatement(stmt.Then); err != nil {
			return err
		}
		if err := r.resolveStatement(stmt.Else); err != nil {
			return err
		}
	case *golox.StatementWhile:
		if err := r.resolveExpression(stmt.Condition); err != nil {
			return err
		}
		if err := r.resolveStatement(stmt.Body); err != nil {
			return err
		}
	default:
		return r.newErrorMissingImplementation(stmt)
	}
	return nil
}

func (r *Resolver) ResolveStatements(
	stmts []golox.Statement,
) (
	map[golox.Expression]int,
	error,
) {
	r.resolvedLocalVars = map[golox.Expression]int{}
	r.scopes = []map[string]bool{}
	r.currFunctionType = FunctionTypeNone
	r.currClassType = ClassTypeNone

	for _, stmt := range stmts {
		if err := r.resolveStatement(stmt); err != nil {
			return nil, err
		}
	}
	return r.resolvedLocalVars, nil
}

func NewResolver(
	isDebug bool,
) *Resolver {
	return &Resolver{
		isDebug: isDebug,
	}
}
