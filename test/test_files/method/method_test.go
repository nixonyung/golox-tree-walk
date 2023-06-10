package method_test

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

func Example_arity() {
	if err := r.RunFile("arity.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "no args"
	// 1
	// 3
	// 6
	// 10
	// 15
	// 21
	// 28
	// 36
}

func Example_empty_block() {
	if err := r.RunFile("empty_block.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// <nil>
}

func Test_extra_arguments(t *testing.T) {
	if err := r.RunFile("extra_arguments.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_missing_arguments(t *testing.T) {
	if err := r.RunFile("missing_arguments.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_not_found(t *testing.T) {
	if err := r.RunFile("not_found.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_print_bound_method() {
	if err := r.RunFile("print_bound_method.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// <fn: method>
}

func Test_refer_to_name(t *testing.T) {
	if err := r.RunFile("refer_to_name.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_too_many_arguments(t *testing.T) {
	if err := r.RunFile("too_many_arguments.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_too_many_parameters(t *testing.T) {
	if err := r.RunFile("too_many_parameters.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}
