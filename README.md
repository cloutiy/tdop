# Top down operator precedence

This repository contains an implementation of a top down operator precedence parser written in Go. It accompanies my article at [cristiandima.com/top-down-operator-precedence-parsing-in-go/](http://www.cristiandima.com/top-down-operator-precedence-parsing-in-go/) where I write more on the method and use the code in this repo to show the algorithm in practice.

The program takes a source file name as an argument and prints out a lisp-like representation of the abstract syntax tree generated for that code. A source code sample file for the language this parser can parse is included in this repo (example.tdop).

## Examples

#### basic types

```go
x = 5;
isHuman = true;
name = "John Doe";
```

#### if blocks

```go
if true and not false {
  do_stuff();
}

if a == true and b == not true {
  do_this();
} else if not a and b {
  do_that();
} else {
  do_this_and_that();
}
```

#### while loops

```go
i = 0;
list = [1, 2, 3];
while i < len(list) {
    do_stuff(list[i]);
    i += 1;
}
```

#### functions

```go
sum = (a, b) -> a + b;
add_one = x -> x + 1;
print_hello = () -> {
    print("Hello World!!");
}
(s -> {
    text = "Hello " + s + "!!!";
    print(text);
})("John Doe");
```

#### tuples

```go
mult_and_sum = (a, b) -> {
    m = a * b;
    return (m, a+b);
}
```
