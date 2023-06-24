package interpreter

import "strings"

type Scope struct {
	NameToValue map[string]any
	Enclosing   *Scope
}

// debug use
func (s *Scope) level() int {
	if s.Enclosing == nil {
		return 0
	} else {
		return 1 + s.Enclosing.level()
	}
}

// debug use
func (s *Scope) string() string {
	var builder strings.Builder
	builder.WriteByte('[')
	i := 0
	for identifier := range s.NameToValue {
		if i != 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(identifier)
		i++
	}
	builder.WriteByte(']')
	return builder.String()
}
