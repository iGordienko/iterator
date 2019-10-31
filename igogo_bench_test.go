package iterator

import (
	"testing"

	//"github.com/VictoriaMetrics/VictoriaMetrics/lib/uint64set"
	"github.com/freepk/arrays"
	//"github.com/freepk/iterator"
)

var (
	resultSetSize = 10000
	firstSetSize  = resultSetSize
	secondSetSize = resultSetSize / 4 * 3
	thirdSetSize  = resultSetSize / 4 * 2
	fourthSetSize = resultSetSize / 4 * 2
)

var (
	resultSet0 = make([]int, resultSetSize)
	resultSet1 = make([]int, resultSetSize)
	firstSet   = randArray(firstSetSize)
	secondSet  = randArray(secondSetSize)
	thirdSet   = randArray(thirdSetSize)
	fourthSet  = randArray(fourthSetSize)
)

func Benchmark_ArrayInter(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arrays.IntersectEx(resultSet0[:0], firstSet, secondSet)
	}
}

func Benchmark_BasicInter(b *testing.B) {
	it := NewInterIter(NewArrayIter(firstSet), NewArrayIter(secondSet))
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

func Benchmark_BasicInter_NextSome(b *testing.B) {
	it := NewInterIter(NewArrayIter(firstSet), NewArrayIter(secondSet))
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

func Benchmark_FasterIntersector(b *testing.B) {
	it := NewFasterIntersectionIterator(NewArrayIter(firstSet), NewArrayIter(secondSet))
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

func Benchmark_FasterIntersector_NextSome(b *testing.B) {
	it := NewFasterIntersectionIterator(NewArrayIter(firstSet), NewArrayIter(secondSet))
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

func Benchmark_BasicInter_FourArrays(b *testing.B) {
	it := NewInterIter(NewInterIter(NewInterIter(NewArrayIter(firstSet), NewArrayIter(secondSet)), NewArrayIter(thirdSet)), NewArrayIter(fourthSet))
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


func Benchmark_FasterIntersector_FourArrays_NextSome(b *testing.B) {
	it := NewFasterIntersectionIterator(NewFasterIntersectionIterator(NewFasterIntersectionIterator(NewArrayIter(firstSet), NewArrayIter(secondSet)), NewArrayIter(thirdSet)), NewArrayIter(fourthSet))
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

func Benchmark_BasicInter_FourArrays2(b *testing.B) {
	it := NewInterIter((NewInterIter(NewArrayIter(firstSet), NewArrayIter(secondSet))), NewInterIter(NewArrayIter(thirdSet), NewArrayIter(fourthSet)))
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

func Benchmark_FasterIntersector_FourArrays2_NextSome(b *testing.B) {
	it := NewFasterIntersectionIterator((NewFasterIntersectionIterator(NewArrayIter(firstSet), NewArrayIter(secondSet))), NewFasterIntersectionIterator(NewArrayIter(thirdSet), NewArrayIter(fourthSet)))
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

func Benchmark_BasicInter_FourSameArrays(b *testing.B) {
	it := NewInterIter((NewInterIter(NewArrayIter(firstSet), NewArrayIter(firstSet))), NewInterIter(NewArrayIter(firstSet), NewArrayIter(firstSet)))
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

func Benchmark_FasterIntersector_FourSameArrays_NextSome(b *testing.B) {
	it := NewFasterIntersectionIterator((NewFasterIntersectionIterator(NewArrayIter(firstSet), NewArrayIter(firstSet))), NewFasterIntersectionIterator(NewArrayIter(firstSet), NewArrayIter(firstSet)))
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

func Benchmark_SmallComplex_BasicIterators(b *testing.B) {
	a0 := randArray(300)
	a1 := randArray(200)
	a2 := randArray(100)
	a3 := randArray(100)
	a4 := randArray(1000)
	a5 := randArray(500)
	a6 := randArray(20)
	a7 := randArray(200)
	a8 := randArray(600)

	step1 := NewUnionIter(NewArrayIter(a0), NewArrayIter(a1))
	step2 := NewUnionIter(NewArrayIter(a2), NewArrayIter(a3))
	step3 := NewUnionIter(NewArrayIter(a4), NewArrayIter(a5))
	step4 := NewUnionIter(NewArrayIter(a6), NewArrayIter(a7))

	step5 := NewInterIter(step1, step2)
	step6 := NewInterIter(step3, step4)

	step7 := NewInterIter(step5, step6)
	it := NewInterIter(step7, NewArrayIter(a8))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			_, ok := it.Next()
			if !ok {
				break
			}
		}
	}
}

func Benchmark_SmallComplex_FasterIterators(b *testing.B) {
	a0 := randArray(300)
	a1 := randArray(200)
	a2 := randArray(100)
	a3 := randArray(100)
	a4 := randArray(1000)
	a5 := randArray(500)
	a6 := randArray(20)
	a7 := randArray(200)
	a8 := randArray(600)

	step1 := NewFasterUnionIterator(NewArrayIter(a0), NewArrayIter(a1))
	step2 := NewFasterUnionIterator(NewArrayIter(a2), NewArrayIter(a3))
	step3 := NewFasterUnionIterator(NewArrayIter(a4), NewArrayIter(a5))
	step4 := NewFasterUnionIterator(NewArrayIter(a6), NewArrayIter(a7))

	step5 := NewFasterIntersectionIterator(step1, step2)
	step6 := NewFasterIntersectionIterator(step3, step4)

	step7 := NewFasterIntersectionIterator(step5, step6)
	it := NewFasterIntersectionIterator(step7, NewArrayIter(a8))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			_, ok := it.Next()
			if !ok {
				break
			}
		}
	}
}

func Benchmark_Complex_BasicIterators(b *testing.B) {
	a0 := randArray(30000)
	a1 := randArray(20000)
	a2 := randArray(10000)
	a3 := randArray(100)
	a4 := randArray(100000)
	a5 := randArray(50000)
	a6 := randArray(2000)
	a7 := randArray(20000)
	a8 := randArray(60000)

	step1 := NewUnionIter(NewArrayIter(a0), NewArrayIter(a1))
	step2 := NewUnionIter(NewArrayIter(a2), NewArrayIter(a3))
	step3 := NewUnionIter(NewArrayIter(a4), NewArrayIter(a5))
	step4 := NewUnionIter(NewArrayIter(a6), NewArrayIter(a7))

	step5 := NewInterIter(step1, step2)
	step6 := NewInterIter(step3, step4)

	step7 := NewInterIter(step5, step6)
	it := NewInterIter(step7, NewArrayIter(a8))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			_, ok := it.Next()
			if !ok {
				break
			}
		}
	}
}

func Benchmark_Complex_FasterIterators(b *testing.B) {
	a0 := randArray(30000)
	a1 := randArray(20000)
	a2 := randArray(10000)
	a3 := randArray(100)
	a4 := randArray(100000)
	a5 := randArray(50000)
	a6 := randArray(2000)
	a7 := randArray(20000)
	a8 := randArray(60000)

	step1 := NewFasterUnionIterator(NewArrayIter(a0), NewArrayIter(a1))
	step2 := NewFasterUnionIterator(NewArrayIter(a2), NewArrayIter(a3))
	step3 := NewFasterUnionIterator(NewArrayIter(a4), NewArrayIter(a5))
	step4 := NewFasterUnionIterator(NewArrayIter(a6), NewArrayIter(a7))

	step5 := NewFasterIntersectionIterator(step1, step2)
	step6 := NewFasterIntersectionIterator(step3, step4)

	step7 := NewFasterIntersectionIterator(step5, step6)
	it := NewFasterIntersectionIterator(step7, NewArrayIter(a8))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			_, ok := it.Next()
			if !ok {
				break
			}
		}
	}
}

func Benchmark_HugeComplex_BasicIterators(b *testing.B) {
	a0 := randArray(300000)
	a1 := randArray(100000)
	a2 := randArray(100000)
	a3 := randArray(1000)
	a4 := randArray(1000000)
	a5 := randArray(500000)
	a6 := randArray(20000)
	a7 := randArray(200000)
	a8 := randArray(600000)

	step1 := NewFasterUnionIterator(NewArrayIter(a0), NewArrayIter(a1))
	step2 := NewFasterUnionIterator(NewArrayIter(a2), NewArrayIter(a3))
	step3 := NewFasterUnionIterator(NewArrayIter(a4), NewArrayIter(a5))
	step4 := NewFasterUnionIterator(NewArrayIter(a6), NewArrayIter(a7))

	step5 := NewFasterIntersectionIterator(step1, step2)
	step6 := NewFasterIntersectionIterator(step3, step4)

	step7 := NewFasterIntersectionIterator(step5, step6)

	step8 := NewFasterIntersectionIterator(step7, NewArrayIter(a6))
	step9 := NewFasterIntersectionIterator(step8, NewArrayIter(a1))

	it := NewFasterIntersectionIterator(step9, NewArrayIter(a8))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			_, ok := it.Next()
			if !ok {
				break
			}
		}
	}
}

