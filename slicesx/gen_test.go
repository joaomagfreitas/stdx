package slicesx_test

import (
	"slices"
	"testing"

	"github.com/joaomagfreitas/stdx/slicesx"
)

func TestGen(t *testing.T) {
	testCases := []struct {
		desc   string
		count  int
		fn     func(int) int
		output []int
	}{
		{
			desc:   "returns empty if count is 0",
			count:  0,
			fn:     func(i int) int { return i * 10 },
			output: nil,
		},
		{
			desc:   "applies generate function to each index",
			count:  5,
			fn:     func(i int) int { return i * 10 },
			output: []int{0, 10, 20, 30, 40},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if o := slicesx.Gen(tC.count, tC.fn); !slices.Equal(o, tC.output) {
				t.Fail()
			}
		})
	}
}
