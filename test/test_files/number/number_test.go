package number_test

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

func Test_decimal_point_at_eof(t *testing.T) {
	if err := r.RunFile("decimal_point_at_eof.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_leading_dot(t *testing.T) {
	if err := r.RunFile("leading_dot.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_literals() {
	if err := r.RunFile("literals.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// 123
	// 987654
	// 0
	// -0
	// 123.456
	// -0.001
}

func Example_nan_equality() {
	if err := r.RunFile("nan_equality.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// false
	// true
	// false
	// true
}

func Test_trailing_dot(t *testing.T) {
	if err := r.RunFile("trailing_dot.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}
