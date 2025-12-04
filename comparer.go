package gounctional

// Comparer is an interface for types that can be compared with each other.
//
// The Compare method should return:
//   - a negative number if the receiver is less than the other value
//   - zero if the receiver is equal to the other value
//   - a positive number if the receiver is greater than the other value
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//
//	func (p Person) Compare(other Person) int {
//		if p.Age != other.Age {
//			return p.Age - other.Age
//		}
//		if p.Name < other.Name {
//			return -1
//		}
//		if p.Name > other.Name {
//			return 1
//		}
//		return 0
//	}
type Comparer[T any] interface {
	Compare(other T) int
}

// IsEqual compares two values of type T using their Compare method and returns true if they are equal.
//
// Two values are considered equal if the Compare method returns 0.
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//
//	func (p Person) Compare(other Person) int {
//		return p.Age - other.Age
//	}
//
//	p1 := Person{Name: "Alice", Age: 30}
//	p2 := Person{Name: "Bob", Age: 30}
//	result := IsEqual(p1, p2) // returns true
func IsEqual[T Comparer[T]](a, b T) bool {
	return a.Compare(b) == 0
}

// IsLess compares two values of type T using their Compare method and returns true if the first is less than the second.
//
// A value is considered less than another if the Compare method returns a negative number.
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//
//	func (p Person) Compare(other Person) int {
//		return p.Age - other.Age
//	}
//
//	p1 := Person{Name: "Alice", Age: 25}
//	p2 := Person{Name: "Bob", Age: 30}
//	result := IsLess(p1, p2) // returns true
func IsLess[T Comparer[T]](a, b T) bool {
	return a.Compare(b) < 0
}

// IsGreater compares two values of type T using their Compare method and returns true if the first is greater than the second.
//
// A value is considered greater than another if the Compare method returns a positive number.
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//
//	func (p Person) Compare(other Person) int {
//		return p.Age - other.Age
//	}
//
//	p1 := Person{Name: "Alice", Age: 35}
//	p2 := Person{Name: "Bob", Age: 30}
//	result := IsGreater(p1, p2) // returns true
func IsGreater[T Comparer[T]](a, b T) bool {
	return a.Compare(b) > 0
}

// IsLessOrEqual compares two values of type T using their Compare method and returns true if the first is less than or equal to the second.
//
// A value is considered less than or equal to another if the Compare method returns a negative number or zero.
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//
//	func (p Person) Compare(other Person) int {
//		return p.Age - other.Age
//	}
//
//	p1 := Person{Name: "Alice", Age: 25}
//	p2 := Person{Name: "Bob", Age: 30}
//	result := IsLessOrEqual(p1, p2) // returns true
//
//	p3 := Person{Name: "Charlie", Age: 25}
//	p4 := Person{Name: "David", Age: 25}
//	result2 := IsLessOrEqual(p3, p4) // returns true
func IsLessOrEqual[T Comparer[T]](a, b T) bool {
	return a.Compare(b) <= 0
}

// IsGreaterOrEqual compares two values of type T using their Compare method and returns true if the first is greater than or equal to the second.
//
// A value is considered greater than or equal to another if the Compare method returns a positive number or zero.
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//
//	func (p Person) Compare(other Person) int {
//		return p.Age - other.Age
//	}
//
//	p1 := Person{Name: "Alice", Age: 35}
//	p2 := Person{Name: "Bob", Age: 30}
//	result := IsGreaterOrEqual(p1, p2) // returns true
//
//	p3 := Person{Name: "Charlie", Age: 30}
//	p4 := Person{Name: "David", Age: 30}
//	result2 := IsGreaterOrEqual(p3, p4) // returns true
func IsGreaterOrEqual[T Comparer[T]](a, b T) bool {
	return a.Compare(b) >= 0
}

// IsDifferent compares two values of type T using their Compare method and returns true if they are not equal.
//
// Two values are considered different if the Compare method returns a non-zero value.
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//
//	func (p Person) Compare(other Person) int {
//		return p.Age - other.Age
//	}
//
//	p1 := Person{Name: "Alice", Age: 30}
//	p2 := Person{Name: "Bob", Age: 35}
//	result := IsDifferent(p1, p2) // returns true
//
//	p3 := Person{Name: "Charlie", Age: 30}
//	p4 := Person{Name: "David", Age: 30}
//	result2 := IsDifferent(p3, p4) // returns false
func IsDifferent[T Comparer[T]](a, b T) bool {
	return a.Compare(b) != 0
}
