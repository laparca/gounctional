package gounctional

import (
	"testing"
)

// TestBind1Of1 tests the Bind1Of1 function
func TestBind1Of1(t *testing.T) {
	fn := func(x int) int {
		return x * 2
	}
	bound := Bind1Of1(fn, 5)
	result := bound()
	expected := 10
	if result != expected {
		t.Errorf("Bind1Of1 result: got %v, want %v", result, expected)
	}
}

// TestBind1Of2 tests the Bind1Of2 function
func TestBind1Of2(t *testing.T) {
	fn := func(x, y int) int {
		return x + y
	}
	bound := Bind1Of2(fn, 5)
	result := bound(3)
	expected := 8
	if result != expected {
		t.Errorf("Bind1Of2 result: got %v, want %v", result, expected)
	}
}

// TestBind2Of2 tests the Bind2Of2 function
func TestBind2Of2(t *testing.T) {
	fn := func(x, y int) int {
		return x * y
	}
	bound := Bind2Of2(fn, 4, 3)
	result := bound()
	expected := 12
	if result != expected {
		t.Errorf("Bind2Of2 result: got %v, want %v", result, expected)
	}
}

// TestBind1Of3 tests the Bind1Of3 function
func TestBind1Of3(t *testing.T) {
	fn := func(x, y, z int) int {
		return x + y + z
	}
	bound := Bind1Of3(fn, 2)
	result := bound(3, 4)
	expected := 9
	if result != expected {
		t.Errorf("Bind1Of3 result: got %v, want %v", result, expected)
	}
}

// TestBind2Of3 tests the Bind2Of3 function
func TestBind2Of3(t *testing.T) {
	fn := func(x, y, z int) int {
		return x + y + z
	}
	bound := Bind2Of3(fn, 2, 3)
	result := bound(4)
	expected := 9
	if result != expected {
		t.Errorf("Bind2Of3 result: got %v, want %v", result, expected)
	}
}

// TestBind3Of3 tests the Bind3Of3 function
func TestBind3Of3(t *testing.T) {
	fn := func(x, y, z int) int {
		return x * y * z
	}
	bound := Bind3Of3(fn, 2, 3, 4)
	result := bound()
	expected := 24
	if result != expected {
		t.Errorf("Bind3Of3 result: got %v, want %v", result, expected)
	}
}

// TestBind1Of4 tests the Bind1Of4 function
func TestBind1Of4(t *testing.T) {
	fn := func(w, x, y, z int) int {
		return w + x + y + z
	}
	bound := Bind1Of4(fn, 1)
	result := bound(2, 3, 4)
	expected := 10
	if result != expected {
		t.Errorf("Bind1Of4 result: got %v, want %v", result, expected)
	}
}

// TestBind2Of4 tests the Bind2Of4 function
func TestBind2Of4(t *testing.T) {
	fn := func(w, x, y, z int) int {
		return w + x + y + z
	}
	bound := Bind2Of4(fn, 1, 2)
	result := bound(3, 4)
	expected := 10
	if result != expected {
		t.Errorf("Bind2Of4 result: got %v, want %v", result, expected)
	}
}

// TestBind3Of4 tests the Bind3Of4 function
func TestBind3Of4(t *testing.T) {
	fn := func(w, x, y, z int) int {
		return w + x + y + z
	}
	bound := Bind3Of4(fn, 1, 2, 3)
	result := bound(4)
	expected := 10
	if result != expected {
		t.Errorf("Bind3Of4 result: got %v, want %v", result, expected)
	}
}

// TestBind4Of4 tests the Bind4Of4 function
func TestBind4Of4(t *testing.T) {
	fn := func(w, x, y, z int) int {
		return w * x * y * z
	}
	bound := Bind4Of4(fn, 1, 2, 3, 4)
	result := bound()
	expected := 24
	if result != expected {
		t.Errorf("Bind4Of4 result: got %v, want %v", result, expected)
	}
}

// TestBind1Of5 tests the Bind1Of5 function
func TestBind1Of5(t *testing.T) {
	fn := func(v, w, x, y, z int) int {
		return v + w + x + y + z
	}
	bound := Bind1Of5(fn, 1)
	result := bound(2, 3, 4, 5)
	expected := 15
	if result != expected {
		t.Errorf("Bind1Of5 result: got %v, want %v", result, expected)
	}
}

