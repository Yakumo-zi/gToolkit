package slice

func Map[T any, D any](src []T, m func(idx int, value T) D) ([]D, error) {
	if err := checkSlice(src); err != nil {
		return nil, err
	}
	dest := make([]D, len(src))
	for i, v := range src {
		dest[i] = m(i, v)
	}
	return dest, nil
}
