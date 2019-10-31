package iterator

type Iterator interface {
	Reset()
	Next() (int, bool)
	NextSome() ([]int, bool)
}
