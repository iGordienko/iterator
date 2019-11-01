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

var (
	aa = randArray(19164)
	ab = randArray(2180)
	ac = randArray(83148)
	ad = randArray(146112)
	ae = randArray(84204)
	af = randArray(143159)
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

func Benchmark_Complex_JustIterate_BasicIterators(b *testing.B) {
	// (a || b) && (c || d) && (e || f)
	u1 := NewUnionIter(NewArrayIter(aa), NewArrayIter(ab))
	u2 := NewUnionIter(NewArrayIter(ac), NewArrayIter(ad))
	u3 := NewUnionIter(NewArrayIter(ae), NewArrayIter(af))

	i1 := NewInterIter(u1, u2)
	it := NewInterIter(i1, u3)

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

func Benchmark_Complex_Exploded_Basicterators(b *testing.B) {
	// (a || b) && (c || d) && (e || f)

	//(a && c && e)
	i11 := NewInterIter(NewArrayIter(aa), NewArrayIter(ac))
	i1 := NewInterIter(i11, NewArrayIter(ae))

	//|| (a && d && e)
	i22 := NewInterIter(NewArrayIter(aa), NewArrayIter(ad))
	i2 := NewInterIter(i22, NewArrayIter(ae))

	//|| (b && c && e)
	i33 := NewInterIter(NewArrayIter(ab), NewArrayIter(ac))
	i3 := NewInterIter(i33, NewArrayIter(ae))

	//|| (b && d && e)
	i44 := NewInterIter(NewArrayIter(ab), NewArrayIter(ad))
	i4 := NewInterIter(i44, NewArrayIter(ae))

	//|| (a && c && f)
	i55 := NewInterIter(NewArrayIter(ab), NewArrayIter(ac))
	i5 := NewInterIter(i55, NewArrayIter(af))

	//|| (a && d && f)
	i66 := NewInterIter(NewArrayIter(aa), NewArrayIter(ad))
	i6 := NewInterIter(i66, NewArrayIter(af))

	//|| (b && c && f)
	i77 := NewInterIter(NewArrayIter(ab), NewArrayIter(ac))
	i7 := NewInterIter(i77, NewArrayIter(af))

	//|| (b && d && f)
	i88 := NewInterIter(NewArrayIter(ab), NewArrayIter(ad))
	i8 := NewInterIter(i88, NewArrayIter(af))

	u1 := NewUnionIter(i1, i2)
	u2 := NewUnionIter(u1, i3)
	u3 := NewUnionIter(u2, i4)
	u4 := NewUnionIter(u3, i5)
	u5 := NewUnionIter(u4, i6)
	u6 := NewUnionIter(u5, i7)

	it := NewUnionIter(u6, i8)

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

func Benchmark_Complex_JustIterate_FasterIterators(b *testing.B) {
	// (a || b) && (c || d) && (e || f)
	u1 := NewFasterUnionIterator(NewArrayIter(aa), NewArrayIter(ab))
	u2 := NewFasterUnionIterator(NewArrayIter(ac), NewArrayIter(ad))
	u3 := NewFasterUnionIterator(NewArrayIter(ae), NewArrayIter(af))

	i1 := NewFasterIntersectionIterator(u1, u2)
	it := NewFasterIntersectionIterator(i1, u3)

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

func Benchmark_Complex_Exploded_FasterIterators(b *testing.B) {
	// (a || b) && (c || d) && (e || f)

	//(a && c && e)
	i11 := NewFasterIntersectionIterator(NewArrayIter(aa), NewArrayIter(ac))
	i1 := NewFasterIntersectionIterator(i11, NewArrayIter(ae))

	//|| (a && d && e)
	i22 := NewFasterIntersectionIterator(NewArrayIter(aa), NewArrayIter(ad))
	i2 := NewFasterIntersectionIterator(i22, NewArrayIter(ae))

	//|| (b && c && e)
	i33 := NewFasterIntersectionIterator(NewArrayIter(ab), NewArrayIter(ac))
	i3 := NewFasterIntersectionIterator(i33, NewArrayIter(ae))

	//|| (b && d && e)
	i44 := NewFasterIntersectionIterator(NewArrayIter(ab), NewArrayIter(ad))
	i4 := NewFasterIntersectionIterator(i44, NewArrayIter(ae))

	//|| (a && c && f)
	i55 := NewFasterIntersectionIterator(NewArrayIter(ab), NewArrayIter(ac))
	i5 := NewFasterIntersectionIterator(i55, NewArrayIter(af))

	//|| (a && d && f)
	i66 := NewFasterIntersectionIterator(NewArrayIter(aa), NewArrayIter(ad))
	i6 := NewFasterIntersectionIterator(i66, NewArrayIter(af))

	//|| (b && c && f)
	i77 := NewFasterIntersectionIterator(NewArrayIter(ab), NewArrayIter(ac))
	i7 := NewFasterIntersectionIterator(i77, NewArrayIter(af))

	//|| (b && d && f)
	i88 := NewFasterIntersectionIterator(NewArrayIter(ab), NewArrayIter(ad))
	i8 := NewFasterIntersectionIterator(i88, NewArrayIter(af))

	u1 := NewFasterUnionIterator(i1, i2)
	u2 := NewFasterUnionIterator(u1, i3)
	u3 := NewFasterUnionIterator(u2, i4)
	u4 := NewFasterUnionIterator(u3, i5)
	u5 := NewFasterUnionIterator(u4, i6)
	u6 := NewFasterUnionIterator(u5, i7)

	it := NewFasterUnionIterator(u6, i8)

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
