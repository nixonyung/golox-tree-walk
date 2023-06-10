package variable_test

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

func Test_collide_with_parameter(t *testing.T) {
	if err := r.RunFile("collide_with_parameter.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_duplicate_local(t *testing.T) {
	if err := r.RunFile("duplicate_local.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_duplicate_parameter(t *testing.T) {
	if err := r.RunFile("duplicate_parameter.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_early_bound() {
	if err := r.RunFile("early_bound.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "outer"
	// "outer"
}

func Example_in_middle_of_block() {
	if err := r.RunFile("in_middle_of_block.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "a"
	// "a b"
	// "a c"
	// "a b d"
}

func Example_in_nested_block() {
	if err := r.RunFile("in_nested_block.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "outer"
}

func Example_local_from_method() {
	if err := r.RunFile("local_from_method.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "variable"
}

func Example_redeclare_global() {
	if err := r.RunFile("redeclare_global.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// <nil>
}

func Example_redefine_global() {
	if err := r.RunFile("redefine_global.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "2"
}

func Example_scope_reuse_in_different_blocks() {
	if err := r.RunFile("scope_reuse_in_different_blocks.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "first"
	// "second"
}

func Example_shadow_and_local() {
	if err := r.RunFile("shadow_and_local.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "outer"
	// "inner"
}

func Example_shadow_global() {
	if err := r.RunFile("shadow_global.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "shadow"
	// "global"
}

func Example_shadow_local() {
	if err := r.RunFile("shadow_local.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "shadow"
	// "local"
}

func Test_undefined_global(t *testing.T) {
	if err := r.RunFile("undefined_global.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_undefined_local(t *testing.T) {
	if err := r.RunFile("undefined_local.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_uninitialized() {
	if err := r.RunFile("uninitialized.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// <nil>
}

func Example_unreached_undefined() {
	if err := r.RunFile("unreached_undefined.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "ok"
}

func Test_use_false_as_var(t *testing.T) {
	if err := r.RunFile("use_false_as_var.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_use_global_in_initializer() {
	if err := r.RunFile("use_global_in_initializer.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "value"
}

func Test_use_local_in_initializer(t *testing.T) {
	if err := r.RunFile("use_local_in_initializer.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_use_nil_as_var(t *testing.T) {
	if err := r.RunFile("use_nil_as_var.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_use_this_as_var(t *testing.T) {
	if err := r.RunFile("use_this_as_var.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}
