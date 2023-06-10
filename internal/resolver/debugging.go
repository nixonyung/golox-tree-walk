package resolver

import (
	"fmt"
	golox "golox/internal"
	"log"
)

const (
	log_prefix = "DEBUG :: resolver"
)

func (r *Resolver) logDeclaredVariableInCurrScope(
	identifier golox.Token,
) {
	if r.isDebug {
		log.Printf("%s: %s: declared variable '%s' in scope level %d",
			log_prefix, identifier.Location, identifier.Lexeme, len(r.scopes),
		)
	}
}

func (r *Resolver) logDefinedVariableInCurrScope(
	identifier golox.Token,
) {
	if r.isDebug {
		log.Printf("%s: %s: defined variable '%s' in scope level %d",
			log_prefix, identifier.Location, identifier.Lexeme, len(r.scopes),
		)
	}
}

func (r *Resolver) logDefinedThisInCurrScope(
	classIdentifier golox.Token,
) {
	if r.isDebug {
		log.Printf("%s: %s: defined 'this' for class '%s' in scope level %d",
			log_prefix, classIdentifier.Location, classIdentifier.Lexeme, len(r.scopes),
		)
	}
}

func (r *Resolver) logDefinedSuperInCurrScope(
	classIdentifier golox.Token,
) {
	if r.isDebug {
		log.Printf("%s: %s: defined 'super' for class '%s' in scope level %d",
			log_prefix, classIdentifier.Location, classIdentifier.Lexeme, len(r.scopes),
		)
	}
}

func (r *Resolver) logResolvedVariable(
	identifier golox.Token,
	expr golox.Expression,
	dist int,
) {
	if r.isDebug {
		log.Printf("%s: %s: resolved '%s' with %s -> %d",
			log_prefix, identifier.Location, identifier.Lexeme, expr, dist,
		)
	}
}

func (r *Resolver) newErrorVariableIsAlreadyDefined(
	identifier golox.Token,
) error {
	return fmt.Errorf("%s: variable '%s' is already defined in current scope",
		identifier.Location, identifier.Lexeme,
	)
}

func (r *Resolver) newErrorVariableInItsOwnInitializer(
	identifier golox.Token,
) error {
	return fmt.Errorf("%s: cannot read variable '%s' in its own initializer",
		identifier.Location, identifier.Lexeme,
	)
}

func (r *Resolver) newErrorTopLevelReturn(
	returnToken golox.Token,
) error {
	return fmt.Errorf("%s: invalid top-level 'return'",
		returnToken.Location,
	)
}

func (r *Resolver) newErrorReturnWithValueInInitializer(
	returnToken golox.Token,
) error {
	return fmt.Errorf("%s: invalid return statement with value in an initializer",
		returnToken.Location,
	)
}

func (r *Resolver) newErrorTopLevelThis(
	thisToken golox.Token,
) error {
	return fmt.Errorf("%s: invalid top-level 'this'",
		thisToken.Location,
	)
}

func (r *Resolver) newErrorSuperOutsideClass(
	superToken golox.Token,
) error {
	return fmt.Errorf("%s: invalid 'super' outside a class",
		superToken.Location,
	)
}

func (r *Resolver) newErrorSuperWithoutSuperclass(
	superToken golox.Token,
) error {
	return fmt.Errorf("%s: invalid 'super' in a class with no superclass",
		superToken.Location,
	)
}

func (r *Resolver) newErrorClassInheritFromItself(
	classIdentifier golox.Token,
	superclassIdentifier golox.Token,
) error {
	return fmt.Errorf("%s: class '%s' inheriting from itself",
		superclassIdentifier.Location, classIdentifier.Lexeme,
	)
}

func (r *Resolver) newErrorMissingImplementation(
	node any,
) error {
	return fmt.Errorf("missing implementation for type %T",
		node,
	)
}
