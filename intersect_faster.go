package iterator

const FasterIntersectionIteratorSliceSize = 256

type FasterIntersectionIterator struct {
	a         Iterator
	b         Iterator
	aPos      int
	bPos      int
	aSlice    []int
	bSlice    []int
	needMoveA bool
	needMoveB bool

	nextSlice []int

	someBuf [FasterIntersectionIteratorSliceSize]int
}

func NewFasterIntersectionIterator(a, b Iterator) *FasterIntersectionIterator {
	return &FasterIntersectionIterator{a: a, b: b, aPos: 0, bPos: 0, aSlice: []int{}, bSlice: []int{}, needMoveA: true, needMoveB: true, nextSlice: []int{}}
}

func (it *FasterIntersectionIterator) Reset() {
	it.a.Reset()
	it.b.Reset()
	it.aPos = 0
	it.bPos = 0
	it.aSlice = []int{}
	it.bSlice = []int{}
	it.nextSlice = []int{}
	it.needMoveA, it.needMoveB = true, true
}

func (it *FasterIntersectionIterator) Next() (a int, ok bool) {
	if len(it.nextSlice) == 0 {
		it.nextSlice, ok = it.NextSome()
		if !ok {
			return 0, false
		}
	}

	a = it.nextSlice[0]
	it.nextSlice = it.nextSlice[1:]
	return a, true
}

func (it *FasterIntersectionIterator) NextSome() (some []int, ok bool) {
	count := 0
	a, b := 0, 0
	for {
		if it.needMoveA {
			if it.aPos == len(it.aSlice) {
				it.aPos = 0
				it.aSlice, ok = it.a.NextSome()
				if !ok {
					break
				}
			}
			a = it.aSlice[it.aPos]
			it.aPos++
		}

		if it.needMoveB {
			if it.bPos == len(it.bSlice) {
				it.bPos = 0
				it.bSlice, ok = it.b.NextSome()
				if !ok {
					break
				}
			}
			b = it.bSlice[it.bPos]
			it.bPos++
		}

		if a < b {
			it.needMoveA, it.needMoveB = true, false
		} else if a > b {
			it.needMoveA, it.needMoveB = false, true
		} else { // a == b
			it.someBuf[count] = a
			count++
			if count == FasterIntersectionIteratorSliceSize {
				break
			}
			it.needMoveA, it.needMoveB = true, true
		}
	}

	return it.someBuf[:count], count > 0
}
