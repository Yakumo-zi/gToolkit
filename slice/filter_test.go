package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	testCases := []struct {
		name   string
		src    []int
		filter func(idx int, value int) bool
		want   []int
	}{
		{
			name: "omit element 3",
			src:  []int{1, 2, 3, 4, 5},
			filter: func(idx int, value int) bool {
				return value != 3
			},
			want: []int{1, 2, 4, 5},
		},
		{
			name: "elements divisible by 2 are omitted ",
			src:  []int{1, 2, 3, 4, 5},
			filter: func(idx int, value int) bool {
				return value%2 != 0
			},
			want: []int{1, 3, 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Filter(tc.src, tc.filter)
			if err != nil {
				return
			}
			assert.Equal(t, tc.want, res)
		})
	}
}
