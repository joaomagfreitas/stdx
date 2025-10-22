package slicesx_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/joaomagfreitas/stdx/slicesx"
)

func TestFilter(t *testing.T) {
	testCases := []struct {
		predicate func(string) bool
		desc      string
		input     []string
		output    []string
	}{
		{
			desc:   "returns nil if nil slice is passed",
			input:  nil,
			output: nil,
			predicate: func(s string) bool {
				return true
			},
		},
		{
			desc:   "applies predicate callback with input slice",
			input:  []string{"foo", "bar", "zol"},
			output: []string{"foo", "zol"},
			predicate: func(s string) bool {
				return strings.Contains(s, "o")
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if o := slicesx.Filter(tC.input, tC.predicate); !slices.Equal(tC.output, o) {
				t.Fail()
			}
		})
	}
}
