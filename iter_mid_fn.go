package iter

// because go can't use generic on method
// so i have to peek it out, and make it to a function

// 转换函数
func Convert[T any, V any](iter Iterable[T], fn func(T) V) Iterable[V] {
	ch := make(chan V)
	go func() {
		for e := range iter.Iter() {
			ch <- fn(e)
		}
		close(ch)
	}()
	return &Iterator[V]{C: ch}
}

func Zip[T any, V any](iter1 Iterable[T], iter2 Iterable[V]) Iterable[struct {
	T T
	V V
}] {
	ch := make(chan struct {
		T T
		V V
	})
	go func() {
		for t := range iter1.Iter() {
			var v, ok = <-iter2.Iter()
			if !ok {
				break
			}
			ch <- struct {
				T T
				V V
			}{
				T: t, V: v,
			}
		}
		// 消耗完所有的 iter2, 避免协程堵塞
		for range iter2.Iter() {
		}
		close(ch)
	}()
	return &Iterator[struct {
		T T
		V V
	}]{
		C: ch,
	}
}
