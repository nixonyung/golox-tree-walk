package comments_test

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

func Example_line_at_eof() {
	if err := r.RunFile("line_at_eof.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "ok"
}

func Example_only_line_comment() {
	if err := r.RunFile("only_line_comment.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:

}

func Example_only_line_comment_and_line() {
	if err := r.RunFile("only_line_comment_and_line.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:

}

func Example_unicode() {
	if err := r.RunFile("unicode.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "ok"
}
