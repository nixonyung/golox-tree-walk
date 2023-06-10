package constructor_test

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

func Example_arguments() {
	if err := r.RunFile("arguments.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "init"
	// 1
	// 2
}

func Example_call_init_early_return() {
	if err := r.RunFile("call_init_early_return.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "init"
	// "init"
	// <instance of <class: Foo>>
}

func Example_call_init_explicitly() {
	if err := r.RunFile("call_init_explicitly.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "Foo.init(one)"
	// "Foo.init(two)"
	// <instance of <class: Foo>>
	// "init"
}

func Example_default() {
	if err := r.RunFile("default.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// <instance of <class: Foo>>
}

func Test_default_arguments(t *testing.T) {
	if err := r.RunFile("default_arguments.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_early_return() {
	if err := r.RunFile("early_return.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "init"
	// <instance of <class: Foo>>
}

func Test_extra_arguments(t *testing.T) {
	if err := r.RunFile("extra_arguments.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_init_not_method() {
	if err := r.RunFile("init_not_method.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "not initializer"
}

func Test_missing_arguments(t *testing.T) {
	if err := r.RunFile("missing_arguments.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_return_in_nested_function() {
	if err := r.RunFile("return_in_nested_function.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "bar"
	// <instance of <class: Foo>>
}

func Test_return_value(t *testing.T) {
	if err := r.RunFile("return_value.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}
