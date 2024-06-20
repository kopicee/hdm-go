package functional

func RemoveDuplicates[T comparable](elems []T) []T {
	identityFunc := func(x T) T { return x }
	return RemoveDuplicatesBy(elems, identityFunc)
}

func RemoveDuplicatesBy[K comparable, T any](elems []T, cmp func(T) K) []T {
	ret := make([]T, 0)
	seen := make(map[K]bool)

	for _, elem := range elems {
		hash := cmp(elem)

		if _, ok := seen[hash]; ok {
			continue
		}

		seen[hash] = true
		ret = append(ret, elem)
	}

	return ret
}

func Map[T, U any](elems []T, mapFunc func(T) U) []U {
	ret := make([]U, len(elems))

	for i, elem := range elems {
		ret[i] = mapFunc(elem)
	}

	return ret
}

func Typecast[TSource, TTarget any](obj TSource) TTarget {
	return any(obj).(TTarget)
}

func Filter[T any](elems []T, predicate func(T) bool) []T {
	ret := make([]T, 0)
	for _, elem := range elems {
		if predicate(elem) {
			ret = append(ret, elem)
		}
	}
	return ret
}
