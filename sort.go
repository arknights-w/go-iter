package iter

type sortableSlice[T any] struct {
	items []T
	less  func(i, j T) bool
}

func (s sortableSlice[T]) Len() int {
	return len(s.items)
}

func (s sortableSlice[T]) Less(i, j int) bool {
	return s.less(s.items[i], s.items[j])
}

func (s sortableSlice[T]) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}
