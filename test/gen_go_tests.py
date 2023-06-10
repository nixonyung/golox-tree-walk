import os
import re
from pathlib import Path


def header(module_name):
    return f"""package {module_name}_test

import (
	"fmt"
	"golox/internal/runner"
	"os"
	"testing"
)

const (
	ANSI_UNDERLINE = "\\x1b[4m"
	ANSI_FG_RED    = "\\x1b[31m"
	ANSI_FG_GREEN  = "\\x1b[32m"
	ANSI_RESET     = "\\x1b[0m"

	SUCCESS_TEXT = ANSI_UNDERLINE + "negative test " + ANSI_FG_GREEN + "SUCCESS" + ANSI_RESET
	FAILED_TEXT  = ANSI_UNDERLINE + "negative test " + ANSI_FG_RED + "FAILED" + ANSI_RESET
)

var (
	r *runner.Runner
)

func TestMain(m *testing.M) {{
	r = runner.NewRunner(false)

	// run tests
	os.Exit(m.Run())
}}

"""


def positive_test(file_name, expected_outputs):
    expected_outputs_text = "\n".join(
        f'    // "{output}"' for output in expected_outputs
    )

    return f"""func Example_{file_name}() {{
    if err := r.RunFile("{file_name}.lox"); err != nil {{
        fmt.Println(err)
    }}

    // Output:
{expected_outputs_text}
}}

"""


def negative_test(file_name):
    return f"""func Test_{file_name}(t *testing.T) {{
	if err := r.RunFile("{file_name}.lox"); err != nil {{
		t.Log(SUCCESS_TEXT+":", err)
	}} else {{
		t.Error(FAILED_TEXT)
	}}
}}

"""


src_dir = Path(__file__).parent / "test_files"

pattern_expect = re.compile(r"// expect: ?(.*)")
pattern_error = re.compile(r"// (Error.*)")
pattern_line_error = re.compile(r"// \[((java|c) )?line (\d+)\] (Error.*)")
pattern_runtime_error = re.compile(r"// expect runtime error: (.+)")
pattern_syntax_error = re.compile(r"\[.*line (\d+)\] (Error.+)")
# pattern_trace = re.compile(r"\[line (\d+)\]")
# pattern_nontest = re.compile(r"// nontest")

for root, _, files in os.walk(src_dir):
    module_name = os.path.split(root)[-1]

    if module_name in (
        "benchmark",
        "expressions",
        "limit",
        "scanning",
    ):
        continue

    with open(
        os.path.join(root, f"{module_name}_test.go"),
        "w",
        encoding="utf-8",
    ) as fp:
        fp.write(header(module_name))

        for file_name in sorted(files):
            file_name, ext = os.path.splitext(file_name)
            if ext != ".lox":
                continue

            text = Path(os.path.join(root, f"{file_name}.lox")).read_text(
                encoding="utf-8"
            )

            pattern_expect_matches = pattern_expect.findall(text)
            pattern_error_matches = pattern_error.findall(text)
            pattern_line_error_matches = pattern_line_error.findall(text)
            pattern_runtime_error_matches = pattern_runtime_error.findall(text)
            pattern_syntax_error_matches = pattern_syntax_error.findall(text)
            # pattern_trace_matches = pattern_trace.findall(text)
            # pattern_nontest_matches = pattern_nontest.findall(text)

            if (
                len(pattern_error_matches) != 0
                or len(pattern_line_error_matches) != 0
                or len(pattern_runtime_error_matches) != 0
                or len(pattern_syntax_error_matches) != 0
            ):
                fp.write(negative_test(file_name))
            else:
                fp.write(positive_test(file_name, pattern_expect_matches))
