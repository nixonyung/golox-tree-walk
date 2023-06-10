# golox

A Go implementation of Lox, as introduced in the book [Crafting Interpreters](http://craftinginterpreters.com/)

---

- This implementation followed [Chapter II of the book - A TREE-WALK INTERPRETER](https://craftinginterpreters.com/a-tree-walk-interpreter.html).
  - No challenges are done. No additional language features are added.
  - A debug mode is added, use the `--debug` flag
  - Tests are adopted from [the official repository](https://github.com/munificent/craftinginterpreters/tree/master/test).

## Getting started

### Prerequisites

- [Go](https://go.dev/dl/)
- run `go mod tidy -v`

### Using the interpreter

- start an interactive session:
  - run `go run cmd/golox/main.go`
  - run `go run cmd/golox/main.go --debug` (or run `./scripts/debug_repl.sh`)
- run a Lox script:
  - run `go run cmd/golox/main.go <script>`
  - run `go run cmd/golox/main.go <script> --debug` (or run `./scripts/debug_file.sh <script>`)

### Testing

- run `go test ./test/...` (or run `./scripts/test.sh`)

---

&copy; 2023 Nixon Yung, All Rights Reserved.
