package gounctional


// Bind1Of1 binds the first 1 arguments of the function f.
func Bind1Of1[T0 any, R any](f func(T0) R, arg0 T0) func() R {
	return func() R {
		return f(arg0)
	}
}

// Bind1Of2 binds the first 1 arguments of the function f.
func Bind1Of2[T0 any, T1 any, R any](f func(T0, T1) R, arg0 T0) func(T1) R {
	return func(arg1 T1) R {
		return f(arg0, arg1)
	}
}

// Bind2Of2 binds the first 2 arguments of the function f.
func Bind2Of2[T0 any, T1 any, R any](f func(T0, T1) R, arg0 T0, arg1 T1) func() R {
	return func() R {
		return f(arg0, arg1)
	}
}

// Bind1Of3 binds the first 1 arguments of the function f.
func Bind1Of3[T0 any, T1 any, T2 any, R any](f func(T0, T1, T2) R, arg0 T0) func(T1, T2) R {
	return func(arg1 T1, arg2 T2) R {
		return f(arg0, arg1, arg2)
	}
}

// Bind2Of3 binds the first 2 arguments of the function f.
func Bind2Of3[T0 any, T1 any, T2 any, R any](f func(T0, T1, T2) R, arg0 T0, arg1 T1) func(T2) R {
	return func(arg2 T2) R {
		return f(arg0, arg1, arg2)
	}
}

// Bind3Of3 binds the first 3 arguments of the function f.
func Bind3Of3[T0 any, T1 any, T2 any, R any](f func(T0, T1, T2) R, arg0 T0, arg1 T1, arg2 T2) func() R {
	return func() R {
		return f(arg0, arg1, arg2)
	}
}

// Bind1Of4 binds the first 1 arguments of the function f.
func Bind1Of4[T0 any, T1 any, T2 any, T3 any, R any](f func(T0, T1, T2, T3) R, arg0 T0) func(T1, T2, T3) R {
	return func(arg1 T1, arg2 T2, arg3 T3) R {
		return f(arg0, arg1, arg2, arg3)
	}
}

// Bind2Of4 binds the first 2 arguments of the function f.
func Bind2Of4[T0 any, T1 any, T2 any, T3 any, R any](f func(T0, T1, T2, T3) R, arg0 T0, arg1 T1) func(T2, T3) R {
	return func(arg2 T2, arg3 T3) R {
		return f(arg0, arg1, arg2, arg3)
	}
}

// Bind3Of4 binds the first 3 arguments of the function f.
func Bind3Of4[T0 any, T1 any, T2 any, T3 any, R any](f func(T0, T1, T2, T3) R, arg0 T0, arg1 T1, arg2 T2) func(T3) R {
	return func(arg3 T3) R {
		return f(arg0, arg1, arg2, arg3)
	}
}

// Bind4Of4 binds the first 4 arguments of the function f.
func Bind4Of4[T0 any, T1 any, T2 any, T3 any, R any](f func(T0, T1, T2, T3) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3) func() R {
	return func() R {
		return f(arg0, arg1, arg2, arg3)
	}
}

// Bind1Of5 binds the first 1 arguments of the function f.
func Bind1Of5[T0 any, T1 any, T2 any, T3 any, T4 any, R any](f func(T0, T1, T2, T3, T4) R, arg0 T0) func(T1, T2, T3, T4) R {
	return func(arg1 T1, arg2 T2, arg3 T3, arg4 T4) R {
		return f(arg0, arg1, arg2, arg3, arg4)
	}
}

// Bind2Of5 binds the first 2 arguments of the function f.
func Bind2Of5[T0 any, T1 any, T2 any, T3 any, T4 any, R any](f func(T0, T1, T2, T3, T4) R, arg0 T0, arg1 T1) func(T2, T3, T4) R {
	return func(arg2 T2, arg3 T3, arg4 T4) R {
		return f(arg0, arg1, arg2, arg3, arg4)
	}
}

// Bind3Of5 binds the first 3 arguments of the function f.
func Bind3Of5[T0 any, T1 any, T2 any, T3 any, T4 any, R any](f func(T0, T1, T2, T3, T4) R, arg0 T0, arg1 T1, arg2 T2) func(T3, T4) R {
	return func(arg3 T3, arg4 T4) R {
		return f(arg0, arg1, arg2, arg3, arg4)
	}
}

// Bind4Of5 binds the first 4 arguments of the function f.
func Bind4Of5[T0 any, T1 any, T2 any, T3 any, T4 any, R any](f func(T0, T1, T2, T3, T4) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3) func(T4) R {
	return func(arg4 T4) R {
		return f(arg0, arg1, arg2, arg3, arg4)
	}
}

