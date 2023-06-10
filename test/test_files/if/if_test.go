package if_test

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

func Test_class_in_else(t *testing.T) {
	if err := r.RunFile("class_in_else.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_class_in_then(t *testing.T) {
	if err := r.RunFile("class_in_then.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_dangling_else() {
	if err := r.RunFile("dangling_else.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "good"
}

func Example_else() {
	if err := r.RunFile("else.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "good"
	// "good"
	// "block"
}

func Test_fun_in_else(t *testing.T) {
	if err := r.RunFile("fun_in_else.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_fun_in_then(t *testing.T) {
	if err := r.RunFile("fun_in_then.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_if() {
	if err := r.RunFile("if.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "good"
	// "block"
	// true
}

func Example_truth() {
	if err := r.RunFile("truth.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "false"
	// "nil"
	// true
	// 0
	// "empty"
}

func Test_var_in_else(t *testing.T) {
	if err := r.RunFile("var_in_else.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_var_in_then(t *testing.T) {
	if err := r.RunFile("var_in_then.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}
