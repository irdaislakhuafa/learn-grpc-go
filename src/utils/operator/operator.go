package operator

func Ternary[T any](isOk bool, a, b T) T {
	if isOk {
		return a
	}
	return b
}
