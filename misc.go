package iterator

import (
	"math/rand"
	"sort"
	"time"
)

func init() {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
}

func copyArray(inArr []int) []int {
	tmp := make([]int, len(inArr))
	copy(tmp, inArr)
	return tmp
}

func randArray(n int) []int {
	if n == 0 {
		return []int{}
	}
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = rand.Intn(n * 2)
	}
	sort.Ints(r)

	unique := []int{r[0]}
	last := r[0]
	for _, v := range r[1:] {
		if v != last {
			unique = append(unique, v)
			last = v
		}
	}

	return unique
}

func isEqual(a, b []int) bool {
	n := len(a)
	if n != len(b) {
		return false
	}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
