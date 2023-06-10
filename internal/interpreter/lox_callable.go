package interpreter

type LoxCallable interface {
	String() string
	Arity() int
	Call(args []any) (any, error)
}