// TestBind2Of5 tests the Bind2Of5 function
func TestBind2Of5(t *testing.T) {
	fn := func(v, w, x, y, z int) int {
		return v + w + x + y + z
	}
	bound := Bind2Of5(fn, 1, 2)
	result := bound(3, 4, 5)
	expected := 15
	if result != expected {
		t.Errorf("Bind2Of5 result: got %v, want %v", result, expected)
	}
}

// TestBind3Of5 tests the Bind3Of5 function
func TestBind3Of5(t *testing.T) {
	fn := func(v, w, x, y, z int) int {
		return v + w + x + y + z
	}
	bound := Bind3Of5(fn, 1, 2, 3)
	result := bound(4, 5)
	expected := 15
	if result != expected {
		t.Errorf("Bind3Of5 result: got %v, want %v", result, expected)
	}
}

// TestBind4Of5 tests the Bind4Of5 function
func TestBind4Of5(t *testing.T) {
	fn := func(v, w, x, y, z int) int {
		return v + w + x + y + z
	}
	bound := Bind4Of5(fn, 1, 2, 3, 4)
	result := bound(5)
	expected := 15
	if result != expected {
		t.Errorf("Bind4Of5 result: got %v, want %v", result, expected)
	}
}

// TestBind5Of5 tests the Bind5Of5 function
func TestBind5Of5(t *testing.T) {
	fn := func(v, w, x, y, z int) int {
		return v * w * x * y * z
	}
	bound := Bind5Of5(fn, 1, 2, 3, 4, 5)
	result := bound()
	expected := 120
	if result != expected {
		t.Errorf("Bind5Of5 result: got %v, want %v", result, expected)
	}
}

// TestBind1Of6 tests the Bind1Of6 function
func TestBind1Of6(t *testing.T) {
	fn := func(u, v, w, x, y, z int) int {
		return u + v + w + x + y + z
	}
	bound := Bind1Of6(fn, 1)
	result := bound(2, 3, 4, 5, 6)
	expected := 21
	if result != expected {
		t.Errorf("Bind1Of6 result: got %v, want %v", result, expected)
	}
}

// TestBind2Of6 tests the Bind2Of6 function
func TestBind2Of6(t *testing.T) {
	fn := func(u, v, w, x, y, z int) int {
		return u + v + w + x + y + z
	}
	bound := Bind2Of6(fn, 1, 2)
	result := bound(3, 4, 5, 6)
	expected := 21
	if result != expected {
		t.Errorf("Bind2Of6 result: got %v, want %v", result, expected)
	}
}

// TestBind3Of6 tests the Bind3Of6 function
func TestBind3Of6(t *testing.T) {
	fn := func(u, v, w, x, y, z int) int {
		return u + v + w + x + y + z
	}
	bound := Bind3Of6(fn, 1, 2, 3)
	result := bound(4, 5, 6)
	expected := 21
	if result != expected {
		t.Errorf("Bind3Of6 result: got %v, want %v", result, expected)
	}
}

// TestBind4Of6 tests the Bind4Of6 function
func TestBind4Of6(t *testing.T) {
	fn := func(u, v, w, x, y, z int) int {
		return u + v + w + x + y + z
	}
	bound := Bind4Of6(fn, 1, 2, 3, 4)
	result := bound(5, 6)
	expected := 21
	if result != expected {
		t.Errorf("Bind4Of6 result: got %v, want %v", result, expected)
	}
}

// TestBind5Of6 tests the Bind5Of6 function
func TestBind5Of6(t *testing.T) {
	fn := func(u, v, w, x, y, z int) int {
		return u + v + w + x + y + z
	}
	bound := Bind5Of6(fn, 1, 2, 3, 4, 5)
	result := bound(6)
	expected := 21
	if result != expected {
		t.Errorf("Bind5Of6 result: got %v, want %v", result, expected)
	}
}