// Bind5Of5 binds the first 5 arguments of the function f.
func Bind5Of5[T0 any, T1 any, T2 any, T3 any, T4 any, R any](f func(T0, T1, T2, T3, T4) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4) func() R {
	return func() R {
		return f(arg0, arg1, arg2, arg3, arg4)
	}
}

// Bind1Of6 binds the first 1 arguments of the function f.
func Bind1Of6[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, R any](f func(T0, T1, T2, T3, T4, T5) R, arg0 T0) func(T1, T2, T3, T4, T5) R {
	return func(arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

// Bind2Of6 binds the first 2 arguments of the function f.
func Bind2Of6[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, R any](f func(T0, T1, T2, T3, T4, T5) R, arg0 T0, arg1 T1) func(T2, T3, T4, T5) R {
	return func(arg2 T2, arg3 T3, arg4 T4, arg5 T5) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

// Bind3Of6 binds the first 3 arguments of the function f.
func Bind3Of6[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, R any](f func(T0, T1, T2, T3, T4, T5) R, arg0 T0, arg1 T1, arg2 T2) func(T3, T4, T5) R {
	return func(arg3 T3, arg4 T4, arg5 T5) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

// Bind4Of6 binds the first 4 arguments of the function f.
func Bind4Of6[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, R any](f func(T0, T1, T2, T3, T4, T5) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3) func(T4, T5) R {
	return func(arg4 T4, arg5 T5) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

// Bind5Of6 binds the first 5 arguments of the function f.
func Bind5Of6[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, R any](f func(T0, T1, T2, T3, T4, T5) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4) func(T5) R {
	return func(arg5 T5) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

// Bind6Of6 binds the first 6 arguments of the function f.
func Bind6Of6[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, R any](f func(T0, T1, T2, T3, T4, T5) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5) func() R {
	return func() R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5)
	}
}

// Bind1Of7 binds the first 1 arguments of the function f.
func Bind1Of7[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, R any](f func(T0, T1, T2, T3, T4, T5, T6) R, arg0 T0) func(T1, T2, T3, T4, T5, T6) R {
	return func(arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

// Bind2Of7 binds the first 2 arguments of the function f.
func Bind2Of7[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, R any](f func(T0, T1, T2, T3, T4, T5, T6) R, arg0 T0, arg1 T1) func(T2, T3, T4, T5, T6) R {
	return func(arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

// Bind3Of7 binds the first 3 arguments of the function f.
func Bind3Of7[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, R any](f func(T0, T1, T2, T3, T4, T5, T6) R, arg0 T0, arg1 T1, arg2 T2) func(T3, T4, T5, T6) R {
	return func(arg3 T3, arg4 T4, arg5 T5, arg6 T6) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

// Bind4Of7 binds the first 4 arguments of the function f.
func Bind4Of7[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, R any](f func(T0, T1, T2, T3, T4, T5, T6) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3) func(T4, T5, T6) R {
	return func(arg4 T4, arg5 T5, arg6 T6) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

// Bind5Of7 binds the first 5 arguments of the function f.
func Bind5Of7[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, R any](f func(T0, T1, T2, T3, T4, T5, T6) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4) func(T5, T6) R {
	return func(arg5 T5, arg6 T6) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

// Bind6Of7 binds the first 6 arguments of the function f.
func Bind6Of7[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, R any](f func(T0, T1, T2, T3, T4, T5, T6) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5) func(T6) R {
	return func(arg6 T6) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

// Bind7Of7 binds the first 7 arguments of the function f.
func Bind7Of7[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, R any](f func(T0, T1, T2, T3, T4, T5, T6) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6) func() R {
	return func() R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	}
}

// Bind1Of8 binds the first 1 arguments of the function f.
func Bind1Of8[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7) R, arg0 T0) func(T1, T2, T3, T4, T5, T6, T7) R {
	return func(arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

// Bind2Of8 binds the first 2 arguments of the function f.
func Bind2Of8[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7) R, arg0 T0, arg1 T1) func(T2, T3, T4, T5, T6, T7) R {
	return func(arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

// Bind3Of8 binds the first 3 arguments of the function f.
func Bind3Of8[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7) R, arg0 T0, arg1 T1, arg2 T2) func(T3, T4, T5, T6, T7) R {
	return func(arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

// Bind4Of8 binds the first 4 arguments of the function f.
func Bind4Of8[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3) func(T4, T5, T6, T7) R {
	return func(arg4 T4, arg5 T5, arg6 T6, arg7 T7) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

// Bind5Of8 binds the first 5 arguments of the function f.
func Bind5Of8[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4) func(T5, T6, T7) R {
	return func(arg5 T5, arg6 T6, arg7 T7) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

// Bind6Of8 binds the first 6 arguments of the function f.
func Bind6Of8[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5) func(T6, T7) R {
	return func(arg6 T6, arg7 T7) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

// Bind7Of8 binds the first 7 arguments of the function f.
func Bind7Of8[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6) func(T7) R {
	return func(arg7 T7) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

// Bind8Of8 binds the first 8 arguments of the function f.
func Bind8Of8[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7) func() R {
	return func() R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
}

// Bind1Of9 binds the first 1 arguments of the function f.
func Bind1Of9[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8) R, arg0 T0) func(T1, T2, T3, T4, T5, T6, T7, T8) R {
	return func(arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7, arg8 T8) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
}

// Bind2Of9 binds the first 2 arguments of the function f.
func Bind2Of9[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8) R, arg0 T0, arg1 T1) func(T2, T3, T4, T5, T6, T7, T8) R {
	return func(arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7, arg8 T8) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
}

// Bind3Of9 binds the first 3 arguments of the function f.
func Bind3Of9[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8) R, arg0 T0, arg1 T1, arg2 T2) func(T3, T4, T5, T6, T7, T8) R {
	return func(arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7, arg8 T8) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
}

// Bind4Of9 binds the first 4 arguments of the function f.
func Bind4Of9[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3) func(T4, T5, T6, T7, T8) R {
	return func(arg4 T4, arg5 T5, arg6 T6, arg7 T7, arg8 T8) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
}

// Bind5Of9 binds the first 5 arguments of the function f.
func Bind5Of9[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4) func(T5, T6, T7, T8) R {
	return func(arg5 T5, arg6 T6, arg7 T7, arg8 T8) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
}

// Bind6Of9 binds the first 6 arguments of the function f.
func Bind6Of9[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5) func(T6, T7, T8) R {
	return func(arg6 T6, arg7 T7, arg8 T8) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
}

// Bind7Of9 binds the first 7 arguments of the function f.
func Bind7Of9[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6) func(T7, T8) R {
	return func(arg7 T7, arg8 T8) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
}

// Bind8Of9 binds the first 8 arguments of the function f.
func Bind8Of9[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7) func(T8) R {
	return func(arg8 T8) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
}

// Bind9Of9 binds the first 9 arguments of the function f.
func Bind9Of9[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7, arg8 T8) func() R {
	return func() R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
}

// Bind1Of10 binds the first 1 arguments of the function f.
func Bind1Of10[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) R, arg0 T0) func(T1, T2, T3, T4, T5, T6, T7, T8, T9) R {
	return func(arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7, arg8 T8, arg9 T9) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	}
}

// Bind2Of10 binds the first 2 arguments of the function f.
func Bind2Of10[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) R, arg0 T0, arg1 T1) func(T2, T3, T4, T5, T6, T7, T8, T9) R {
	return func(arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7, arg8 T8, arg9 T9) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	}
}

// Bind3Of10 binds the first 3 arguments of the function f.
func Bind3Of10[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) R, arg0 T0, arg1 T1, arg2 T2) func(T3, T4, T5, T6, T7, T8, T9) R {
	return func(arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7, arg8 T8, arg9 T9) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	}
}

// Bind4Of10 binds the first 4 arguments of the function f.
func Bind4Of10[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3) func(T4, T5, T6, T7, T8, T9) R {
	return func(arg4 T4, arg5 T5, arg6 T6, arg7 T7, arg8 T8, arg9 T9) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	}
}

// Bind5Of10 binds the first 5 arguments of the function f.
func Bind5Of10[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4) func(T5, T6, T7, T8, T9) R {
	return func(arg5 T5, arg6 T6, arg7 T7, arg8 T8, arg9 T9) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	}
}

// Bind6Of10 binds the first 6 arguments of the function f.
func Bind6Of10[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5) func(T6, T7, T8, T9) R {
	return func(arg6 T6, arg7 T7, arg8 T8, arg9 T9) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	}
}

// Bind7Of10 binds the first 7 arguments of the function f.
func Bind7Of10[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6) func(T7, T8, T9) R {
	return func(arg7 T7, arg8 T8, arg9 T9) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	}
}

// Bind8Of10 binds the first 8 arguments of the function f.
func Bind8Of10[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7) func(T8, T9) R {
	return func(arg8 T8, arg9 T9) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	}
}

// Bind9Of10 binds the first 9 arguments of the function f.
func Bind9Of10[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7, arg8 T8) func(T9) R {
	return func(arg9 T9) R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	}
}

// Bind10Of10 binds the first 10 arguments of the function f.
func Bind10Of10[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, R any](f func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) R, arg0 T0, arg1 T1, arg2 T2, arg3 T3, arg4 T4, arg5 T5, arg6 T6, arg7 T7, arg8 T8, arg9 T9) func() R {
	return func() R {
		return f(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	}
}
