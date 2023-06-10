package runner

import (
	"bufio"
	"bytes"
	"fmt"
	"golox/internal/interpreter"
	"golox/internal/lexer"
	"golox/internal/parser"
	"golox/internal/resolver"
	"os"
	"path/filepath"
)

type Runner struct {
	// configs:
	isDebug bool

	// states:
	srcPath     string
	interpreter *interpreter.Interpreter
}

func (r *Runner) run(source []rune) error {
	tokens, err := lexer.
		NewLexer(r.isDebug, r.srcPath).
		TokensFromSource(source)
	if err != nil {
		return err
	}

	stmts, err := parser.
		NewParser(r.isDebug).
		StatementsFromTokens(tokens)
	if err != nil {
		return err
	}

	resolvedLocalVars, err := resolver.
		NewResolver(r.isDebug).
		ResolveStatements(stmts)
	if err != nil {
		return err
	}

	// reuse interpreter to persist scopes in a run session
	if err := r.interpreter.
		InterpretStatements(stmts, resolvedLocalVars); err != nil {
		return err
	}

	return nil
}

func (r *Runner) RunFile(path string) error {
	r.srcPath = path
	if pwd, err := os.Getwd(); err == nil {
		r.srcPath = filepath.Join(pwd, path)
	}
	r.interpreter = interpreter.NewInterpreter(r.isDebug)

	source, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := r.run(bytes.Runes(source)); err != nil {
		return err
	}

	return nil
}

func (r *Runner) RunPrompt(errHandler func(error)) {
	r.srcPath = "REPL"
	r.interpreter = interpreter.NewInterpreter(r.isDebug)

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("An interactive session of golox. Press Ctrl-d to end.")
		fmt.Print("> ")
		if ok := reader.Scan(); !ok {
			// e.g. detected ctrl+d
			break
		} else if err := r.run(bytes.Runes(reader.Bytes())); err != nil {
			errHandler(err)
		}
	}
}

func NewRunner(
	isDebug bool,
) *Runner {
	return &Runner{
		isDebug:     isDebug,
		interpreter: nil,
	}
}