// TestBind6Of6 tests the Bind6Of6 function
func TestBind6Of6(t *testing.T) {
	fn := func(u, v, w, x, y, z int) int {
		return u * v * w * x * y * z
	}
	bound := Bind6Of6(fn, 1, 2, 3, 4, 5, 6)
	result := bound()
	expected := 720
	if result != expected {
		t.Errorf("Bind6Of6 result: got %v, want %v", result, expected)
	}
}

// TestBind1Of7 tests the Bind1Of7 function
func TestBind1Of7(t *testing.T) {
	fn := func(t, u, v, w, x, y, z int) int {
		return t + u + v + w + x + y + z
	}
	bound := Bind1Of7(fn, 1)
	result := bound(2, 3, 4, 5, 6, 7)
	expected := 28
	if result != expected {
		t.Errorf("Bind1Of7 result: got %v, want %v", result, expected)
	}
}

// TestBind2Of7 tests the Bind2Of7 function
func TestBind2Of7(t *testing.T) {
	fn := func(t, u, v, w, x, y, z int) int {
		return t + u + v + w + x + y + z
	}
	bound := Bind2Of7(fn, 1, 2)
	result := bound(3, 4, 5, 6, 7)
	expected := 28
	if result != expected {
		t.Errorf("Bind2Of7 result: got %v, want %v", result, expected)
	}
}

// TestBind3Of7 tests the Bind3Of7 function
func TestBind3Of7(t *testing.T) {
	fn := func(t, u, v, w, x, y, z int) int {
		return t + u + v + w + x + y + z
	}
	bound := Bind3Of7(fn, 1, 2, 3)
	result := bound(4, 5, 6, 7)
	expected := 28
	if result != expected {
		t.Errorf("Bind3Of7 result: got %v, want %v", result, expected)
	}
}

// TestBind4Of7 tests the Bind4Of7 function
func TestBind4Of7(t *testing.T) {
	fn := func(t, u, v, w, x, y, z int) int {
		return t + u + v + w + x + y + z
	}
	bound := Bind4Of7(fn, 1, 2, 3, 4)
	result := bound(5, 6, 7)
	expected := 28
	if result != expected {
		t.Errorf("Bind4Of7 result: got %v, want %v", result, expected)
	}
}

// TestBind5Of7 tests the Bind5Of7 function
func TestBind5Of7(t *testing.T) {
	fn := func(t, u, v, w, x, y, z int) int {
		return t + u + v + w + x + y + z
	}
	bound := Bind5Of7(fn, 1, 2, 3, 4, 5)
	result := bound(6, 7)
	expected := 28
	if result != expected {
		t.Errorf("Bind5Of7 result: got %v, want %v", result, expected)
	}
}

// TestBind6Of7 tests the Bind6Of7 function
func TestBind6Of7(t *testing.T) {
	fn := func(t, u, v, w, x, y, z int) int {
		return t + u + v + w + x + y + z
	}
	bound := Bind6Of7(fn, 1, 2, 3, 4, 5, 6)
	result := bound(7)
	expected := 28
	if result != expected {
		t.Errorf("Bind6Of7 result: got %v, want %v", result, expected)
	}
}

// TestBind7Of7 tests the Bind7Of7 function
func TestBind7Of7(t *testing.T) {
	fn := func(t, u, v, w, x, y, z int) int {
		return t * u * v * w * x * y * z
	}
	bound := Bind7Of7(fn, 1, 2, 3, 4, 5, 6, 7)
	result := bound()
	expected := 5040
	if result != expected {
		t.Errorf("Bind7Of7 result: got %v, want %v", result, expected)
	}
}

// TestBind1Of8 tests the Bind1Of8 function
func TestBind1Of8(t *testing.T) {
	fn := func(s, t, u, v, w, x, y, z int) int {
		return s + t + u + v + w + x + y + z
	}
	bound := Bind1Of8(fn, 1)
	result := bound(2, 3, 4, 5, 6, 7, 8)
	expected := 36
	if result != expected {
		t.Errorf("Bind1Of8 result: got %v, want %v", result, expected)
	}
}

