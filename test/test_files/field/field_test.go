package field_test

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

func Example_call_function_field() {
	if err := r.RunFile("call_function_field.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "bar"
	// 1
	// 2
}

func Test_call_nonfunction_field(t *testing.T) {
	if err := r.RunFile("call_nonfunction_field.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_get_and_set_method() {
	if err := r.RunFile("get_and_set_method.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "other"
	// 1
	// "method"
	// 2
}

func Test_get_on_bool(t *testing.T) {
	if err := r.RunFile("get_on_bool.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_get_on_class(t *testing.T) {
	if err := r.RunFile("get_on_class.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_get_on_function(t *testing.T) {
	if err := r.RunFile("get_on_function.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_get_on_nil(t *testing.T) {
	if err := r.RunFile("get_on_nil.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_get_on_num(t *testing.T) {
	if err := r.RunFile("get_on_num.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_get_on_string(t *testing.T) {
	if err := r.RunFile("get_on_string.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Example_many() {
	if err := r.RunFile("many.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "apple"
	// "apricot"
	// "avocado"
	// "banana"
	// "bilberry"
	// "blackberry"
	// "blackcurrant"
	// "blueberry"
	// "boysenberry"
	// "cantaloupe"
	// "cherimoya"
	// "cherry"
	// "clementine"
	// "cloudberry"
	// "coconut"
	// "cranberry"
	// "currant"
	// "damson"
	// "date"
	// "dragonfruit"
	// "durian"
	// "elderberry"
	// "feijoa"
	// "fig"
	// "gooseberry"
	// "grape"
	// "grapefruit"
	// "guava"
	// "honeydew"
	// "huckleberry"
	// "jabuticaba"
	// "jackfruit"
	// "jambul"
	// "jujube"
	// "juniper"
	// "kiwifruit"
	// "kumquat"
	// "lemon"
	// "lime"
	// "longan"
	// "loquat"
	// "lychee"
	// "mandarine"
	// "mango"
	// "marionberry"
	// "melon"
	// "miracle"
	// "mulberry"
	// "nance"
	// "nectarine"
	// "olive"
	// "orange"
	// "papaya"
	// "passionfruit"
	// "peach"
	// "pear"
	// "persimmon"
	// "physalis"
	// "pineapple"
	// "plantain"
	// "plum"
	// "plumcot"
	// "pomegranate"
	// "pomelo"
	// "quince"
	// "raisin"
	// "rambutan"
	// "raspberry"
	// "redcurrant"
	// "salak"
	// "salmonberry"
	// "satsuma"
	// "strawberry"
	// "tamarillo"
	// "tamarind"
	// "tangerine"
	// "tomato"
	// "watermelon"
	// "yuzu"
}

func Example_method() {
	if err := r.RunFile("method.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "got method"
	// "arg"
}

func Example_method_binds_this() {
	if err := r.RunFile("method_binds_this.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "foo1"
	// 1
}

func Example_on_instance() {
	if err := r.RunFile("on_instance.lox"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// "bar value"
	// "baz value"
	// "bar value"
	// "baz value"
}

func Test_set_evaluation_order(t *testing.T) {
	if err := r.RunFile("set_evaluation_order.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_set_on_bool(t *testing.T) {
	if err := r.RunFile("set_on_bool.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_set_on_class(t *testing.T) {
	if err := r.RunFile("set_on_class.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_set_on_function(t *testing.T) {
	if err := r.RunFile("set_on_function.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_set_on_nil(t *testing.T) {
	if err := r.RunFile("set_on_nil.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_set_on_num(t *testing.T) {
	if err := r.RunFile("set_on_num.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_set_on_string(t *testing.T) {
	if err := r.RunFile("set_on_string.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}

func Test_undefined(t *testing.T) {
	if err := r.RunFile("undefined.lox"); err != nil {
		t.Log(SUCCESS_TEXT+":", err)
	} else {
		t.Error(FAILED_TEXT)
	}
}
