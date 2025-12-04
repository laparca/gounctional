# gounctional

[![Go Test](https://github.com/laparca/gounctional/actions/workflows/test.yml/badge.svg)](https://github.com/laparca/gounctional/actions/workflows/test.yml) [![Coverage](https://codecov.io/gh/laparca/gounctional/branch/main/graph/badge.svg)](https://codecov.io/gh/laparca/gounctional)

`gounctional` is a Go library that provides functional programming utilities and higher-order functions for common operations. It includes functions for comparison, arithmetic operations, function composition, and currying/binding operations, all built with Go's generic type system.

## Table of Contents

- [Features](#features)
- [Installation](#installation)  
- [Usage](#usage)
- [API Reference](#api-reference)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Generic Type Support**: All functions use Go generics for type safety
- **Comparison Operations**: Functions for comparing comparable types
- **Arithmetic Operations**: Generic arithmetic functions and partial application
- **Function Composition**: Utilities to compose functions together
- **Currying and Binding**: Generate curried functions with partial applications
- **Functional Utilities**: Identity functions, conditionals, and utility operations

## Installation

```bash
go get github.com/laparca/gounctional
```

## Usage

Import the package in your Go code:

```go
import "github.com/laparca/gounctional"
```

## API Reference

### Comparer Interface and Functions

`Comparer[T any]` interface defines a comparison operation for types that can be compared:

```go
type Comparer[T any] interface {
    Compare(other T) int
}
```

The interface should return:
- A negative number if the receiver is less than the other value
- Zero if the receiver is equal to the other value
- A positive number if the receiver is greater than the other value

#### Comparison Functions

- `IsEqual[T Comparer[T]](a, b T) bool`: Returns true if two values are equal
- `IsLess[T Comparer[T]](a, b T) bool`: Returns true if first value is less than second
- `IsGreater[T Comparer[T]](a, b T) bool`: Returns true if first value is greater than second
- `IsLessOrEqual[T Comparer[T]](a, b T) bool`: Returns true if first value is less than or equal to second
- `IsGreaterOrEqual[T Comparer[T]](a, b T) bool`: Returns true if first value is greater than or equal to second
- `IsDifferent[T Comparer[T]](a, b T) bool`: Returns true if two values are not equal

### Identity and Utility Functions

- `NotFunc() func(bool) bool`: Returns a function that negates a boolean value
- `IDFunc[T any]() func(T) T`: Returns the identity function that returns its input unchanged
- `EqualFunc[T comparable](value T) func(T) bool`: Returns a function that checks if a value is equal to the given value
- `DifferentFunc[T comparable](value T) func(T) bool`: Returns a function that checks if a value is different from the given value

### Relational Functions

- `LessThanFunc[T cmp.Ordered](value T) func(T) bool`: Returns a function that checks if a value is less than the given value
- `GreaterThanFunc[T cmp.Ordered](value T) func(T) bool`: Returns a function that checks if a value is greater than the given value
- `LessThanOrEqualFunc[T cmp.Ordered](value T) func(T) bool`: Checks if value is less than or equal to the given value
- `GreaterThanOrEqualFunc[T cmp.Ordered](value T) func(T) bool`: Checks if value is greater than or equal to the given value

### Arithmetic Functions

Type constraints:
- `Number`: Includes all numeric types (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64)
- `Integer`: Includes all integer types for modulo operations

#### Basic Operations
- `AddFunc[T Number]() func(T, T) T`: Returns a function that adds two values
- `SubFunc[T Number]() func(T, T) T`: Returns a function that subtracts second from first
- `MulFunc[T Number]() func(T, T) T`: Returns a function that multiplies two values
- `DivFunc[T Number]() func(T, T) T`: Returns a function that divides first by second
- `ModFunc[T Integer]() func(T, T) T`: Returns a function that calculates modulo

#### Partial Application Operations
- `AddNumberFunc[T Number](a T) func(T) T`: Partially applied add function
- `SubNumberFunc[T Number](a T) func(T) T`: Partially applied subtract function
- `MulNumberFunc[T Number](a T) func(T) T`: Partially applied multiply function
- `DivNumberFunc[T Number](a T) func(T) T`: Partially applied divide function
- `ModNumberFunc[T Integer](a T) func(T) T`: Partially applied modulo function

### Higher-Order Functions

- `Compose[T any, U any, R any](f func(T) U, g func(U) R) func(T) R`: Composes two functions (g(f(x)))
- `SwapFunc[T any, U any]() func(T, U) (U, T)`: Returns a function that swaps argument order
- `Iff[T any, R any](f func(T) bool, a func() R, b func() R) func(T) R`: Conditional function
- `Split[T any, R1 any, R2 any](f func(T) R1, g func(T) R2) func(T) (R1, R2)`: Applies two functions to the same input

### Binding Functions

The library provides binding functions from 1-of-1 to 10-of-10, allowing partial application of functions with multiple arguments. These are generated automatically using the build system.

## Examples

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/laparca/gounctional"
)

func main() {
    // Identity function
    id := gounctional.IDFunc[int]()
    fmt.Println(id(42)) // Output: 42

    // Not function
    not := gounctional.NotFunc()
    fmt.Println(not(true))  // Output: false
    fmt.Println(not(false)) // Output: true

    // Equal function
    isZero := gounctional.EqualFunc(0)
    fmt.Println(isZero(5)) // Output: false
    fmt.Println(isZero(0)) // Output: true

    // Arithmetic operations
    add := gounctional.AddFunc[int]()
    fmt.Println(add(3, 5)) // Output: 8

    // Partial application
    addFive := gounctional.AddNumberFunc(5)
    fmt.Println(addFive(10)) // Output: 15

    // Function composition
    double := gounctional.MulNumberFunc(2)
    addOne := gounctional.AddNumberFunc(1)
    doubleThenAddOne := gounctional.Compose(double, addOne)
    fmt.Println(doubleThenAddOne(5)) // Output: 11 (5*2=10, then 10+1=11)
}
```

### Comparison with Comparer Interface

```go
package main

import (
    "fmt"
    "github.com/laparca/gounctional"
)

type Person struct {
    Name string
    Age  int
}

// Implement the Comparer interface
func (p Person) Compare(other Person) int {
    if p.Age != other.Age {
        return p.Age - other.Age
    }
    if p.Name < other.Name {
        return -1
    }
    if p.Name > other.Name {
        return 1
    }
    return 0
}

func main() {
    p1 := Person{Name: "Alice", Age: 30}
    p2 := Person{Name: "Bob", Age: 25}

    fmt.Println(gounctional.IsGreater(p1, p2))  // true
    fmt.Println(gounctional.IsLess(p1, p2))     // false
    fmt.Println(gounctional.IsEqual(p1, p2))    // false
}
```

### Using Relational Functions

```go
package main

import (
    "fmt"
    "github.com/laparca/gounctional"
)

func main() {
    // Create functions that compare against a fixed value
    lessThanTen := gounctional.LessThanFunc(10)
    greaterThanFive := gounctional.GreaterThanFunc(5)

    fmt.Println(lessThanTen(7))   // true
    fmt.Println(lessThanTen(15))  // false
    fmt.Println(greaterThanFive(7))  // true
    fmt.Println(greaterThanFive(3))  // false
}
```

### Function Composition and Higher-Order Functions

```go
package main

import (
    "fmt"
    "github.com/laparca/gounctional"
)

func main() {
    // Function composition: apply f first, then g
    double := gounctional.MulNumberFunc(2)
    addOne := gounctional.AddNumberFunc(1)
    doubleThenAddOne := gounctional.Compose(double, addOne)
    fmt.Println(doubleThenAddOne(5)) // 11

    // Split function: apply two functions to the same input
    square := func(x int) int { return x * x }
    doubleFunc := func(x int) int { return x * 2 }
    split := gounctional.Split(square, doubleFunc)
    sq, dbl := split(5)
    fmt.Println(sq, dbl) // 25 10

    // Conditional function
    isPositive := func(x int) bool { return x > 0 }
    getPositive := func() string { return "positive" }
    getNegative := func() string { return "negative or zero" }
    iff := gounctional.Iff(isPositive, getPositive, getNegative)
    fmt.Println(iff(5))  // "positive"
    fmt.Println(iff(-3)) // "negative or zero"
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the terms found in the [LICENSE](LICENSE) file.