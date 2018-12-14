# Monkey

This is a language called _Monkey_ I am building based on the book [Writing
an Interpreter in Go][1] by Thorsten Ball.

Here are the features of the Monkey language, taken directly from the book:

- C-like syntax
- Variable bindings
- Integers and booleans
- Arithmetic expressions
- Built-in functions
- First-class and higher-order functions
- Closures
- A string data structure
- An array data structure
- A hash data structure

### Setup

Monkey is written in Go, [install it][2] for your OS.

Then, clone the repository into your GOPATH and run the tests:

```
$ git clone http://github.com:jwworth/monkey.git
$ cd monkey
$ go test ./...
```

### REPL

Monkey includes the Monkey REPL. Load it as follows:

```
$ go run main.go
```

### Code of Conduct

This project is intended to be a safe, welcoming space for collaboration, and
contributors are expected to adhere to the [Contributor
Covenant](http://contributor-covenant.org) code of conduct. Please see [CODE OF
CONDUCT](CODE_OF_CONDUCT.md) for more information.

### License

None of this content was created by me. This project is released under the same
license terms as the book.

[1]: https://interpreterbook.com/
[2]: https://golang.org/doc/install/
