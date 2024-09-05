package iterator

type Iterator[T any] interface {
	Next() (T, error)
}
