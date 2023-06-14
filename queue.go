package iter

// queue
type queue[T any] struct {
	cup  int
	size int
	head *item[T]
	tail *item[T]
}

type item[T any] struct {
	Value T
	Next  *item[T]
}

func NewQueue[T any](cup int) *queue[T] {
	return &queue[T]{
		cup:  cup,
		size: 0,
	}
}

func (q *queue[T]) Push(val T) {
	var value *item[T]

	if q.size == q.cup {
		value = q.head
		q.head = q.head.Next
		value.Value = val
		value.Next = nil
	} else {
		q.size++
		value = &item[T]{Value: val}
	}
	if q.size == 1 {
		q.head = value
		q.tail = value
	} else {
		q.tail.Next = value
		q.tail = q.tail.Next
	}
}

// get the header and delete it from the list
func (q *queue[T]) Pop() T {
	var empty T
	if q.size > 1 {
	} else if q.size == 1 {
		q.tail = nil
	} else {
		return empty
	}
	pop := q.head
	q.head = q.head.Next
	q.size--
	return pop.Value
}

// get the header but don't delete it
func (q *queue[T]) Peek() T {
	return q.head.Value
}

// the queue is full(size == cup)
func (q *queue[T]) IsFull() bool {
	return q.size == q.cup
}

// the queue is empty(size == 0)
func (q *queue[T]) IsEmpty() bool {
	return q.size == 0
}

// turn to list
func (q *queue[T]) List() []T {
	list := make([]T, 0, q.size)
	ptr := q.head
	for ptr != nil {
		list = append(list, ptr.Value)
		ptr = ptr.Next
	}
	return list
}