// TestBind2Of8 tests the Bind2Of8 function
func TestBind2Of8(t *testing.T) {
	fn := func(s, t, u, v, w, x, y, z int) int {
		return s + t + u + v + w + x + y + z
	}
	bound := Bind2Of8(fn, 1, 2)
	result := bound(3, 4, 5, 6, 7, 8)
	expected := 36
	if result != expected {
		t.Errorf("Bind2Of8 result: got %v, want %v", result, expected)
	}
}

// TestBind3Of8 tests the Bind3Of8 function
func TestBind3Of8(t *testing.T) {
	fn := func(s, t, u, v, w, x, y, z int) int {
		return s + t + u + v + w + x + y + z
	}
	bound := Bind3Of8(fn, 1, 2, 3)
	result := bound(4, 5, 6, 7, 8)
	expected := 36
	if result != expected {
		t.Errorf("Bind3Of8 result: got %v, want %v", result, expected)
	}
}

// TestBind4Of8 tests the Bind4Of8 function
func TestBind4Of8(t *testing.T) {
	fn := func(s, t, u, v, w, x, y, z int) int {
		return s + t + u + v + w + x + y + z
	}
	bound := Bind4Of8(fn, 1, 2, 3, 4)
	result := bound(5, 6, 7, 8)
	expected := 36
	if result != expected {
		t.Errorf("Bind4Of8 result: got %v, want %v", result, expected)
	}
}

// TestBind5Of8 tests the Bind5Of8 function
func TestBind5Of8(t *testing.T) {
	fn := func(s, t, u, v, w, x, y, z int) int {
		return s + t + u + v + w + x + y + z
	}
	bound := Bind5Of8(fn, 1, 2, 3, 4, 5)
	result := bound(6, 7, 8)
	expected := 36
	if result != expected {
		t.Errorf("Bind5Of8 result: got %v, want %v", result, expected)
	}
}

// TestBind6Of8 tests the Bind6Of8 function
func TestBind6Of8(t *testing.T) {
	fn := func(s, t, u, v, w, x, y, z int) int {
		return s + t + u + v + w + x + y + z
	}
	bound := Bind6Of8(fn, 1, 2, 3, 4, 5, 6)
	result := bound(7, 8)
	expected := 36
	if result != expected {
		t.Errorf("Bind6Of8 result: got %v, want %v", result, expected)
	}
}

// TestBind7Of8 tests the Bind7Of8 function
func TestBind7Of8(t *testing.T) {
	fn := func(s, t, u, v, w, x, y, z int) int {
		return s + t + u + v + w + x + y + z
	}
	bound := Bind7Of8(fn, 1, 2, 3, 4, 5, 6, 7)
	result := bound(8)
	expected := 36
	if result != expected {
		t.Errorf("Bind7Of8 result: got %v, want %v", result, expected)
	}
}

// TestBind8Of8 tests the Bind8Of8 function
func TestBind8Of8(t *testing.T) {
	fn := func(s, t, u, v, w, x, y, z int) int {
		return s * t * u * v * w * x * y * z
	}
	bound := Bind8Of8(fn, 1, 2, 3, 4, 5, 6, 7, 8)
	result := bound()
	expected := 40320
	if result != expected {
		t.Errorf("Bind8Of8 result: got %v, want %v", result, expected)
	}
}

// TestBind1Of9 tests the Bind1Of9 function
func TestBind1Of9(t *testing.T) {
	fn := func(r, s, t, u, v, w, x, y, z int) int {
		return r + s + t + u + v + w + x + y + z
	}
	bound := Bind1Of9(fn, 1)
	result := bound(2, 3, 4, 5, 6, 7, 8, 9)
	expected := 45
	if result != expected {
		t.Errorf("Bind1Of9 result: got %v, want %v", result, expected)
	}
}

// TestBind2Of9 tests the Bind2Of9 function
func TestBind2Of9(t *testing.T) {
	fn := func(r, s, t, u, v, w, x, y, z int) int {
		return r + s + t + u + v + w + x + y + z
	}
	bound := Bind2Of9(fn, 1, 2)
	result := bound(3, 4, 5, 6, 7, 8, 9)
	expected := 45
	if result != expected {
		t.Errorf("Bind2Of9 result: got %v, want %v", result, expected)
	}
}

