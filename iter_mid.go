package iter

import (
	"sort"
)

// traverse all
func (it *Iterator[T]) Map(f func(T) T) Iterable[T] {
	ch := make(chan T)
	go func() {
		for item := range it.Iter() {
			ch <- f(item)
		}
		close(ch)
	}()
	return &Iterator[T]{C: ch}
}

func (it *Iterator[T]) Filter(f func(T) bool) Iterable[T] {
	ch := make(chan T)
	go func() {
		for item := range it.Iter() {
			if f(item) {
				ch <- item
			}
		}
		close(ch)
	}()
	return &Iterator[T]{C: ch}
}

func (it *Iterator[T]) Step(count int) Iterable[T] {
	ch := make(chan T)
	go func() {
		var i = 0
		for item := range it.Iter() {
			if i++; i > count {
				ch <- item
			}
		}
		close(ch)
	}()
	return &Iterator[T]{C: ch}
}

func (it *Iterator[T]) Head(count int) Iterable[T] {
	ch := make(chan T)
	go func() {
		var i = 0
		for item := range it.Iter() {
			if i++; i <= count {
				ch <- item
			}
		}
		close(ch)
	}()
	return &Iterator[T]{C: ch}
}

func (it *Iterator[T]) Tail(count int) Iterable[T] {
	ch := make(chan T)
	queue := NewQueue[T](count)
	go func() {
		for item := range it.Iter() {
			queue.Push(item)
		}
		for !queue.IsEmpty() {
			val := queue.Pop()
			ch <- val
		}
		close(ch)
	}()
	return &Iterator[T]{C: ch}
}

func (it *Iterator[T]) Next() Iterable[T] {
	ch := make(chan T)
	go func() {
		var flag = false
		for item := range it.Iter() {
			if !flag {
				flag = true
				continue
			}
			ch <- item
		}
		close(ch)
	}()
	return &Iterator[T]{C: ch}
}

func (it *Iterator[T]) Sort(less func(i, j T) bool) Iterable[T] {
	ch := make(chan T)
	go func() {
		items := make([]T, 0, 16)
		for item := range it.Iter() {
			items = append(items, item)
		}

		// 对切片进行排序
		sort.Sort(sortableSlice[T]{
			items: items,
			less:  less,
		})

		for _, item := range items {
			ch <- item
		}
		close(ch)
	}()
	return &Iterator[T]{C: ch}
}

func (it *Iterator[T]) Distinct() Iterable[T] {
	ch := make(chan T)
	go func() {
		var maps = make(map[any]struct{})
		for item := range it.Iter() {
			maps[item] = struct{}{}
		}
		var item T
		for k := range maps {
			item = k.(T)
			ch <- item
		}
		close(ch)
	}()
	return &Iterator[T]{C: ch}
}

// If two inputs result in the same output,
// we consider it a duplicate
func (it *Iterator[T]) DistinctWithFn(fn func(T) any) Iterable[T] {
	ch := make(chan T)
	go func() {
		var maps = make(map[any]T)
		for item := range it.Iter() {
			key := fn(item)
			maps[key] = item
		}
		var item T
		for k := range maps {
			item = k.(T)
			ch <- item
		}
		close(ch)
	}()
	return &Iterator[T]{C: ch}
}

func (it *Iterator[T]) Concat(iters ...Iterable[T]) Iterable[T] {
	var _it Iterable[T] = it
	for _, iter := range iters {
		_it = _it.ConcatOne(iter)
	}
	return _it
}

func (it *Iterator[T]) ConcatOne(iter Iterable[T]) Iterable[T] {
	ch := make(chan T)
	go func() {
		for e := range it.Iter() {
			ch <- e
		}
		for e := range iter.Iter() {
			ch <- e
		}
		close(ch)
	}()
	return &Iterator[T]{C: ch}
}
