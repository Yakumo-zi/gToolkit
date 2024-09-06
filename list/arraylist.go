package list

import "fmt"

var _ List[any] = (*ArrayList[any])(nil)

type ArrayList[T any] struct {
	data []T
}

func (a *ArrayList[T]) checkIndex(index int) error {
	if index < 0 || index >= len(a.data) {
		return fmt.Errorf("index out of range")
	}
	return nil
}

// Add implements List.
func (a *ArrayList[T]) Add(index int, t T) error {
	if err := a.checkIndex(index); err != nil {
		return err
	}
	a.data = append(append(a.data[:index], t), a.data[index:]...)
	return nil
}

// Append implements List.
func (a *ArrayList[T]) Append(ts ...T) error {
	a.data = append(a.data, ts...)
	return nil
}

// AsSlice implements List.
func (a *ArrayList[T]) AsSlice() []T {
	return a.data
}

// Cap implements List.
func (a *ArrayList[T]) Cap() int {
	return cap(a.data)
}

// Delete implements List.
func (a *ArrayList[T]) Delete(index int) (res T, err error) {
	if err = a.checkIndex(index); err != nil {
		return
	}
	res = a.data[index]
	a.data = append(a.data[:index], a.data[index+1:]...)
	return
}

// Get implements List.
func (a *ArrayList[T]) Get(index int) (res T, err error) {
	if err = a.checkIndex(index); err != nil {
		return
	}
	res = a.data[index]
	return
}

// Len implements List.
func (a *ArrayList[T]) Len() int {
	return len(a.data)
}

// Range implements List.
func (a *ArrayList[T]) Range(fn func(index int, t T) error) (err error) {
	for i, t := range a.data {
		if err = fn(i, t); err != nil {
			return
		}
	}
	return
}

// Set implements List.
func (a *ArrayList[T]) Set(index int, t T) error {
	if err := a.checkIndex(index); err != nil {
		return err
	}
	a.data[index] = t
	return nil
}
