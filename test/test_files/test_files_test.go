package test_files_test

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

func Example_empty_file() {
	if err := r.RunFile("empty_file.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:

}

func Example_precedence() {
	if err := r.RunFile("precedence.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// 14
	// 8
	// 4
	// 0
	// true
	// true
	// true
	// true
	// 0
	// 0
	// 0
	// 0
	// 4
}

func Test_unexpected_character(t *testing.T) {
	if err := r.RunFile("unexpected_character.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}
