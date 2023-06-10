package while_test

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

func Test_class_in_body(t *testing.T) {
	if err := r.RunFile("class_in_body.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_closure_in_body() {
	if err := r.RunFile("closure_in_body.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// 1
	// 2
	// 3
}

func Test_fun_in_body(t *testing.T) {
	if err := r.RunFile("fun_in_body.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_return_closure() {
	if err := r.RunFile("return_closure.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "i"
}

func Example_return_inside() {
	if err := r.RunFile("return_inside.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "i"
}

func Example_syntax() {
	if err := r.RunFile("syntax.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// 1
	// 2
	// 3
	// 0
	// 1
	// 2
}

func Test_var_in_body(t *testing.T) {
	if err := r.RunFile("var_in_body.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}
