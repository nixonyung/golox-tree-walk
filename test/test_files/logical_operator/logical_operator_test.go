package logical_operator_test

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

func Example_and() {
	if err := r.RunFile("and.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// false
	// 1
	// false
	// true
	// 3
	// true
	// false
}

func Example_and_truth() {
	if err := r.RunFile("and_truth.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// false
	// <nil>
	// "ok"
	// "ok"
	// "ok"
}

func Example_or() {
	if err := r.RunFile("or.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// 1
	// 1
	// true
	// false
	// false
	// false
	// true
}

func Example_or_truth() {
	if err := r.RunFile("or_truth.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "ok"
	// "ok"
	// true
	// 0
	// "s"
}
