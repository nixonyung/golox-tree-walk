package return_test

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

func Example_after_else() {
	if err := r.RunFile("after_else.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "ok"
}

func Example_after_if() {
	if err := r.RunFile("after_if.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "ok"
}

func Example_after_while() {
	if err := r.RunFile("after_while.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "ok"
}

func Test_at_top_level(t *testing.T) {
	if err := r.RunFile("at_top_level.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_in_function() {
	if err := r.RunFile("in_function.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "ok"
}

func Example_in_method() {
	if err := r.RunFile("in_method.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "ok"
}

func Example_return_nil_if_no_value() {
	if err := r.RunFile("return_nil_if_no_value.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// <nil>
}
