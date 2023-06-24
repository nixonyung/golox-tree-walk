package golox

import (
	"fmt"
	"log"
)

//go:generate stringer -type=Module
type Module int

const (
	ModuleLexer       Module = iota
	ModuleParser      Module = iota
	ModuleResolver    Module = iota
	ModuleInterpreter Module = iota
)

func Logf(module Module, format string, args ...any) {
	if ConfigIsDebug {
		if format == "" {
			log.Printf("DEBUG[%s]",
				module,
			)
		} else {
			log.Printf("DEBUG[%s] %s",
				module, fmt.Sprintf(format, args...),
			)
		}
	}
}
