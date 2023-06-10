package lox

import "fmt"

type Location struct {
	SrcPath string
	Line    int
	Col     int
}

func (l Location) String() string {
	return fmt.Sprintf("%s:%d:%d",
		l.SrcPath, l.Line, l.Col,
	)
}
