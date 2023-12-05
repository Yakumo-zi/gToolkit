package iterator

type Iterator[T any] interface {
	Next() (T, error)
	Prev() (T, error)
	End() bool
}
