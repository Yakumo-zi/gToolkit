package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReduce(t *testing.T) {
	testCases := []struct {
		name   string
		src    []int
		reduce func(init int, idx int, src []int) int
		want   int
	}{
		{
			name: "sum of slice",
			src:  []int{1, 2, 3, 4, 5},
			reduce: func(init int, idx int, src []int) int {
				return init + src[idx]
			},
			want: 15,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Reduce(0, tc.src, tc.reduce)
			if err != nil {

			}
			assert.Equal(t, tc.want, res)
		})
	}
}
