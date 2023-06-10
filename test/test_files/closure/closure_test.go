package closure_test

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

func Example_assign_to_closure() {
	if err := r.RunFile("assign_to_closure.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "local"
	// "after f"
	// "after f"
	// "after g"
}

func Example_assign_to_shadowed_later() {
	if err := r.RunFile("assign_to_shadowed_later.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "inner"
	// "assigned"
}

func Example_close_over_function_parameter() {
	if err := r.RunFile("close_over_function_parameter.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "param"
}

func Example_close_over_later_variable() {
	if err := r.RunFile("close_over_later_variable.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "b"
	// "a"
}

func Example_close_over_method_parameter() {
	if err := r.RunFile("close_over_method_parameter.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "param"
}

func Example_closed_closure_in_function() {
	if err := r.RunFile("closed_closure_in_function.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "local"
}

func Example_nested_closure() {
	if err := r.RunFile("nested_closure.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "a"
	// "b"
	// "c"
}

func Example_open_closure_in_function() {
	if err := r.RunFile("open_closure_in_function.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "local"
}

func Example_reference_closure_multiple_times() {
	if err := r.RunFile("reference_closure_multiple_times.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "a"
	// "a"
}

func Example_reuse_closure_slot() {
	if err := r.RunFile("reuse_closure_slot.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "a"
}

func Example_shadow_closure_with_local() {
	if err := r.RunFile("shadow_closure_with_local.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "closure"
	// "shadow"
	// "closure"
}

func Example_unused_closure() {
	if err := r.RunFile("unused_closure.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "ok"
}

func Example_unused_later_closure() {
	if err := r.RunFile("unused_later_closure.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "a"
}
