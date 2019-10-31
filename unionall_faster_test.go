package iterator

import (
	"fmt"
	"github.com/freepk/arrays"
	"testing"
)

func TestFasterUnionAll(t *testing.T) {
	a0 := randArray(3000)
	a1 := randArray(2000)
	a0Copy := copyArray(a0)
	a1Copy := copyArray(a1)
	expectedUnionAll := copyArray(arrays.UnionAll(a0, a1))

	it := NewFasterUnionAllIterator(NewArrayIter(a0), NewArrayIter(a1))
	it.Reset()
	out := make([]int, 0)
	for {
		if v, ok := it.Next(); !ok {
			break
		} else {
			out = append(out, v)
		}
	}
	if !arrays.IsEqual(out, expectedUnionAll) {
		fmt.Println("out")
		fmt.Println(out)
		fmt.Println("correct")
		fmt.Println(expectedUnionAll)
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

func TestFasterUnionAll_NextSomeFor2RandomArrays(t *testing.T) {
	a0 := randArray(3000)
	a1 := randArray(2000)
	//a0 := randArray(8)
	//a1 := randArray(5)
	//a0 := randArray(6)
	//a1 := randArray(6)
	//a0 := []int{1, 3, 7, 10, 11}
	//a1 := []int{2, 4, 6, 11}
	a0Copy := copyArray(a0)
	a1Copy := copyArray(a1)
	expectedUnionAll := copyArray(arrays.UnionAll(a0, a1))

	it := NewFasterUnionAllIterator(NewArrayIter(a0), NewArrayIter(a1))
	it.Reset()
	out := make([]int, 0)
	for {
		if v, ok := it.NextSome(); !ok {
			break
		} else {
			out = append(out, v...)
		}
	}

	if !arrays.IsEqual(out, expectedUnionAll) {
		fmt.Println("out")
		fmt.Println(out)
		fmt.Println("correct")
		fmt.Println(expectedUnionAll)
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

func TestFasterUnionAll_SameArrayFourTimes(t *testing.T) {
	a0 := randArray(300)
	a0Copy := copyArray(a0)
	//a0 := []int{0, 1, 3}
	//a1 := []int{1, 3}arrays.UnionAll(a0, a0)
	res1 := arrays.UnionAll(a0, a0)
	expectedUnionAll := copyArray(arrays.UnionAll(res1, res1))
	itStep1 := NewFasterUnionAllIterator(NewArrayIter(a0), NewArrayIter(a0))
	itStep2 := NewFasterUnionAllIterator(NewArrayIter(a0), NewArrayIter(a0))
	it := NewFasterUnionAllIterator(itStep1, itStep2)
	it.Reset()
	out := make([]int, 0)
	for {
		if v, ok := it.NextSome(); !ok {
			break
		} else {
			out = append(out, v...)
		}
	}
	if !arrays.IsEqual(out, expectedUnionAll) {
		fmt.Println("out")
		fmt.Println(out)
		fmt.Println("correct")
		fmt.Println(expectedUnionAll)
		t.Fail()
	}
	if !arrays.IsEqual(a0, a0Copy) {
		t.Fail()
	}
}

func TestFasterUnionAll_DoubleUnionOfTwoArrays(t *testing.T) {
	a0 := randArray(3000)
	a1 := randArray(2000)
	a0Copy := copyArray(a0)
	a1Copy := copyArray(a1)
	//a0 := []int{0, 1, 3}
	//a1 := []int{1, 3}
	single := arrays.UnionAll(a0, a1)
	expectedUnionAll := copyArray(arrays.UnionAll(single, single))
	it := NewFasterUnionAllIterator(NewFasterUnionAllIterator(NewArrayIter(a0), NewArrayIter(a1)), NewFasterUnionAllIterator(NewArrayIter(a0), NewArrayIter(a1)))
	it.Reset()
	out := make([]int, 0)
	for {
		if v, ok := it.NextSome(); !ok {
			break
		} else {
			out = append(out, v...)
		}
	}

	if !arrays.IsEqual(out, expectedUnionAll) {
		fmt.Println("out")
		fmt.Println(out)
		fmt.Println("correct")
		fmt.Println(expectedUnionAll)
		t.Fail()
	}
	if !arrays.IsEqual(a0, a0Copy) {
		t.Fail()
	}
	if !arrays.IsEqual(a1, a1Copy) {
		t.Fail()
	}
}

func TestFasterUnionAll_SmallThreeArrays(t *testing.T) {
	a0 := randArray(16)
	a1 := randArray(14)
	a2 := randArray(12)
	a0Copy := copyArray(a0)
	a1Copy := copyArray(a1)
	a2Copy := copyArray(a2)
	a01 := copyArray(arrays.UnionAll(a0, a1))
	expectedUnionAll := copyArray(arrays.UnionAll(a01, a2))

	it := NewFasterUnionAllIterator(NewFasterUnionAllIterator(NewArrayIter(a0), NewArrayIter(a1)), NewArrayIter(a2))
	it.Reset()
	out := make([]int, 0)
	for {
		if v, ok := it.NextSome(); !ok {
			break
		} else {
			out = append(out, v...)
		}
	}

	if !arrays.IsEqual(out, expectedUnionAll) {
		fmt.Println("out")
		fmt.Println(out)
		fmt.Println("correct")
		fmt.Println(expectedUnionAll)
		fmt.Println("a0")
		fmt.Println(a0Copy)
		fmt.Println("a1")
		fmt.Println(a1Copy)
		fmt.Println("a2")
		fmt.Println(a2Copy)
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

func TestFasterUnionAll_BigThreeArrays(t *testing.T) {
	a0 := randArray(600000)
	a1 := randArray(400000)
	a2 := randArray(200000)
	a0Copy := copyArray(a0)
	a1Copy := copyArray(a1)
	a2Copy := copyArray(a2)
	a01 := copyArray(arrays.UnionAll(a0, a1))
	expectedUnionAll := copyArray(arrays.UnionAll(a01, a2))

	it := NewFasterUnionAllIterator(NewFasterUnionAllIterator(NewArrayIter(a0), NewArrayIter(a1)), NewArrayIter(a2))
	it.Reset()
	out := make([]int, 0)
	for {
		if v, ok := it.NextSome(); !ok {
			break
		} else {
			out = append(out, v...)
		}
	}

	if !arrays.IsEqual(out, expectedUnionAll) {
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

func BenchmarkFasterUnionAll_ThreeSmallArrays_Next(b *testing.B) {
	a0 := randArray(3000)
	a1 := randArray(2000)
	a2 := randArray(1000)
	it := NewFasterUnionAllIterator(NewFasterUnionAllIterator(NewArrayIter(a0), NewArrayIter(a1)), NewArrayIter(a2))
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

func BenchmarkFasterUnionAll_ThreeSmallArrays_NextSome(b *testing.B) {
	a0 := randArray(3000)
	a1 := randArray(2000)
	a2 := randArray(1000)
	it := NewFasterUnionAllIterator(NewFasterUnionAllIterator(NewArrayIter(a0), NewArrayIter(a1)), NewArrayIter(a2))
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

func BenchmarkFasterUnionAll_ThreeBigArrays_Next(b *testing.B) {
	a0 := randArray(30000)
	a1 := randArray(20000)
	a2 := randArray(10000)
	it := NewFasterUnionAllIterator(NewFasterUnionAllIterator(NewArrayIter(a0), NewArrayIter(a1)), NewArrayIter(a2))
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

func BenchmarkFasterUnionAll_ThreeBigArrays_NextSome(b *testing.B) {
	a0 := randArray(30000)
	a1 := randArray(20000)
	a2 := randArray(10000)
	it := NewFasterUnionAllIterator(NewFasterUnionAllIterator(NewArrayIter(a0), NewArrayIter(a1)), NewArrayIter(a2))
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
