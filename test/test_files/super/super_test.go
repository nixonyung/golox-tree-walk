package super_test

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

func Example_bound_method() {
	if err := r.RunFile("bound_method.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "A.method(arg)"
}

func Example_call_other_method() {
	if err := r.RunFile("call_other_method.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "Derived.bar()"
	// "Base.foo()"
}

func Example_call_same_method() {
	if err := r.RunFile("call_same_method.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "Derived.foo()"
	// "Base.foo()"
}

func Example_closure() {
	if err := r.RunFile("closure.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "Base"
}

func Example_constructor() {
	if err := r.RunFile("constructor.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "Derived.init()"
	// "Base.init(a, b)"
}

func Test_extra_arguments(t *testing.T) {
	if err := r.RunFile("extra_arguments.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_indirectly_inherited() {
	if err := r.RunFile("indirectly_inherited.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "C.foo()"
	// "A.foo()"
}

func Test_missing_arguments(t *testing.T) {
	if err := r.RunFile("missing_arguments.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_no_superclass_bind(t *testing.T) {
	if err := r.RunFile("no_superclass_bind.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_no_superclass_call(t *testing.T) {
	if err := r.RunFile("no_superclass_call.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_no_superclass_method(t *testing.T) {
	if err := r.RunFile("no_superclass_method.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_parenthesized(t *testing.T) {
	if err := r.RunFile("parenthesized.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_reassign_superclass() {
	if err := r.RunFile("reassign_superclass.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "Base.method()"
	// "Base.method()"
}

func Test_super_at_top_level(t *testing.T) {
	if err := r.RunFile("super_at_top_level.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_super_in_closure_in_inherited_method() {
	if err := r.RunFile("super_in_closure_in_inherited_method.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "A"
}

func Example_super_in_inherited_method() {
	if err := r.RunFile("super_in_inherited_method.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "A"
}

func Test_super_in_top_level_function(t *testing.T) {
	if err := r.RunFile("super_in_top_level_function.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_super_without_dot(t *testing.T) {
	if err := r.RunFile("super_without_dot.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_super_without_name(t *testing.T) {
	if err := r.RunFile("super_without_name.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_this_in_superclass_method() {
	if err := r.RunFile("this_in_superclass_method.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "a"
	// "b"
}
