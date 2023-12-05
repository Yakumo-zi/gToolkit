package list

import "fmt"

type ArrayList[T any] struct {
	content []T
}

func (a *ArrayList[T]) Get(index int) (T, error) {
	if index >= len(a.content) || index < 0 {
		var t T
		return t, fmt.Errorf("index %d out of bound", index)
	}
	return a.content[index], nil
}

func (a *ArrayList[T]) Append(ts ...T) error {
	a.content = append(a.content, ts...)
	return nil
}

func (a *ArrayList[T]) Add(index int, t T) error {
	if index > len(a.content) || index < 0 {
		return fmt.Errorf("index %d out of bound", index)
	}
	if index == len(a.content) {
		return a.Append(t)
	}
	left := a.content[:index]
	right := a.content[index:]
	a.content = append(left, t)
	a.content = append(a.content, right...)
	return nil
}

func (a *ArrayList[T]) Set(index int, t T) error {
	if index >= len(a.content) || index < 0 {
		return fmt.Errorf("index %d out of bound", index)
	}
	a.content[index] = t
	return nil
}

func (a *ArrayList[T]) Delete(index int) (T, error) {
	if index >= len(a.content) || index < 0 {
		var t T
		return t, fmt.Errorf("index %d out of bound", index)
	}
	item := a.content[index]
	a.content = append(a.content[:index], a.content[index+1:]...)
	return item, nil
}

func (a *ArrayList[T]) Len() int {
	return len(a.content)
}

func (a *ArrayList[T]) Cap() int {
	return cap(a.content)
}

func (a *ArrayList[T]) Range(fn func(index int, t T) error) error {
	for i, v := range a.content {
		if err := fn(i, v); err != nil {
			return err
		}
	}
	return nil
}

func (a *ArrayList[T]) AsSlice() []T {
	res := make([]T, 0, len(a.content))
	return append(res, a.content...)
}
