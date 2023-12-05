package list

type Node[T any] struct {
	val  T
	next *Node[T]
	prev *Node[T]
}

type LinkedList[T any] struct {
	head     *Node[T]
	tail     *Node[T]
	length   int
	capacity int
}

func (l *LinkedList[T]) Get(index int) (T, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Append(ts T) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Add(index int, t T) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Set(index int, t T) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Delete(index int) (T, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Len() int {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Cap() int {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Range(fn func(index int, t T) error) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) AsSlice() []T {
	//TODO implement me
	panic("implement me")
}