// TestBind3Of9 tests the Bind3Of9 function
func TestBind3Of9(t *testing.T) {
	fn := func(r, s, t, u, v, w, x, y, z int) int {
		return r + s + t + u + v + w + x + y + z
	}
	bound := Bind3Of9(fn, 1, 2, 3)
	result := bound(4, 5, 6, 7, 8, 9)
	expected := 45
	if result != expected {
		t.Errorf("Bind3Of9 result: got %v, want %v", result, expected)
	}
}

// TestBind4Of9 tests the Bind4Of9 function
func TestBind4Of9(t *testing.T) {
	fn := func(r, s, t, u, v, w, x, y, z int) int {
		return r + s + t + u + v + w + x + y + z
	}
	bound := Bind4Of9(fn, 1, 2, 3, 4)
	result := bound(5, 6, 7, 8, 9)
	expected := 45
	if result != expected {
		t.Errorf("Bind4Of9 result: got %v, want %v", result, expected)
	}
}

// TestBind5Of9 tests the Bind5Of9 function
func TestBind5Of9(t *testing.T) {
	fn := func(r, s, t, u, v, w, x, y, z int) int {
		return r + s + t + u + v + w + x + y + z
	}
	bound := Bind5Of9(fn, 1, 2, 3, 4, 5)
	result := bound(6, 7, 8, 9)
	expected := 45
	if result != expected {
		t.Errorf("Bind5Of9 result: got %v, want %v", result, expected)
	}
}

// TestBind6Of9 tests the Bind6Of9 function
func TestBind6Of9(t *testing.T) {
	fn := func(r, s, t, u, v, w, x, y, z int) int {
		return r + s + t + u + v + w + x + y + z
	}
	bound := Bind6Of9(fn, 1, 2, 3, 4, 5, 6)
	result := bound(7, 8, 9)
	expected := 45
	if result != expected {
		t.Errorf("Bind6Of9 result: got %v, want %v", result, expected)
	}
}

// TestBind7Of9 tests the Bind7Of9 function
func TestBind7Of9(t *testing.T) {
	fn := func(r, s, t, u, v, w, x, y, z int) int {
		return r + s + t + u + v + w + x + y + z
	}
	bound := Bind7Of9(fn, 1, 2, 3, 4, 5, 6, 7)
	result := bound(8, 9)
	expected := 45
	if result != expected {
		t.Errorf("Bind7Of9 result: got %v, want %v", result, expected)
	}
}

// TestBind8Of9 tests the Bind8Of9 function
func TestBind8Of9(t *testing.T) {
	fn := func(r, s, t, u, v, w, x, y, z int) int {
		return r + s + t + u + v + w + x + y + z
	}
	bound := Bind8Of9(fn, 1, 2, 3, 4, 5, 6, 7, 8)
	result := bound(9)
	expected := 45
	if result != expected {
		t.Errorf("Bind8Of9 result: got %v, want %v", result, expected)
	}
}

