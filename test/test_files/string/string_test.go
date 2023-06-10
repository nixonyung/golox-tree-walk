package string_test

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

func Test_error_after_multiline(t *testing.T) {
	if err := r.RunFile("error_after_multiline.lox"); err != nil {
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
	// "()"
	// "a string"
	// "A~¶Þॐஃ"
}

func Example_multiline() {
	if err := r.RunFile("multiline.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "1
	// 2
	// 3"
}

func Test_unterminated(t *testing.T) {
	if err := r.RunFile("unterminated.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}
