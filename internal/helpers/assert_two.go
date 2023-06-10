package helpers

func AssertTwo[S any, T any](lhs any, rhs any) (S, T, bool) {
	val1, ok1 := lhs.(S)
	if !ok1 {
		return *new(S), *new(T), false
	}

	val2, ok2 := rhs.(T)
	if !ok2 {
		return *new(S), *new(T), false
	}

	return val1, val2, true
}
