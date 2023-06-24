package interpreter

import (
	"fmt"
	golox "golox/internal"
	"log"
)

const (
	log_prefix = "DEBUG :: interpreter"
)

func (itp *Interpreter) logDefinedVar(
	identifier golox.Token,
	val any,
) {
	if itp.isDebug {
		log.Printf("%s: %s: defined var '%s' = %v at scope index %d, level %d, currently %s",
			log_prefix, identifier.Location, identifier.Lexeme, val, len(itp.scopes), itp.currScope().level(), itp.currScope().string(),
		)
	}
}

func (itp *Interpreter) logAssignedVar(
	identifier golox.Token,
	val any,
	scope *Scope,
) {
	if itp.isDebug {
		log.Printf("%s: %s: assigned var '%s' = %v at scope index %d, level %d, currently %s",
			log_prefix, identifier.Location, identifier.Lexeme, val, len(itp.scopes), scope.level(), scope.string(),
		)
	}
}

func (itp *Interpreter) logGotVar(
	identifier golox.Token,
	val any,
	scope *Scope,
) {
	if itp.isDebug {
		log.Printf("%s: %s: got var '%s' = %v at scope index %d, level %d, currently %s",
			log_prefix, identifier.Location, identifier.Lexeme, val, len(itp.scopes), scope.level(), scope.string(),
		)
	}
}

func (itp *Interpreter) logGlobalScope() {
	if itp.isDebug {
		log.Printf("%s: global scope starts, currently has %s",
			log_prefix, itp.globals.string(),
		)
	}
}

func (itp *Interpreter) logBeginBlockScope() {
	if itp.isDebug {
		log.Printf("%s: scope index %d, level %d starts",
			log_prefix, len(itp.scopes), itp.currScope().level(),
		)
	}
}

func (itp *Interpreter) logEndBlockScope() {
	if itp.isDebug {
		log.Printf("%s: scope index %d, level %d ends",
			log_prefix, len(itp.scopes), itp.currScope().level(),
		)
	}
}

func (itp *Interpreter) logBeginFunctionScope() {
	if itp.isDebug {
		log.Printf("%s: scope index %d starts at level %d",
			log_prefix, len(itp.scopes), itp.currScope().level(),
		)
	}
}

func (itp *Interpreter) logEndFunctionScope() {
	if itp.isDebug {
		log.Printf("%s: scope index %d ends at level %d",
			log_prefix, len(itp.scopes), itp.currScope().level(),
		)
	}
}

func (itp *Interpreter) logExecutedStatementExpression(
	stmt *golox.StatementExpression,
	val any,
) {
	if itp.isDebug {
		log.Printf("%s: %s: executed StatementExpression: evaluated %s = %v",
			log_prefix, stmt.GetLocation(), stmt.Expression, val,
		)
	}
}

func (itp *Interpreter) logExecutedStatementVar(
	stmt *golox.StatementVar,
	val any,
) {
	if itp.isDebug {
		log.Printf("%s: %s: executed StatementVar: defined variable %s = %v",
			log_prefix, stmt.GetLocation(), stmt.Identifier.Lexeme, val,
		)
	}
}

func (itp *Interpreter) logEvaluatedStatementIfCondition(
	stmt *golox.StatementIf,
	val any,
	isTrue bool,
) {
	if itp.isDebug {
		log.Printf("%s: %s: in StatementIf: evaluated if condition %s = %v, which is %t",
			log_prefix, stmt.GetLocation(), stmt.Condition, val, isTrue,
		)
	}
}

func (itp *Interpreter) logEvaluatedStatementWhileCondition(
	stmt *golox.StatementWhile,
	val any,
	isTrue bool,
) {
	if itp.isDebug {
		log.Printf("%s: %s: in StatementWhile: evaluated while condition %s = %v, which is %t",
			log_prefix, stmt.GetLocation(), stmt.Condition, val, isTrue,
		)
	}
}

func (itp *Interpreter) logExecutedStatementFun(
	stmt *golox.StatementFun,
) {
	if itp.isDebug {
		log.Printf("%s: %s: executed StatementFun: defined function %s",
			log_prefix, stmt.GetLocation(), stmt.Identifier.Lexeme,
		)
	}
}

func (itp *Interpreter) logEvaluatedStatementReturnExpression(
	stmt *golox.StatementReturn,
	val any,
) {
	if itp.isDebug {
		log.Printf("%s: %s: in StatementReturn: evaluated %s = %v",
			log_prefix, stmt.GetLocation(), stmt.Expression, val,
		)
	}
}

func (itp *Interpreter) logExecutedStatementClass(
	stmt *golox.StatementClass,
	class *LoxClass,
) {
	if itp.isDebug {
		log.Printf("%s: %s: executed StatementClass: defined class %s with methods %v",
			log_prefix, stmt.GetLocation(), stmt.Identifier.Lexeme, class.Methods,
		)
	}
}

func (itp *Interpreter) logEvaluatedStatementPrintExpression(
	stmt *golox.StatementPrint,
	val any,
) {
	if itp.isDebug {
		log.Printf("%s: %s: in StatementPrint: evaluated %s = %v",
			log_prefix, stmt.GetLocation(), stmt.Expression, val,
		)
	}
}

func (itp *Interpreter) newErrorUndefinedVariable(
	identifier golox.Token,
) error {
	return fmt.Errorf("%s: undefined variable '%s'",
		identifier.Location, identifier.Lexeme,
	)
}

func (itp *Interpreter) newErrorOperandMustBe(
	message string, // e.g. "a number"
	opTkn golox.Token,
) error {
	return fmt.Errorf("%s: operand must be %s",
		opTkn.Location, message,
	)
}

func (itp *Interpreter) newErrorOperandsMustBe(
	message string, // e.g. "both numbers"
	opTkn golox.Token,
) error {
	return fmt.Errorf("%s: operands must be %s",
		opTkn.Location, message,
	)
}

func (itp *Interpreter) newErrorInvalidFunctionCallee(
	callee golox.Expression,
) error {
	return fmt.Errorf("%s: invalid function callee %s",
		callee.GetLocation(), callee,
	)
}

func (itp *Interpreter) newErrorFunctionArityMismatch(
	expr golox.Expression,
	want int,
	got int,
) error {
	return fmt.Errorf("%s: function call %s expected %d arguments, got %d",
		expr.GetLocation(), expr, want, got,
	)
}

func (itp *Interpreter) newErrorInvalidObjectInstance(
	expr golox.Expression,
) error {
	return fmt.Errorf("%s: invalid object instance %s",
		expr.GetLocation(), expr,
	)
}

func (itp *Interpreter) newErrorInvalidThisValue(
	expr golox.Expression,
	val any,
) error {
	return fmt.Errorf("%s: invalid value (%v) for 'this'",
		expr.GetLocation(), val,
	)
}

func (itp *Interpreter) newErrorInvalidClass(
	expr golox.Expression,
) error {
	return fmt.Errorf("%s: invalid class %s",
		expr.GetLocation(), expr,
	)
}

func (itp *Interpreter) newErrorInvalidSuperclassValue(
	expr golox.Expression,
	val any,
) error {
	return fmt.Errorf("%s: invalid superclass %v",
		expr.GetLocation(), val,
	)
}

func (c *LoxClass) newErrorUndefinedProperty(
	identifier golox.Token,
) error {
	return fmt.Errorf("%s: undefined property '%s'",
		identifier.Location, identifier.Lexeme,
	)
}

func (itp *Interpreter) newErrorMissingImplementation(
	node any,
) error {
	return fmt.Errorf("missing implementation for type %T",
		node,
	)
}
