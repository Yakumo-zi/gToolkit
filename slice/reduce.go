package slice

func Reduce[T any, D any](init D, src []T, m func(init D, idx int, src []T) D) (D, error) {
	if err := checkSlice(src); err != nil {
		return init, nil
	}
	for i, _ := range src {
		init = m(init, i, src)
	}
	return init, nil
}
