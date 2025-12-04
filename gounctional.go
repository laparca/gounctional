// Package gounctional provides functional programming utilities and higher-order functions
// for common operations in Go. It includes functions for comparison, arithmetic operations,
// function composition, and currying/binding operations.
package gounctional

import "cmp"

// NotFunc returns a function that negates a boolean value.
//
// Example:
//
//	negate := gounctional.NotFunc()
//	result := negate(true)  // false
//	result = negate(false)  // true
func NotFunc() func(bool) bool {
	return func(value bool) bool {
		return !value
	}
}

// IDFunc returns the identity function that returns its input unchanged.
//
// Example:
//
//	id := gounctional.IDFunc[int]()
//	result := id(42)  // 42
//	idStr := gounctional.IDFunc[string]()
//	resultStr := idStr("hello")  // "hello"
func IDFunc[T any]() func(T) T {
	return func(value T) T {
		return value
	}
}

// EqualFunc returns a function that checks if a value is equal to the given value.
//
// Example:
//
//	isZero := gounctional.EqualFunc(0)
//	result := isZero(5)   // false
//	result = isZero(0)    // true
//
//	isHello := gounctional.EqualFunc("hello")
//	result := isHello("world")  // false
//	result = isHello("hello")   // true
func EqualFunc[T comparable](value T) func(T) bool {
	return func(other T) bool {
		return value == other
	}
}

// DifferentFunc returns a function that checks if a value is different from the given value.
//
// Example:
//
//	notZero := gounctional.DifferentFunc(0)
//	result := notZero(5)   // true
//	result = notZero(0)    // false
//
//	notHello := gounctional.DifferentFunc("hello")
//	result := notHello("world")  // true
//	result = notHello("hello")   // false
func DifferentFunc[T comparable](value T) func(T) bool {
	return func(other T) bool {
		return value != other
	}
}

// LessThanFunc returns a function that checks if a value is less than the given value.
//
// Example:
//
//	lessThanFive := gounctional.LessThanFunc(5)
//	result := lessThanFive(3)   // true
//	result = lessThanFive(7)    // false
//
//	lessThanA := gounctional.LessThanFunc("b")
//	result = lessThanA("a")     // true
//	result = lessThanA("z")     // false
func LessThanFunc[T cmp.Ordered](value T) func(T) bool {
	return func(other T) bool {
		return cmp.Compare(other, value) < 0
	}
}

// GreaterThanFunc returns a function that checks if a value is greater than the given value.
//
// Example:
//
//	greaterThanFive := gounctional.GreaterThanFunc(5)
//	result := greaterThanFive(7)   // true
//	result = greaterThanFive(3)    // false
//
//	greaterThanB := gounctional.GreaterThanFunc("b")
//	result = greaterThanB("c")     // true
//	result = greaterThanB("a")     // false
func GreaterThanFunc[T cmp.Ordered](value T) func(T) bool {
	return func(other T) bool {
		return cmp.Compare(other, value) > 0
	}
}

// LessThanOrEqualFunc returns a function that checks if a value is less than or equal to the given value.
//
// Example:
//
//	lessOrEqualFive := gounctional.LessThanOrEqualFunc(5)
//	result := lessOrEqualFive(3)   // true
//	result = lessOrEqualFive(5)    // true
//	result = lessOrEqualFive(7)    // false
func LessThanOrEqualFunc[T cmp.Ordered](value T) func(T) bool {
	return func(other T) bool {
		return cmp.Compare(other, value) <= 0
	}
}

// GreaterThanOrEqualFunc returns a function that checks if a value is greater than or equal to the given value.
//
// Example:
//
//	greaterOrEqualFive := gounctional.GreaterThanOrEqualFunc(5)
//	result := greaterOrEqualFive(7)   // true
//	result = greaterOrEqualFive(5)    // true
//	result = greaterOrEqualFive(3)    // false
func GreaterThanOrEqualFunc[T cmp.Ordered](value T) func(T) bool {
	return func(other T) bool {
		return cmp.Compare(other, value) >= 0
	}
}

// SwapFunc returns a function that swaps the order of two arguments.
//
// Example:
//
//	swap := gounctional.SwapFunc[int, string]()
//	first, second := swap(42, "hello")  // "hello", 42
func SwapFunc[T any, U any]() func(T, U) (U, T) {
	return func(a T, b U) (U, T) {
		return b, a
	}
}

// Number is a type constraint that includes all numeric types that support arithmetic operations.
// This includes signed integers, unsigned integers, and floating point numbers.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Integer is a type constraint that includes all integer types, both signed and unsigned.
// This constraint is used by operations that require integer types only (like Modulo operations).
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// AddFunc returns a function that adds two values of the same numeric type.
//
// Example:
//
//	add := gounctional.AddFunc[int]()
//	result := add(3, 5)  // 8
//
//	addFloat := gounctional.AddFunc[float64]()
//	result = addFloat(2.5, 1.5)  // 4.0
func AddFunc[T Number]() func(T, T) T {
	return func(a T, b T) T {
		return a + b
	}
}

// SubFunc returns a function that subtracts the second value from the first value.
//
// Example:
//
//	sub := gounctional.SubFunc[int]()
//	result := sub(10, 3)  // 7
//
//	subFloat := gounctional.SubFunc[float64]()
//	result = subFloat(5.5, 2.0)  // 3.5
func SubFunc[T Number]() func(T, T) T {
	return func(a T, b T) T {
		return a - b
	}
}

