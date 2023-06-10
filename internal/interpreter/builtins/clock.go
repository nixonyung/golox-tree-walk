package builtins

import "time"

// builtin functions:

type Clock struct{}

func (c *Clock) String() string {
	return "<native fn: clock>"
}

func (c *Clock) Arity() int {
	return 0
}

func (c *Clock) Call(args []any) (any, error) {
	return float64(time.Now().UnixMilli() / 1000), nil
}
