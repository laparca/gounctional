package gounctional

import (
	"testing"
)

// TestNotFunc tests the NotFunc function
func TestNotFunc(t *testing.T) {
	fn := NotFunc()
	tests := []struct {
		input    bool
		expected bool
	}{
		{true, false},
		{false, true},
	}

	for _, tt := range tests {
		result := fn(tt.input)
		if result != tt.expected {
			t.Errorf("NotFunc()(%v) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

// TestIDFunc tests the IDFunc function
func TestIDFunc(t *testing.T) {
	intFn := IDFunc[int]()
	strFn := IDFunc[string]()

	tests := []struct {
		name     string
		intInput int
		intWant  int
		strInput string
		strWant  string
	}{
		{
			name:     "integer identity",
			intInput: 42,
			intWant:  42,
		},
		{
			name:     "string identity",
			strInput: "hello",
			strWant:  "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "integer identity" {
				result := intFn(tt.intInput)
				if result != tt.intWant {
					t.Errorf("IDFunc[int()](%v) = %v, want %v", tt.intInput, result, tt.intWant)
				}
			} else if tt.name == "string identity" {
				result := strFn(tt.strInput)
				if result != tt.strWant {
					t.Errorf("IDFunc[string()](%v) = %v, want %v", tt.strInput, result, tt.strWant)
				}
			}
		})
	}
}

// TestEqualFunc tests the EqualFunc function
func TestEqualFunc(t *testing.T) {
	isZero := EqualFunc(0)
	isHello := EqualFunc("hello")

	tests := []struct {
		name     string
		fn       func(interface{}) bool
		input    interface{}
		expected bool
	}{
		{
			name:     "integer equal",
			fn:       func(i interface{}) bool { return isZero(i.(int)) },
			input:    0,
			expected: true,
		},
		{
			name:     "integer not equal",
			fn:       func(i interface{}) bool { return isZero(i.(int)) },
			input:    5,
			expected: false,
		},
		{
			name:     "string equal",
			fn:       func(s interface{}) bool { return isHello(s.(string)) },
			input:    "hello",
			expected: true,
		},
		{
			name:     "string not equal",
			fn:       func(s interface{}) bool { return isHello(s.(string)) },
			input:    "world",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn(tt.input)
			if result != tt.expected {
				t.Errorf("%T(%v) = %v, want %v", tt.fn, tt.input, result, tt.expected)
			}
		})
	}
}

// TestDifferentFunc tests the DifferentFunc function
func TestDifferentFunc(t *testing.T) {
	notZero := DifferentFunc(0)
	notHello := DifferentFunc("hello")

	tests := []struct {
		name     string
		fn       func(interface{}) bool
		input    interface{}
		expected bool
	}{
		{
			name:     "integer different",
			fn:       func(i interface{}) bool { return notZero(i.(int)) },
			input:    5,
			expected: true,
		},
		{
			name:     "integer not different",
			fn:       func(i interface{}) bool { return notZero(i.(int)) },
			input:    0,
			expected: false,
		},
		{
			name:     "string different",
			fn:       func(s interface{}) bool { return notHello(s.(string)) },
			input:    "world",
			expected: true,
		},
		{
			name:     "string not different",
			fn:       func(s interface{}) bool { return notHello(s.(string)) },
			input:    "hello",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn(tt.input)
			if result != tt.expected {
				t.Errorf("%T(%v) = %v, want %v", tt.fn, tt.input, result, tt.expected)
			}
		})
	}
}

