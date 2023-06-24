package golox

import "fmt"

func NewErrorf(loc Location, format string, args ...any) error {
	return fmt.Errorf("%s: %s",
		loc, fmt.Sprintf(format, args...),
	)
}
