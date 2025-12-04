package gounctional

import (
	"testing"
)

// testStruct implements the Comparer interface for testing purposes
type testStruct struct {
	value int
}

// Compare implements the Comparer interface
func (t testStruct) Compare(other testStruct) int {
	if t.value < other.value {
		return -1
	} else if t.value > other.value {
		return 1
	}
	return 0
}

// TestIsEqual tests the IsEqual function
func TestIsEqual(t *testing.T) {
	tests := []struct {
		name     string
		a, b     testStruct
		expected bool
	}{
		{
			name:     "equal values",
			a:        testStruct{value: 5},
			b:        testStruct{value: 5},
			expected: true,
		},
		{
			name:     "different values",
			a:        testStruct{value: 3},
			b:        testStruct{value: 7},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEqual(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("IsEqual(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestIsLess tests the IsLess function
func TestIsLess(t *testing.T) {
	tests := []struct {
		name     string
		a, b     testStruct
		expected bool
	}{
		{
			name:     "first less than second",
			a:        testStruct{value: 3},
			b:        testStruct{value: 7},
			expected: true,
		},
		{
			name:     "first equal to second",
			a:        testStruct{value: 5},
			b:        testStruct{value: 5},
			expected: false,
		},
		{
			name:     "first greater than second",
			a:        testStruct{value: 8},
			b:        testStruct{value: 6},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsLess(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("IsLess(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestIsGreater tests the IsGreater function
func TestIsGreater(t *testing.T) {
	tests := []struct {
		name     string
		a, b     testStruct
		expected bool
	}{
		{
			name:     "first greater than second",
			a:        testStruct{value: 8},
			b:        testStruct{value: 6},
			expected: true,
		},
		{
			name:     "first equal to second",
			a:        testStruct{value: 5},
			b:        testStruct{value: 5},
			expected: false,
		},
		{
			name:     "first less than second",
			a:        testStruct{value: 3},
			b:        testStruct{value: 7},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsGreater(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("IsGreater(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestIsLessOrEqual tests the IsLessOrEqual function
func TestIsLessOrEqual(t *testing.T) {
	tests := []struct {
		name     string
		a, b     testStruct
		expected bool
	}{
		{
			name:     "first less than second",
			a:        testStruct{value: 3},
			b:        testStruct{value: 7},
			expected: true,
		},
		{
			name:     "first equal to second",
			a:        testStruct{value: 5},
			b:        testStruct{value: 5},
			expected: true,
		},
		{
			name:     "first greater than second",
			a:        testStruct{value: 8},
			b:        testStruct{value: 6},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsLessOrEqual(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("IsLessOrEqual(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestIsGreaterOrEqual tests the IsGreaterOrEqual function
func TestIsGreaterOrEqual(t *testing.T) {
	tests := []struct {
		name     string
		a, b     testStruct
		expected bool
	}{
		{
			name:     "first greater than second",
			a:        testStruct{value: 8},
			b:        testStruct{value: 6},
			expected: true,
		},
		{
			name:     "first equal to second",
			a:        testStruct{value: 5},
			b:        testStruct{value: 5},
			expected: true,
		},
		{
			name:     "first less than second",
			a:        testStruct{value: 3},
			b:        testStruct{value: 7},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsGreaterOrEqual(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("IsGreaterOrEqual(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestIsDifferent tests the IsDifferent function
func TestIsDifferent(t *testing.T) {
	tests := []struct {
		name     string
		a, b     testStruct
		expected bool
	}{
		{
			name:     "different values",
			a:        testStruct{value: 3},
			b:        testStruct{value: 7},
			expected: true,
		},
		{
			name:     "equal values",
			a:        testStruct{value: 5},
			b:        testStruct{value: 5},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsDifferent(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("IsDifferent(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
