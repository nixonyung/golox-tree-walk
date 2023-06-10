package bool_test

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

func Example_equality() {
	if err := r.RunFile("equality.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// true
	// false
	// false
	// true
	// false
	// false
	// false
	// false
	// false
	// false
	// true
	// true
	// false
	// true
	// true
	// true
	// true
	// true
}

func Example_not() {
	if err := r.RunFile("not.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// false
	// true
	// true
}
