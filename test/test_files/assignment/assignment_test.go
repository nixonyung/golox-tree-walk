package assignment_test

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

func Example_associativity() {
	if err := r.RunFile("associativity.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "c"
	// "c"
	// "c"
}

func Example_global() {
	if err := r.RunFile("global.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "before"
	// "after"
	// "arg"
	// "arg"
}

func Test_grouping(t *testing.T) {
	if err := r.RunFile("grouping.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_infix_operator(t *testing.T) {
	if err := r.RunFile("infix_operator.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_local() {
	if err := r.RunFile("local.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "before"
	// "after"
	// "arg"
	// "arg"
}

func Test_prefix_operator(t *testing.T) {
	if err := r.RunFile("prefix_operator.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_syntax() {
	if err := r.RunFile("syntax.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "var"
	// "var"
}

func Test_to_this(t *testing.T) {
	if err := r.RunFile("to_this.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_undefined(t *testing.T) {
	if err := r.RunFile("undefined.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}
