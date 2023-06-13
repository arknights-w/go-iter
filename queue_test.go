package iter_test

import (
	"fmt"
	"iter"
	"testing"
)

func TestQueue(t *testing.T) {
	q := iter.NewQueue[int](3)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(6)
	i := q.List()
	fmt.Printf("i: %v\n", i)
}
