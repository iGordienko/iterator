package iterator

const FasterUnionAllIteratorSliceSize = 256

type FasterUnionAllIterator struct {
	a      Iterator
	b      Iterator
	aPos   int
	bPos   int
	aSlice []int
	bSlice []int
	hasA   bool
	hasB   bool
	curA   int
	curB   int

	nextSlice []int

	someBuf [FasterUnionAllIteratorSliceSize]int
}

func NewFasterUnionAllIterator(a, b Iterator) *FasterUnionAllIterator {
	return &FasterUnionAllIterator{a: a, b: b, aPos: 0, bPos: 0, aSlice: []int{}, bSlice: []int{}, hasA: true, hasB: true, nextSlice: []int{}}
}

func (it *FasterUnionAllIterator) Reset() {
	it.a.Reset()
	it.b.Reset()
	it.aPos = 0
	it.bPos = 0
	it.nextSlice = []int{}

	ok := true
	it.aSlice, ok = it.a.NextSome()
	if ok {
		it.curA = it.aSlice[0]
		it.aPos = 1
		it.hasA = true
	} else {
		it.aPos = 0
		it.hasA = false
	}

	it.bSlice, ok = it.b.NextSome()
	if ok {
		it.curB = it.bSlice[0]
		it.bPos = 1
		it.hasB = true
	} else {
		it.bPos = 0
		it.hasB = false
	}
}

func (it *FasterUnionAllIterator) Next() (a int, ok bool) {
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

func (it *FasterUnionAllIterator) NextSome() (some []int, ok bool) {
	count := 0
	needMoveA, needMoveB := true, true
	for {
		if it.hasB {
			if it.hasA && it.curA < it.curB {
				it.someBuf[count] = it.curA
				needMoveA, needMoveB = true, false
			} else {
				it.someBuf[count] = it.curB
				needMoveA, needMoveB = false, true
			}
		} else if it.hasA {
			it.someBuf[count] = it.curA
			needMoveA, needMoveB = true, false
		} else {
			break
		}
		count++

		if needMoveA {
			if it.aPos == len(it.aSlice) {
				it.aSlice, ok = it.a.NextSome()
				if ok {
					it.curA = it.aSlice[0]
					it.aPos = 1
				} else {
					it.aPos = 0
					it.hasA = false
				}
			} else {
				it.curA = it.aSlice[it.aPos]
				it.aPos++
			}
		}

		if needMoveB {
			if it.bPos == len(it.bSlice) {
				it.bSlice, ok = it.b.NextSome()
				if ok {
					it.curB = it.bSlice[0]
					it.bPos = 1
				} else {
					it.bPos = 0
					it.hasB = false
				}
			} else {
				it.curB = it.bSlice[it.bPos]
				it.bPos++
			}
		}

		if count == FasterUnionAllIteratorSliceSize {
			break
		}
	}

	return it.someBuf[:count], count > 0
}
