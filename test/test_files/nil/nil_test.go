package nil_test

import (
	"fmt"
	"golox/internal/runner"
	"os"
	"testing"
)

const (
	ANSI_UNDERLINE = "\x1b[4m"
	ANSI_FG_RED    = "\x1b[31m"
	ANSI_FG_GREEN  = "\x1b[32m"
	ANSI_RESET     = "\x1b[0m"

	SUCCESS_TEXT = ANSI_UNDERLINE + "negative test " + ANSI_FG_GREEN + "SUCCESS" + ANSI_RESET
	FAILED_TEXT  = ANSI_UNDERLINE + "negative test " + ANSI_FG_RED + "FAILED" + ANSI_RESET
)

var (
	r *runner.Runner
)

func TestMain(m *testing.M) {
	r = runner.NewRunner(false)

	// run tests
	os.Exit(m.Run())
}

func Example_literal() {
	if err := r.RunFile("literal.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// <nil>
}
