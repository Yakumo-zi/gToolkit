package slice

func Filter[T any](src []T, m func(idx int, value T) bool) ([]T, error) {
	if err := checkSlice(src); err != nil {
		return src, err
	}
	dest := make([]T, 0, len(src))
	for i, v := range src {
		if m(i, v) {
			dest = append(dest, v)
		}
	}
	return dest, nil
}
