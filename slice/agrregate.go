package slice

// Sum 传入一个函数定义两个类型如何相加
func Sum[T any](src []T, f func(lhs, rhs T) T) (T, error) {
	if err := checkSlice(src); err != nil {
		var t T
		return t, err
	}
	sum := src[0]
	for i := 1; i < len(src); i++ {
		sum = f(sum, src[i])
	}
	return sum, nil
}

// Max 传入一个函数定义两个类型如何比较 返回true说明 lhs 值大
func Max[T any](src []T, f func(lhs, rhs T) bool) (T, error) {
	if err := checkSlice(src); err != nil {
		var t T
		return t, err
	}
	max := src[0]
	for i := 1; i < len(src); i++ {
		if !f(max, src[i]) {
			max = src[i]
		}
	}
	return max, nil
}

// Min 传入一个函数定义两个类型如何比较 返回true说明 lhs 值小
func Min[T any](src []T, f func(lhs, rhs T) bool) (T, error) {
	if err := checkSlice(src); err != nil {
		var t T
		return t, err
	}
	min := src[0]
	for i := 1; i < len(src); i++ {
		if !f(min, src[i]) {
			min = src[i]
		}
	}
	return min, nil
}

// Count 传入一个函数用于判断一个元素是否是要统计的元素
func Count[T any](src []T, f func(val T) bool) (int, error) {
	if err := checkSlice(src); err != nil {
		return 0, err
	}
	count := 0
	for i := range src {
		if f(src[i]) {
			count++
		}
	}
	return count, nil
}

// Contains 传入一个函数用于判断一个元素是否存在
func Contains[T any](src []T, f func(val T) bool) (bool, error) {
	if err := checkSlice(src); err != nil {
		return false, err
	}
	for i := range src {
		if f(src[i]) {
			return true, nil
		}
	}
	return false, nil
}
