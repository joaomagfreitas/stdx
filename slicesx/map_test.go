package slicesx_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/joaomagfreitas/stdx/slicesx"
)

func TestMap(t *testing.T) {
	testCases := []struct {
		transform func(string) string
		desc      string
		input     []string
		output    []string
	}{
		{
			desc:   "returns nil if nil slice is passed",
			input:  nil,
			output: nil,
			transform: func(s string) string {
				return s
			},
		},
		{
			desc:   "applies mapper callback with input slice",
			input:  []string{"foo", "bar"},
			output: []string{"FOO", "BAR"},
			transform: func(s string) string {
				return strings.ToUpper(s)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if o := slicesx.Map(tC.input, tC.transform); !slices.Equal(tC.output, o) {
				t.Fail()
			}
		})
	}
}
