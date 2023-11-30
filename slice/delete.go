package slice

import (
	"errors"
	"fmt"
)

// checkSlice 检查切片的合法性，是否是 nil ，以及长度是否为 0
func checkSlice[T any](src []T) error {
	if src == nil {
		return errors.New("argument src is nil")
	}
	if len(src) == 0 {
		return errors.New("slice length is zero")
	}
	return nil
}

// DeleteIndex 通过索引删除切片中的元素，当索引不合法或者切片不合法时会抛出错误
func DeleteIndex[T any](src []T, idx uint) ([]T, error) {
	if err := checkSlice(src); err != nil {
		return src, err
	}
	if len(src) <= int(idx) {
		return src, errors.New(fmt.Sprintf("idx %d is out of length %d \n", idx, len(src)))
	}
	res := append(src[:idx], src[idx+1:]...)
	return res, nil
}

// DeleteValue 通过值删除切片中的元素，但是该值必须是可比较的，当切片不合法或者切片中没有指定元素时会抛出错误
func DeleteValue[T comparable](src []T, value T) ([]T, error) {
	if err := checkSlice(src); err != nil {
		return src, err
	}
	for i, v := range src {
		if v == value {
			return DeleteIndex(src, uint(i))
		}
	}
	return src, errors.New(fmt.Sprintf("slice is not contains element %+v", value))
}

// DeleteWithFunc 通过传入一个过滤函数来删除且片中的元素，删除满足过滤函数的第一个元素
// filter 是用于过滤元素的函数，该函数会传入两个值，一个是当前元素的索引，一个是当前元素的值，该函数需要返回一个 bool 值来判断
// 当没有删除元素时，不会报错
func DeleteWithFunc[T comparable](src []T, filter func(i int, src []T) bool) ([]T, error) {
	if err := checkSlice(src); err != nil {
		return src, err
	}
	for i, _ := range src {
		if filter(i, src) {
			return DeleteIndex(src, uint(i))
		}
	}
	return src, nil
}
