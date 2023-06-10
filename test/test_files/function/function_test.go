package function_test

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

func Test_body_must_be_block(t *testing.T) {
	if err := r.RunFile("body_must_be_block.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_empty_body() {
	if err := r.RunFile("empty_body.lox"); err != nil {
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

func Test_local_mutual_recursion(t *testing.T) {
	if err := r.RunFile("local_mutual_recursion.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_local_recursion() {
	if err := r.RunFile("local_recursion.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// 21
}

func Test_missing_arguments(t *testing.T) {
	if err := r.RunFile("missing_arguments.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_missing_comma_in_parameters(t *testing.T) {
	if err := r.RunFile("missing_comma_in_parameters.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_mutual_recursion() {
	if err := r.RunFile("mutual_recursion.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// true
	// true
}

func Example_nested_call_with_arguments() {
	if err := r.RunFile("nested_call_with_arguments.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "hello world"
}

func Example_parameters() {
	if err := r.RunFile("parameters.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// 0
	// 1
	// 3
	// 6
	// 10
	// 15
	// 21
	// 28
	// 36
}

func Example_print() {
	if err := r.RunFile("print.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// <fn: foo>
	// <native fn: clock>
}

func Example_recursion() {
	if err := r.RunFile("recursion.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// 21
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
