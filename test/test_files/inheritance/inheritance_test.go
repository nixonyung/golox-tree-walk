package inheritance_test

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

func Example_constructor() {
	if err := r.RunFile("constructor.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "value"
}

func Test_inherit_from_function(t *testing.T) {
	if err := r.RunFile("inherit_from_function.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_inherit_from_nil(t *testing.T) {
	if err := r.RunFile("inherit_from_nil.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_inherit_from_number(t *testing.T) {
	if err := r.RunFile("inherit_from_number.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_inherit_methods() {
	if err := r.RunFile("inherit_methods.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "foo"
	// "bar"
	// "bar"
}

func Test_parenthesized_superclass(t *testing.T) {
	if err := r.RunFile("parenthesized_superclass.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_set_fields_from_base_class() {
	if err := r.RunFile("set_fields_from_base_class.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "foo 1"
	// "foo 2"
	// "bar 1"
	// "bar 2"
	// "bar 1"
	// "bar 2"
}
