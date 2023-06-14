package iter_test

import (
	"fmt"
	"go-iter"

	"testing"
)

func TestOtherThing(t *testing.T) {
	var list = []int{1, 2, 3}
	var i = 5
	list = list[len(list)-i:]
	fmt.Printf("list: %v\n", list)
}

func TestSimple(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	it := iter.NewIterator(items)
	// 2, 4, 6, 8, 10, 12, 14, 16, 18, 0
	res := it.Map(func(i int) int { return i * 2 }).
		// 2, 6, 8, 12, 14, 18, 0
		Filter(func(t int) bool { return t%3 != 1 }).
		// 12, 14, 18, 0
		Tail(4).
		// 0, 12, 14, 18
		Sort(func(i, j int) bool { return i < j }).
		Slice()

	fmt.Printf("res: %v\n", res)
}

func TestZip(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	items2 := []string{"abc", "cba", "bac", "bca", "acb", "cab"}
	i := iter.NewIterator(items)
	i2 := iter.NewIterator(items2)

	i3 := iter.Zip(i, i2)
	m := iter.CollectZip(i3)

	fmt.Printf("m: %v\n", m)
}

func TestGroupBy(t *testing.T) {
	items := []string{"abc", "cba", "bac", "bca", "acb", "cab"}
	it := iter.NewIterator(items)
	m := iter.GroupBy(it,
		func(str string) rune { var p = []rune(str); return p[0] })
	fmt.Printf("m: %v\n", m)
}

func TestConvert(t *testing.T) {
	items := []string{"abc", "cba", "bac", "bca", "acb", "cab"}
	it := iter.NewIterator(items)
	res := iter.Convert(it,
		func(str string) int { var p = []rune(str); return int(p[0]) }).
		Slice()
	fmt.Printf("res: %v\n", res)
}