// TestBind9Of9 tests the Bind9Of9 function
func TestBind9Of9(t *testing.T) {
	fn := func(r, s, t, u, v, w, x, y, z int) int {
		return r * s * t * u * v * w * x * y * z
	}
	bound := Bind9Of9(fn, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	result := bound()
	expected := 362880
	if result != expected {
		t.Errorf("Bind9Of9 result: got %v, want %v", result, expected)
	}
}

// TestBind1Of10 tests the Bind1Of10 function
func TestBind1Of10(t *testing.T) {
	fn := func(q, r, s, t, u, v, w, x, y, z int) int {
		return q + r + s + t + u + v + w + x + y + z
	}
	bound := Bind1Of10(fn, 1)
	result := bound(2, 3, 4, 5, 6, 7, 8, 9, 10)
	expected := 55
	if result != expected {
		t.Errorf("Bind1Of10 result: got %v, want %v", result, expected)
	}
}

// TestBind2Of10 tests the Bind2Of10 function
func TestBind2Of10(t *testing.T) {
	fn := func(q, r, s, t, u, v, w, x, y, z int) int {
		return q + r + s + t + u + v + w + x + y + z
	}
	bound := Bind2Of10(fn, 1, 2)
	result := bound(3, 4, 5, 6, 7, 8, 9, 10)
	expected := 55
	if result != expected {
		t.Errorf("Bind2Of10 result: got %v, want %v", result, expected)
	}
}

// TestBind3Of10 tests the Bind3Of10 function
func TestBind3Of10(t *testing.T) {
	fn := func(q, r, s, t, u, v, w, x, y, z int) int {
		return q + r + s + t + u + v + w + x + y + z
	}
	bound := Bind3Of10(fn, 1, 2, 3)
	result := bound(4, 5, 6, 7, 8, 9, 10)
	expected := 55
	if result != expected {
		t.Errorf("Bind3Of10 result: got %v, want %v", result, expected)
	}
}

// TestBind4Of10 tests the Bind4Of10 function
func TestBind4Of10(t *testing.T) {
	fn := func(q, r, s, t, u, v, w, x, y, z int) int {
		return q + r + s + t + u + v + w + x + y + z
	}
	bound := Bind4Of10(fn, 1, 2, 3, 4)
	result := bound(5, 6, 7, 8, 9, 10)
	expected := 55
	if result != expected {
		t.Errorf("Bind4Of10 result: got %v, want %v", result, expected)
	}
}

// TestBind5Of10 tests the Bind5Of10 function
func TestBind5Of10(t *testing.T) {
	fn := func(q, r, s, t, u, v, w, x, y, z int) int {
		return q + r + s + t + u + v + w + x + y + z
	}
	bound := Bind5Of10(fn, 1, 2, 3, 4, 5)
	result := bound(6, 7, 8, 9, 10)
	expected := 55
	if result != expected {
		t.Errorf("Bind5Of10 result: got %v, want %v", result, expected)
	}
}

// TestBind6Of10 tests the Bind6Of10 function
func TestBind6Of10(t *testing.T) {
	fn := func(q, r, s, t, u, v, w, x, y, z int) int {
		return q + r + s + t + u + v + w + x + y + z
	}
	bound := Bind6Of10(fn, 1, 2, 3, 4, 5, 6)
	result := bound(7, 8, 9, 10)
	expected := 55
	if result != expected {
		t.Errorf("Bind6Of10 result: got %v, want %v", result, expected)
	}
}

// TestBind7Of10 tests the Bind7Of10 function
func TestBind7Of10(t *testing.T) {
	fn := func(q, r, s, t, u, v, w, x, y, z int) int {
		return q + r + s + t + u + v + w + x + y + z
	}
	bound := Bind7Of10(fn, 1, 2, 3, 4, 5, 6, 7)
	result := bound(8, 9, 10)
	expected := 55
	if result != expected {
		t.Errorf("Bind7Of10 result: got %v, want %v", result, expected)
	}
}

// TestBind8Of10 tests the Bind8Of10 function
func TestBind8Of10(t *testing.T) {
	fn := func(q, r, s, t, u, v, w, x, y, z int) int {
		return q + r + s + t + u + v + w + x + y + z
	}
	bound := Bind8Of10(fn, 1, 2, 3, 4, 5, 6, 7, 8)
	result := bound(9, 10)
	expected := 55
	if result != expected {
		t.Errorf("Bind8Of10 result: got %v, want %v", result, expected)
	}
}

// TestBind9Of10 tests the Bind9Of10 function
func TestBind9Of10(t *testing.T) {
	fn := func(q, r, s, t, u, v, w, x, y, z int) int {
		return q + r + s + t + u + v + w + x + y + z
	}
	bound := Bind9Of10(fn, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	result := bound(10)
	expected := 55
	if result != expected {
		t.Errorf("Bind9Of10 result: got %v, want %v", result, expected)
	}
}

// TestBind10Of10 tests the Bind10Of10 function
func TestBind10Of10(t *testing.T) {
	fn := func(q, r, s, t, u, v, w, x, y, z int) int {
		return q * r * s * t * u * v * w * x * y * z
	}
	bound := Bind10Of10(fn, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	result := bound()
	expected := 3628800
	if result != expected {
		t.Errorf("Bind10Of10 result: got %v, want %v", result, expected)
	}
}