// TestLessThanFunc tests the LessThanFunc function
func TestLessThanFunc(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "less than",
			input:    3,
			expected: true,
		},
		{
			name:     "equal to",
			input:    5,
			expected: false,
		},
		{
			name:     "greater than",
			input:    7,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := LessThanFunc(5) // Use consistent comparison value
			result := fn(tt.input)
			if result != tt.expected {
				t.Errorf("LessThanFunc(5)(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}

	// Test with strings
	lessThanB := LessThanFunc("b")
	if !lessThanB("a") {
		t.Errorf("LessThanFunc(\"b\")(\"a\") = false, want true")
	}
	if lessThanB("b") {
		t.Errorf("LessThanFunc(\"b\")(\"b\") = true, want false")
	}
	if lessThanB("c") {
		t.Errorf("LessThanFunc(\"b\")(\"c\") = true, want false")
	}
}

// TestGreaterThanFunc tests the GreaterThanFunc function
func TestGreaterThanFunc(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "greater than",
			input:    7,
			expected: true,
		},
		{
			name:     "equal to",
			input:    5,
			expected: false,
		},
		{
			name:     "less than",
			input:    3,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := GreaterThanFunc(5)
			result := fn(tt.input)
			if result != tt.expected {
				t.Errorf("GreaterThanFunc(5)(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}

	// Test with strings
	greaterThanB := GreaterThanFunc("b")
	if !greaterThanB("c") {
		t.Errorf("GreaterThanFunc(\"b\")(\"c\") = false, want true")
	}
	if greaterThanB("b") {
		t.Errorf("GreaterThanFunc(\"b\")(\"b\") = true, want false")
	}
	if greaterThanB("a") {
		t.Errorf("GreaterThanFunc(\"b\")(\"a\") = true, want false")
	}
}

// TestLessThanOrEqualFunc tests the LessThanOrEqualFunc function
func TestLessThanOrEqualFunc(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "less than",
			input:    3,
			expected: true,
		},
		{
			name:     "equal to",
			input:    5,
			expected: true,
		},
		{
			name:     "greater than",
			input:    7,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := LessThanOrEqualFunc(5)
			result := fn(tt.input)
			if result != tt.expected {
				t.Errorf("LessThanOrEqualFunc(5)(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestGreaterThanOrEqualFunc tests the GreaterThanOrEqualFunc function
func TestGreaterThanOrEqualFunc(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "greater than",
			input:    7,
			expected: true,
		},
		{
			name:     "equal to",
			input:    5,
			expected: true,
		},
		{
			name:     "less than",
			input:    3,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := GreaterThanOrEqualFunc(5)
			result := fn(tt.input)
			if result != tt.expected {
				t.Errorf("GreaterThanOrEqualFunc(5)(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestSwapFunc tests the SwapFunc function
func TestSwapFunc(t *testing.T) {
	swap := SwapFunc[int, string]()
	first, second := swap(42, "hello")
	if first != "hello" || second != 42 {
		t.Errorf("SwapFunc[int, string()](42, \"hello\") = (%v, %v), want (\"hello\", 42)", first, second)
	}

	// Test with different types
	swapStrInt := SwapFunc[string, int]()
	first2, second2 := swapStrInt("world", 100)
	if first2 != 100 || second2 != "world" {
		t.Errorf("SwapFunc[string, int()](\"world\", 100) = (%v, %v), want (100, \"world\")", first2, second2)
	}
}

// TestAddFunc tests the AddFunc function
func TestAddFunc(t *testing.T) {
	add := AddFunc[int]()
	result := add(3, 5)
	if result != 8 {
		t.Errorf("AddFunc[int()]()(3, 5) = %v, want 8", result)
	}

	// Test with float64
	addFloat := AddFunc[float64]()
	resultFloat := addFloat(2.5, 1.5)
	if resultFloat != 4.0 {
		t.Errorf("AddFunc[float64()](2.5, 1.5) = %v, want 4.0", resultFloat)
	}
}

// TestSubFunc tests the SubFunc function
func TestSubFunc(t *testing.T) {
	sub := SubFunc[int]()
	result := sub(10, 3)
	if result != 7 {
		t.Errorf("SubFunc[int()]()(10, 3) = %v, want 7", result)
	}

	// Test with float64
	subFloat := SubFunc[float64]()
	resultFloat := subFloat(5.5, 2.0)
	if resultFloat != 3.5 {
		t.Errorf("SubFunc[float64()](5.5, 2.0) = %v, want 3.5", resultFloat)
	}
}

// TestMulFunc tests the MulFunc function
func TestMulFunc(t *testing.T) {
	mul := MulFunc[int]()
	result := mul(4, 5)
	if result != 20 {
		t.Errorf("MulFunc[int()]()(4, 5) = %v, want 20", result)
	}

	// Test with float64
	mulFloat := MulFunc[float64]()
	resultFloat := mulFloat(2.5, 4.0)
	if resultFloat != 10.0 {
		t.Errorf("MulFunc[float64()](2.5, 4.0) = %v, want 10.0", resultFloat)
	}
}

// TestDivFunc tests the DivFunc function
func TestDivFunc(t *testing.T) {
	div := DivFunc[int]()
	result := div(10, 2)
	if result != 5 {
		t.Errorf("DivFunc[int()]()(10, 2) = %v, want 5", result)
	}

	// Test with float64
	divFloat := DivFunc[float64]()
	resultFloat := divFloat(7.5, 2.5)
	if resultFloat != 3.0 {
		t.Errorf("DivFunc[float64()](7.5, 2.5) = %v, want 3.0", resultFloat)
	}
}

// TestModFunc tests the ModFunc function
func TestModFunc(t *testing.T) {
	mod := ModFunc[int]()
	result := mod(10, 3)
	if result != 1 {
		t.Errorf("ModFunc[int()]()(10, 3) = %v, want 1", result)
	}

	// Test with int64
	mod64 := ModFunc[int64]()
	result64 := mod64(15, 4)
	if result64 != 3 {
		t.Errorf("ModFunc[int64()](15, 4) = %v, want 3", result64)
	}
}

// TestAddNumberFunc tests the AddNumberFunc function
func TestAddNumberFunc(t *testing.T) {
	add5 := AddNumberFunc(5)
	result := add5(3)
	if result != 8 {
		t.Errorf("AddNumberFunc(5)(3) = %v, want 8", result)
	}

	// Test with float64
	addFloat := AddNumberFunc(2.5)
	resultFloat := addFloat(1.5)
	if resultFloat != 4.0 {
		t.Errorf("AddNumberFunc(2.5)(1.5) = %v, want 4.0", resultFloat)
	}
}

// TestSubNumberFunc tests the SubNumberFunc function
func TestSubNumberFunc(t *testing.T) {
	sub5 := SubNumberFunc(5)
	result := sub5(10)
	if result != 5 {
		t.Errorf("SubNumberFunc(5)(10) = %v, want 5", result)
	}

	// Test with float64
	subFloat := SubNumberFunc(2.5)
	resultFloat := subFloat(5.0)
	if resultFloat != 2.5 {
		t.Errorf("SubNumberFunc(2.5)(5.0) = %v, want 2.5", resultFloat)
	}
}

// TestMulNumberFunc tests the MulNumberFunc function
func TestMulNumberFunc(t *testing.T) {
	mul5 := MulNumberFunc(5)
	result := mul5(4)
	if result != 20 {
		t.Errorf("MulNumberFunc(5)(4) = %v, want 20", result)
	}

	// Test with float64
	mulFloat := MulNumberFunc(2.5)
	resultFloat := mulFloat(4.0)
	if resultFloat != 10.0 {
		t.Errorf("MulNumberFunc(2.5)(4.0) = %v, want 10.0", resultFloat)
	}
}

// TestDivNumberFunc tests the DivNumberFunc function
func TestDivNumberFunc(t *testing.T) {
	div2 := DivNumberFunc(2)
	result := div2(10)
	if result != 5 {
		t.Errorf("DivNumberFunc(2)(10) = %v, want 5", result)
	}

	// Test with float64
	divFloat := DivNumberFunc(2.5)
	resultFloat := divFloat(7.5)
	if resultFloat != 3.0 {
		t.Errorf("DivNumberFunc(2.5)(7.5) = %v, want 3.0", resultFloat)
	}
}

// TestModNumberFunc tests the ModNumberFunc function
func TestModNumberFunc(t *testing.T) {
	mod3 := ModNumberFunc(3)
	result := mod3(10)
	if result != 1 {
		t.Errorf("ModNumberFunc(3)(10) = %v, want 1", result)
	}

	// Test with int64
	mod64 := ModNumberFunc(int64(4))
	result64 := mod64(15)
	if result64 != 3 {
		t.Errorf("ModNumberFunc(int64(4))(15) = %v, want 3", result64)
	}
}

// TestCompose tests the Compose function
func TestCompose(t *testing.T) {
	double := MulNumberFunc(2)
	addOne := AddNumberFunc(1)
	doubleThenAddOne := Compose(double, addOne)

	result := doubleThenAddOne(5)
	if result != 11 {
		t.Errorf("Compose(double, addOne)(5) = %v, want 11", result) // (5*2=10, then 10+1=11)
	}

	// Test compose with different function types
	toString := func(i int) string {
		return "value: " + string(rune('0'+i))
	}
	toLength := func(s string) int {
		return len(s)
	}
	composed := Compose(toString, toLength)
	result2 := composed(5)
	// "value: " + "5" = "value: 5", which has length 8 (v-a-l-u-e-:-space-5)
	expectedLen := len("value: 5") // This should be 8
	if result2 != expectedLen {
		t.Errorf("Compose(toString, toLength)(5) = %v, want %v", result2, expectedLen)
	}
}

// TestIff tests the Iff function
func TestIff(t *testing.T) {
	isPositive := func(x int) bool { return x > 0 }
	getPositive := func() int { return 1 }
	getNegative := func() int { return -1 }

	iffFunc := Iff(isPositive, getPositive, getNegative)

	result1 := iffFunc(5)
	if result1 != 1 {
		t.Errorf("Iff with positive input: got %v, want 1", result1)
	}

	result2 := iffFunc(-3)
	if result2 != -1 {
		t.Errorf("Iff with negative input: got %v, want -1", result2)
	}

	result3 := iffFunc(0)
	if result3 != -1 {
		t.Errorf("Iff with zero input: got %v, want -1", result3)
	}
}

// TestSplit tests the Split function
func TestSplit(t *testing.T) {
	square := func(x int) int { return x * x }
	double := func(x int) int { return x * 2 }

	splitFunc := Split(square, double)
	first, second := splitFunc(5)

	if first != 25 {
		t.Errorf("Split first result: got %v, want 25", first)
	}
	if second != 10 {
		t.Errorf("Split second result: got %v, want 10", second)
	}
}
