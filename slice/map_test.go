package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	testCases := []struct {
		name    string
		src     []int
		convert func(idx int, value int) string
		want    []string
	}{
		{
			name: "convert to string",
			src:  []int{1, 2, 3, 4, 5},
			convert: func(idx int, value int) string {
				return fmt.Sprintf("%d", value)
			},
			want: []string{"1", "2", "3", "4", "5"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Map(tc.src, tc.convert)
			if err != nil {
				return
			}
			assert.Equal(t, tc.want, res)
		})
	}
}
