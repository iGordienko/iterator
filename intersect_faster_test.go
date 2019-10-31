package iterator

import (
	"fmt"
	"github.com/freepk/arrays"
	"testing"
)

func TestFasterIntersector(t *testing.T) {
	a0 := randArray(3000)
	a1 := randArray(2000)
	a0Copy := copyArray(a0)
	a1Copy := copyArray(a1)
	expectedIntersect := copyArray(arrays.Intersect(a0, a1))

	it := NewFasterIntersectionIterator(NewArrayIter(a0), NewArrayIter(a1))
	it.Reset()
	out := make([]int, 0)
	for {
		if v, ok := it.Next(); !ok {
			break
		} else {
			out = append(out, v)
		}
	}
	if !arrays.IsEqual(out, expectedIntersect) {
		fmt.Println("out")
		fmt.Println(out)
		fmt.Println("correct")
		fmt.Println(expectedIntersect)
		fmt.Println("a0")
		fmt.Println(a0Copy)
		fmt.Println("a1")
		fmt.Println(a1Copy)
		t.Fail()
	}
	if !arrays.IsEqual(a0, a0Copy) {
		t.Fail()
	}
	if !arrays.IsEqual(a1, a1Copy) {
		t.Fail()
	}
}

func TestFasterIntersector_NextSomeFor2RandomArrays(t *testing.T) {
	a0 := randArray(3000)
	a1 := randArray(2000)
	//a0 := randArray(8)
	//a1 := randArray(5)
	//a0 := randArray(600)
	//a1 := randArray(400)
	//a0 := []int{1, 2, 4, 5, 13}
	//a1 := []int{4, 5, 6, 9}
	a0Copy := copyArray(a0)
	a1Copy := copyArray(a1)
	expectedIntersect := copyArray(arrays.Intersect(a0, a1))

	//fmt.Println("correct1")
	//fmt.Println(arrays.Intersect(a0, a1))

	it := NewFasterIntersectionIterator(NewArrayIter(a0), NewArrayIter(a1))
	it.Reset()
	out := make([]int, 0)
	for {
		if v, ok := it.NextSome(); !ok {
			break
		} else {
			out = append(out, v...)
		}
	}

	if !arrays.IsEqual(out, expectedIntersect) {
		fmt.Println("out")
		fmt.Println(out)
		fmt.Println("correct")
		fmt.Println(expectedIntersect)
		fmt.Println("a0")
		fmt.Println(a0Copy)
		fmt.Println("a1")
		fmt.Println(a1Copy)
		t.Fail()
	}
	if !arrays.IsEqual(a0, a0Copy) {
		t.Fail()
	}
	if !arrays.IsEqual(a1, a1Copy) {
		t.Fail()
	}
}

func TestFasterIntersector_SameArrayInFourIntersections(t *testing.T) {
	a0 := randArray(3000)
	a0Copy := copyArray(a0)
	//a0 := []int{0, 1, 3}
	//a1 := []int{1, 3}
	expectedIntersect := copyArray(arrays.Intersect(a0, a0))
	it := NewFasterIntersectionIterator(NewFasterIntersectionIterator(NewArrayIter(a0), NewArrayIter(a0)), NewFasterIntersectionIterator(NewArrayIter(a0), NewArrayIter(a0)))
	it.Reset()
	out := make([]int, 0)
	for {
		if v, ok := it.NextSome(); !ok {
			break
		} else {
			out = append(out, v...)
		}
	}
	if !arrays.IsEqual(out, expectedIntersect) {
		fmt.Println("out")
		fmt.Println(out)
		fmt.Println("correct")
		fmt.Println(expectedIntersect)
		t.Fail()
	}
	if !arrays.IsEqual(a0, a0Copy) {
		t.Fail()
	}
}

func TestFasterIntersector_DoubleIntercetionOfTwoArrays(t *testing.T) {
	a0 := randArray(3000)
	a1 := randArray(2000)
	a0Copy := copyArray(a0)
	a1Copy := copyArray(a1)
	//a0 := []int{0, 1, 3}
	//a1 := []int{1, 3}
	expectedIntersect := copyArray(arrays.Intersect(a0, a1))
	it := NewFasterIntersectionIterator(NewFasterIntersectionIterator(NewArrayIter(a0), NewArrayIter(a1)), NewFasterIntersectionIterator(NewArrayIter(a0), NewArrayIter(a1)))
	it.Reset()
	out := make([]int, 0)
	for {
		if v, ok := it.NextSome(); !ok {
			break
		} else {
			out = append(out, v...)
		}
	}

	if !arrays.IsEqual(out, expectedIntersect) {
		fmt.Println("out")
		fmt.Println(out)
		fmt.Println("correct")
		fmt.Println(expectedIntersect)
		t.Fail()
	}
	if !arrays.IsEqual(a0, a0Copy) {
		t.Fail()
	}
	if !arrays.IsEqual(a1, a1Copy) {
		t.Fail()
	}
}

func TestFasterIntersector_BigThreeArrays(t *testing.T) {
	a0 := randArray(600000)
	a1 := randArray(400000)
	a2 := randArray(200000)
	a0Copy := copyArray(a0)
	a1Copy := copyArray(a1)
	a2Copy := copyArray(a2)
	a01 := copyArray(arrays.Intersect(a0, a1))
	expectedIntersect := copyArray(arrays.Intersect(a01, a2))

	it := NewFasterIntersectionIterator(NewFasterIntersectionIterator(NewArrayIter(a0), NewArrayIter(a1)), NewArrayIter(a2))
	out := make([]int, 0)
	for {
		if v, ok := it.NextSome(); !ok {
			break
		} else {
			out = append(out, v...)
		}
	}

	if !arrays.IsEqual(out, expectedIntersect) {
		t.Fail()
	}
	if !arrays.IsEqual(a0, a0Copy) {
		t.Fail()
	}
	if !arrays.IsEqual(a1, a1Copy) {
		t.Fail()
	}
	if !arrays.IsEqual(a2, a2Copy) {
		t.Fail()
	}
}

func BenchmarkFasterIntersector_Next(b *testing.B) {
	a0 := randArray(300000)
	a1 := randArray(200000)
	a2 := randArray(100000)
	it := NewFasterIntersectionIterator(NewFasterIntersectionIterator(NewArrayIter(a0), NewArrayIter(a1)), NewArrayIter(a2))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			if _, ok := it.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkFasterIntersector_NextSome(b *testing.B) {
	a0 := randArray(300000)
	a1 := randArray(200000)
	a2 := randArray(100000)
	it := NewFasterIntersectionIterator(NewFasterIntersectionIterator(NewArrayIter(a0), NewArrayIter(a1)), NewArrayIter(a2))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			buf, ok := it.NextSome()
			if !ok {
				break
			}
			for range buf {

			}
		}
	}
}