// MulFunc returns a function that multiplies two values of the same numeric type.
//
// Example:
//
//	mul := gounctional.MulFunc[int]()
//	result := mul(4, 5)  // 20
//
//	mulFloat := gounctional.MulFunc[float64]()
//	result = mulFloat(2.5, 4.0)  // 10.0
func MulFunc[T Number]() func(T, T) T {
	return func(a T, b T) T {
		return a * b
	}
}

// DivFunc returns a function that divides the first value by the second value.
//
// Example:
//
//	div := gounctional.DivFunc[int]()
//	result := div(10, 2)  // 5
//
//	divFloat := gounctional.DivFunc[float64]()
//	result = divFloat(7.5, 2.5)  // 3.0
func DivFunc[T Number]() func(T, T) T {
	return func(a T, b T) T {
		return a / b
	}
}

// ModFunc returns a function that calculates the modulo (remainder) of dividing the first value by the second value.
// This function works with integer types only.
//
// Example:
//
//	mod := gounctional.ModFunc[int]()
//	result := mod(10, 3)  // 1
//
//	mod64 := gounctional.ModFunc[int64]()
//	result = mod64(15, 4)  // 3
func ModFunc[T Integer]() func(T, T) T {
	return func(a T, b T) T {
		return a % b
	}
}

// AddNumberFunc returns a partially applied function that adds a fixed value to its input.
// The fixed value is provided as the parameter to this function.
//
// Example:
//
//	add5 := gounctional.AddNumberFunc(5)
//	result := add5(3)  // 8
//
//	addFloat := gounctional.AddNumberFunc(2.5)
//	result = addFloat(1.5)  // 4.0
func AddNumberFunc[T Number](a T) func(T) T {
	return func(b T) T {
		return b + a
	}
}

// SubNumberFunc returns a partially applied function that subtracts a fixed value from its input.
// The fixed value is provided as the parameter to this function.
//
// Example:
//
//	sub5 := gounctional.SubNumberFunc(5)
//	result := sub5(10)  // 5
//
//	subFloat := gounctional.SubNumberFunc(2.5)
//	result = subFloat(5.0)  // 2.5
func SubNumberFunc[T Number](a T) func(T) T {
	return func(b T) T {
		return b - a
	}
}

// MulNumberFunc returns a partially applied function that multiplies its input by a fixed value.
// The fixed value is provided as the parameter to this function.
//
// Example:
//
//	mul5 := gounctional.MulNumberFunc(5)
//	result := mul5(4)  // 20
//
//	mulFloat := gounctional.MulNumberFunc(2.5)
//	result = mulFloat(4.0)  // 10.0
func MulNumberFunc[T Number](a T) func(T) T {
	return func(b T) T {
		return b * a
	}
}

// DivNumberFunc returns a partially applied function that divides its input by a fixed value.
// The fixed value is provided as the parameter to this function.
//
// Example:
//
//	div2 := gounctional.DivNumberFunc(2)
//	result := div2(10)  // 5
//
//	divFloat := gounctional.DivNumberFunc(2.5)
//	result = divFloat(7.5)  // 3.0
func DivNumberFunc[T Number](a T) func(T) T {
	return func(b T) T {
		return b / a
	}
}

// ModNumberFunc returns a partially applied function that calculates the modulo (remainder) of dividing its input by a fixed value.
// The fixed value is provided as the parameter to this function. This works with integer types only.
//
// Example:
//
//	mod3 := gounctional.ModNumberFunc(3)
//	result := mod3(10)  // 1
//
//	mod64 := gounctional.ModNumberFunc(int64(4))
//	result = mod64(15)  // 3
func ModNumberFunc[T Integer](a T) func(T) T {
	return func(b T) T {
		return b % a
	}
}

// Compose returns a function that represents the composition of two functions f and g.
// The resulting function applies f first, then applies g to the result of f.
// This is equivalent to g(f(x)).
//
// Example:
//
//	double := gounctional.MulNumberFunc(2)
//	addOne := gounctional.AddNumberFunc(1)
//	doubleThenAddOne := gounctional.Compose(double, addOne)
//	result := doubleThenAddOne(5)  // 11 (5*2=10, then 10+1=11)
func Compose[T any, U any, R any](f func(T) U, g func(U) R) func(T) R {
	return func(t T) R {
		return g(f(t))
	}
}

// Iff returns a conditional function that executes one of two functions based on a predicate.
// If the predicate function returns true, it executes the 'a' function and returns its result.
// Otherwise, it executes the 'b' function and returns its result.
//
// Example:
//
//	isPositive := func(x int) bool { return x > 0 }
//	getPositiveMsg := func() string { return "positive" }
//	getNegativeMsg := func() string { return "negative or zero" }
//	conditionalMsg := gounctional.Iff(isPositive, getPositiveMsg, getNegativeMsg)
//
//	result1 := conditionalMsg(5)   // "positive"
//	result2 := conditionalMsg(-3)  // "negative or zero"
func Iff[T any, R any](f func(T) bool, a func() R, b func() R) func(T) R {
	return func(t T) R {
		if f(t) {
			return a()
		}
		return b()
	}
}

// Split returns a function that applies two given functions to the same input and returns both results as a tuple.
// This is useful for applying multiple transformations to the same value simultaneously.
//
// Example:
//
//	square := func(x int) int { return x * x }
//	double := func(x int) int { return x * 2 }
//	splitFunc := gounctional.Split(square, double)
//
//	sq, dbl := splitFunc(5)  // sq = 25, dbl = 10
func Split[T any, R1 any, R2 any](f func(T) R1, g func(T) R2) func(T) (R1, R2) {
	return func(t T) (R1, R2) {
		return f(t), g(t)
	}
}

//go:generate go run cmd/generate_bind.go -s 1 -e 10 -o gounctional_bind.go
