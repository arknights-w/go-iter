package iter

func GroupBy[T any, V comparable](iter Iterable[T], fn func(t T) V) map[V][]T {
	groups := make(map[V][]T)
	for e := range iter.Iter() {
		key := fn(e)
		groups[key] = append(groups[key], e)
	}
	return groups
}

func CollectZip[T comparable, V any](iter Iterable[struct {
	T T
	V V
}]) map[T]V {
	col := make(map[T]V)
	for e := range iter.Iter() {
		col[e.T] = e.V
	}
	return col
}
