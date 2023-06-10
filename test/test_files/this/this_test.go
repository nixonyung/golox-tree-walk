package this_test

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

func Example_closure() {
	if err := r.RunFile("closure.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "Foo"
}

func Example_nested_class() {
	if err := r.RunFile("nested_class.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// <instance of <class: Outer>>
	// <instance of <class: Outer>>
	// <instance of <class: Inner>>
}

func Example_nested_closure() {
	if err := r.RunFile("nested_closure.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "Foo"
}

func Test_this_at_top_level(t *testing.T) {
	if err := r.RunFile("this_at_top_level.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_this_in_method() {
	if err := r.RunFile("this_in_method.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "baz"
}

func Test_this_in_top_level_function(t *testing.T) {
	if err := r.RunFile("this_in_top_level_function.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}
