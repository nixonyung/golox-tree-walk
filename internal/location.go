package golox

import "fmt"

// token location in source file
type Location struct {
	SrcPath string
	Line    int
	Col     int
}

func (loc Location) String() string {
	return fmt.Sprintf("%s:%d:%d",
		loc.SrcPath, loc.Line, loc.Col,
	)
}
