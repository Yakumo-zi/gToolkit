package slice

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteIndex(t *testing.T) {
	testCases := []struct {
		name        string
		src         []int
		deleteIndex uint
		want        []int
		wantErr     error
	}{
		{
			name:        "pass a nil to function",
			src:         nil,
			deleteIndex: 0,
			want:        nil,
			wantErr:     errors.New("argument src is nil"),
		},
		{
			name:        "slice length is zero",
			src:         []int{},
			deleteIndex: 0,
			want:        []int{},
			wantErr:     errors.New("slice length is zero"),
		},
		{
			name:        "delete success",
			src:         []int{1, 2, 3, 4, 5},
			deleteIndex: 0,
			want:        []int{2, 3, 4, 5},
			wantErr:     nil,
		},
		{
			name:        "out of length",
			src:         []int{1, 2, 3, 4, 5},
			deleteIndex: 10,
			want:        []int{1, 2, 3, 4, 5},
			wantErr:     fmt.Errorf("idx %d is out of length %d \n", 10, 5),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			want, err := DeleteIndex(tc.src, tc.deleteIndex)
			if err != nil {
				assert.Equal(t, err, tc.wantErr)
			}
			assert.Equal(t, tc.want, want)
		})
	}
}

func TestDeleteValue(t *testing.T) {
	testCases := []struct {
		name        string
		src         []int
		deleteValue int
		want        []int
		wantErr     error
	}{
		{
			name:        "pass a nil to function",
			src:         nil,
			deleteValue: 0,
			want:        nil,
			wantErr:     errors.New("argument src is nil"),
		},
		{
			name:        "slice length is zero",
			src:         []int{},
			deleteValue: 0,
			want:        []int{},
			wantErr:     errors.New("slice length is zero"),
		},
		{
			name:        "delete success",
			src:         []int{1, 2, 3, 4, 5},
			deleteValue: 1,
			want:        []int{2, 3, 4, 5},
			wantErr:     nil,
		},
		{
			name:        "slice is not contains the element",
			src:         []int{1, 2, 3, 4, 5},
			deleteValue: 10,
			want:        []int{1, 2, 3, 4, 5},
			wantErr:     fmt.Errorf("slice is not contains element %+v", 10),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			want, err := DeleteValue(tc.src, tc.deleteValue)
			if err != nil {
				assert.Equal(t, err, tc.wantErr)
			}
			assert.Equal(t, tc.want, want)
		})
	}
}

func TestDeleteWithFunc(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		filter  func(idx int, src []int) bool
		want    []int
		wantErr error
	}{
		{
			name:    "pass a nil to function",
			src:     nil,
			filter:  nil,
			want:    nil,
			wantErr: errors.New("argument src is nil"),
		},
		{
			name:    "slice length is zero",
			src:     []int{},
			filter:  nil,
			want:    []int{},
			wantErr: errors.New("slice length is zero"),
		},
		{
			name: "delete value success",
			src:  []int{1, 2, 3, 4, 5},
			filter: func(idx int, src []int) bool {
				return src[idx] == 1
			},
			want:    []int{2, 3, 4, 5},
			wantErr: nil,
		},
		{
			name: "delete index success",
			src:  []int{1, 2, 3, 4, 5},
			filter: func(idx int, src []int) bool {
				return idx == 2
			},
			want:    []int{1, 2, 4, 5},
			wantErr: nil,
		},
		{
			name: "slice is not contains the element",
			src:  []int{1, 2, 3, 4, 5},
			filter: func(idx int, src []int) bool {
				return src[idx] == 10
			},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			want, err := DeleteWithFunc(tc.src, tc.filter)
			if err != nil {
				assert.Equal(t, err, tc.wantErr)
			}
			assert.Equal(t, tc.want, want)
		})
	}
}
