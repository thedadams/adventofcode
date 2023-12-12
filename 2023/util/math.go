package util

func AbsoluteValue[T Number](a, b T) T {
	if a > b {
		return a - b
	}
	return b - a
}
