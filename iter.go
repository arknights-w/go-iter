package iter

/* there are also 4 function for Iterable interface,
 * witch in xxx_fn.go
 *
 * in iter_mid_fn.go:
 * Convert[T any, V any](iter Iterable[T], fn func(T) V) Iterable[V]
 * Zip[T any, V any](iter1 Iterable[T], iter2 Iterable[V]) Iterable[struct{T T;V V}]
 *
 * in iter_end_fn.go:
 * GroupBy[T any, V comparable](iter Iterable[T], fn func(t T) V) map[V][]T
 * CollectZip[T comparable, V any](iter Iterable[struct {T T;V V}]) map[T]V
 */

type Iterable[T any] interface {
	Iter() <-chan T
	/* mid method, turn iter*/

	Map(f func(T) T) Iterable[T]
	Filter(f func(T) bool) Iterable[T]
	Step(count int) Iterable[T]
	Head(count int) Iterable[T]
	Tail(count int) Iterable[T]
	Next() Iterable[T]
	Sort(less func(i, j T) bool) Iterable[T]
	Distinct() Iterable[T]
	DistinctWithFn(fn func(T) any) Iterable[T]
	Concat(iters ...Iterable[T]) Iterable[T]
	ConcatOne(iter Iterable[T]) Iterable[T]

	/* end method, turn collection */

	GroupByInt(fn func(t T) int) map[int][]T
	GroupByStr(fn func(t T) string) map[string][]T
	Count() int
	Slice() []T
	Clear()
}

type Iterator[T any] struct {
	C <-chan T
}

func NewIterator[T any](items []T) Iterable[T] {
	ch := make(chan T)
	go func() {
		for _, item := range items {
			ch <- item
		}
		close(ch)
	}()
	return &Iterator[T]{C: ch}
}

// consume one item in channel
func (it *Iterator[T]) Iter() <-chan T {
	return it.C
}
