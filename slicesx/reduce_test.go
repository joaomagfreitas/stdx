package slicesx_test

import (
	"testing"

	"github.com/joaomagfreitas/stdx/slicesx"
)

func TestReduce(t *testing.T) {
	testCases := []struct {
		combine func(previous int, next int) int
		desc    string
		input   []int
		output  int
	}{
		{
			desc:   "returns empty value if nil slice is passed",
			input:  nil,
			output: 0,
			combine: func(previous, next int) int {
				return previous + next
			},
		},
		{
			desc:   "returns empty value if empty slice is passed",
			input:  []int{},
			output: 0,
			combine: func(previous, next int) int {
				return previous + next
			},
		},
		{
			desc:   "applies combine function",
			input:  []int{5, 5, 5},
			output: 15,
			combine: func(previous, next int) int {
				return previous + next
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if o := slicesx.Reduce(tC.input, tC.combine); o != tC.output {
				t.Fail()
			}
		})
	}
}
