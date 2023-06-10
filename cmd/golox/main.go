package main

import (
	"fmt"
	"golox/internal/runner"
	"log"

	"github.com/spf13/pflag"
)

func handleErr(err error) {
	fmt.Println(err)
}

func main() {
	// parse flags:
	isDebug := pflag.Bool("debug", false, "enables debug logs")
	pflag.Parse()

	// parse args:
	args := pflag.Args()

	// init logger:
	log.SetFlags(0)

	// init runner:
	r := runner.NewRunner(*isDebug)

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
