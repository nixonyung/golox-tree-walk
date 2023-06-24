package main

import (
	"fmt"
	lox "golox/internal"
	"golox/internal/runner"
	"log"

	"github.com/spf13/pflag"
)

func handleErr(err error) {
	fmt.Println(err)
}

func main() {
	// parse flags:
	pflag.BoolVar(&lox.ConfigIsDebug, "debug", false, "enables debug logs")
	pflag.Parse()

	// parse args:
	args := pflag.Args()

	// init logger:
	log.SetFlags(0)

	// init runner:
	r := runner.NewRunner(lox.ConfigIsDebug)

	switch {
	case len(args) > 1:
		panic("usage: lox [script_path]")
	case len(args) == 1:
		if err := r.RunFile(args[0]); err != nil {
			handleErr(err)
		}
	default:
		r.RunPrompt(handleErr)
	}
}
