package slicesx_test

import (
	"testing"

	"github.com/joaomagfreitas/stdx/slicesx"
)

func TestFold(t *testing.T) {
	testCases := []struct {
		combine func(previous int, next int) int
		desc    string
		input   []int
		initial int
		output  int
	}{
		{
			desc:    "returns initial value if nil slice is passed",
			input:   nil,
			initial: 0,
			output:  0,
			combine: func(previous, next int) int {
				return previous + next
			},
		},
		{
			desc:    "returns initial value if empty slice is passed",
			input:   []int{},
			initial: 0,
			output:  0,
			combine: func(previous, next int) int {
				return previous + next
			},
		},
		{
			desc:    "uses initial value with combine function",
			input:   []int{5, 5, 5},
			initial: 5,
			output:  20,
			combine: func(previous, next int) int {
				return previous + next
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if o := slicesx.Fold(tC.input, tC.initial, tC.combine); o != tC.output {
				t.Fail()
			}
		})
	}
}
