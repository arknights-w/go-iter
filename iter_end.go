package iter

func (it *Iterator[T]) GroupByInt(fn func(t T) int) map[int][]T {
	groups := make(map[int][]T)
	for e := range it.Iter() {
		key := fn(e)
		groups[key] = append(groups[key], e)
	}
	return groups
}

func (it *Iterator[T]) GroupByStr(fn func(t T) string) map[string][]T {
	groups := make(map[string][]T)
	for e := range it.Iter() {
		key := fn(e)
		groups[key] = append(groups[key], e)
	}
	return groups
}

func (it *Iterator[T]) Count() int {
	var count = 0
	for range it.Iter() {
		count++
	}
	return count
}

func (it *Iterator[T]) Slice() []T {
	slice := make([]T, 0, 2)
	for e := range it.Iter() {
		slice = append(slice, e)
	}
	return slice
}

func (it *Iterator[T]) Clear() {
	for range it.Iter() {
	}
}
