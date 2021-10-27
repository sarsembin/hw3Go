package fibonacci

func Fibonacci() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		fNext := f1 + f2
		f1 = f2
		f2 = fNext
		return fNext
	}
}
