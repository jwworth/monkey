# Monkey

This is a language called _Monkey_ I am building based on the book [Writing
an Interpreter in Go][1] by Thorsten Ball.

Here are the features of the Monkey language, taken directly from the book:

- C-like syntax
- variable bindings
- integers and booleans
- arithmetic expressions
- built-in functions
- first-class and higher-order functions
- closures
- a string data structure
- an array data structure
- a hash data structure

### Setup

Monkey is written in Go, [install it][2] for your OS.

Then, clone the repository into your GOPATH and run the tests:

```
$ git clone http://github.com:jwworth/monkey.git
$ cd monkey
$ go test ./...
```

[1]: https://interpreterbook.com/
[2]: https://golang.org/doc/install/

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
