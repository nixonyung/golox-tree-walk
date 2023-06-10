package operator_test

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

func Example_add() {
	if err := r.RunFile("add.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// 579
	// "string"
}

func Test_add_bool_nil(t *testing.T) {
	if err := r.RunFile("add_bool_nil.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_add_bool_num(t *testing.T) {
	if err := r.RunFile("add_bool_num.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_add_bool_string(t *testing.T) {
	if err := r.RunFile("add_bool_string.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_add_nil_nil(t *testing.T) {
	if err := r.RunFile("add_nil_nil.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_add_num_nil(t *testing.T) {
	if err := r.RunFile("add_num_nil.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_add_string_nil(t *testing.T) {
	if err := r.RunFile("add_string_nil.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_comparison() {
	if err := r.RunFile("comparison.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// true
	// false
	// false
	// true
	// true
	// false
	// false
	// false
	// true
	// false
	// true
	// true
	// false
	// false
	// false
	// false
	// true
	// true
	// true
	// true
}

func Example_divide() {
	if err := r.RunFile("divide.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// 4
	// 1
}

func Test_divide_nonnum_num(t *testing.T) {
	if err := r.RunFile("divide_nonnum_num.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_divide_num_nonnum(t *testing.T) {
	if err := r.RunFile("divide_num_nonnum.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_equals() {
	if err := r.RunFile("equals.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// true
	// true
	// false
	// true
	// false
	// true
	// false
	// false
	// false
	// false
}

func Example_equals_class() {
	if err := r.RunFile("equals_class.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// true
	// false
	// false
	// true
	// false
	// false
	// false
	// false
}

func Example_equals_method() {
	if err := r.RunFile("equals_method.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// true
	// false
}

func Test_greater_nonnum_num(t *testing.T) {
	if err := r.RunFile("greater_nonnum_num.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_greater_num_nonnum(t *testing.T) {
	if err := r.RunFile("greater_num_nonnum.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_greater_or_equal_nonnum_num(t *testing.T) {
	if err := r.RunFile("greater_or_equal_nonnum_num.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_greater_or_equal_num_nonnum(t *testing.T) {
	if err := r.RunFile("greater_or_equal_num_nonnum.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_less_nonnum_num(t *testing.T) {
	if err := r.RunFile("less_nonnum_num.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_less_num_nonnum(t *testing.T) {
	if err := r.RunFile("less_num_nonnum.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_less_or_equal_nonnum_num(t *testing.T) {
	if err := r.RunFile("less_or_equal_nonnum_num.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_less_or_equal_num_nonnum(t *testing.T) {
	if err := r.RunFile("less_or_equal_num_nonnum.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_multiply() {
	if err := r.RunFile("multiply.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// 15
	// 3.702
}

func Test_multiply_nonnum_num(t *testing.T) {
	if err := r.RunFile("multiply_nonnum_num.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_multiply_num_nonnum(t *testing.T) {
	if err := r.RunFile("multiply_num_nonnum.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_negate() {
	if err := r.RunFile("negate.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// -3
	// 3
	// -3
}

func Test_negate_nonnum(t *testing.T) {
	if err := r.RunFile("negate_nonnum.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_not() {
	if err := r.RunFile("not.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// false
	// true
	// true
	// false
	// false
	// true
	// false
	// false
}

func Example_not_class() {
	if err := r.RunFile("not_class.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// false
	// false
}

func Example_not_equals() {
	if err := r.RunFile("not_equals.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// false
	// false
	// true
	// false
	// true
	// false
	// true
	// true
	// true
	// true
}

func Example_subtract() {
	if err := r.RunFile("subtract.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// 1
	// 0
}

func Test_subtract_nonnum_num(t *testing.T) {
	if err := r.RunFile("subtract_nonnum_num.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_subtract_num_nonnum(t *testing.T) {
	if err := r.RunFile("subtract_num_nonnum.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}