func Benchmark_HugeComplex_FasterIterators(b *testing.B) {
	a0 := randArray(300000)
	a1 := randArray(100000)
	a2 := randArray(100000)
	a3 := randArray(1000)
	a4 := randArray(1000000)
	a5 := randArray(500000)
	a6 := randArray(20000)
	a7 := randArray(200000)
	a8 := randArray(600000)

	step1 := NewFasterUnionIterator(NewArrayIter(a0), NewArrayIter(a1))
	step2 := NewFasterUnionIterator(NewArrayIter(a2), NewArrayIter(a3))
	step3 := NewFasterUnionIterator(NewArrayIter(a4), NewArrayIter(a5))
	step4 := NewFasterUnionIterator(NewArrayIter(a6), NewArrayIter(a7))

	step5 := NewFasterIntersectionIterator(step1, step2)
	step6 := NewFasterIntersectionIterator(step3, step4)

	step7 := NewFasterIntersectionIterator(step5, step6)

	step8 := NewFasterIntersectionIterator(step7, NewArrayIter(a6))
	step9 := NewFasterIntersectionIterator(step8, NewArrayIter(a1))

	it := NewFasterIntersectionIterator(step9, NewArrayIter(a8))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			_, ok := it.Next()
			if !ok {
				break
			}
		}
	}
}

func Benchmark_HugeComplex_FasterIterators_NextSome(b *testing.B) {
	a0 := randArray(300000)
	a1 := randArray(100000)
	a2 := randArray(100000)
	a3 := randArray(1000)
	a4 := randArray(1000000)
	a5 := randArray(500000)
	a6 := randArray(20000)
	a7 := randArray(200000)
	a8 := randArray(600000)

	step1 := NewFasterUnionIterator(NewArrayIter(a0), NewArrayIter(a1))
	step2 := NewFasterUnionIterator(NewArrayIter(a2), NewArrayIter(a3))
	step3 := NewFasterUnionIterator(NewArrayIter(a4), NewArrayIter(a5))
	step4 := NewFasterUnionIterator(NewArrayIter(a6), NewArrayIter(a7))

	step5 := NewFasterIntersectionIterator(step1, step2)
	step6 := NewFasterIntersectionIterator(step3, step4)

	step7 := NewFasterIntersectionIterator(step5, step6)

	step8 := NewFasterIntersectionIterator(step7, NewArrayIter(a6))
	step9 := NewFasterIntersectionIterator(step8, NewArrayIter(a1))

	it := NewFasterIntersectionIterator(step9, NewArrayIter(a8))

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
