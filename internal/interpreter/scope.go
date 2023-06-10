package interpreter

type Scope struct {
	NameToValue map[string]any
	Enclosing   *Scope
}
