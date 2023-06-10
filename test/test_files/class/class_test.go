package class_test

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

func Example_empty() {
	if err := r.RunFile("empty.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// <class: Foo>
}

func Test_inherit_self(t *testing.T) {
	if err := r.RunFile("inherit_self.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_inherited_method() {
	if err := r.RunFile("inherited_method.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "in foo"
	// "in bar"
	// "in baz"
}

func Example_local_inherit_other() {
	if err := r.RunFile("local_inherit_other.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// <class: B>
}

func Test_local_inherit_self(t *testing.T) {
	if err := r.RunFile("local_inherit_self.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_local_reference_self() {
	if err := r.RunFile("local_reference_self.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// <class: Foo>
}

func Example_reference_self() {
	if err := r.RunFile("reference_self.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// <class: Foo>
}
