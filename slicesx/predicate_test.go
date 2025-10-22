package slicesx_test

import (
	"strings"
	"testing"

	"github.com/joaomagfreitas/stdx/slicesx"
)

func TestAny(t *testing.T) {
	testCases := []struct {
		predicate func(string) bool
		desc      string
		input     []string
		output    bool
	}{
		{
			desc:   "returns false if nil slice is passed",
			input:  nil,
			output: false,
			predicate: func(s string) bool {
				return true
			},
		},
		{
			desc:   "returns false if empty slice is passed",
			input:  []string{},
			output: false,
			predicate: func(s string) bool {
				return true
			},
		},
		{
			desc:   "returns true if at least one item matches",
			input:  []string{"foo", "bar"},
			output: true,
			predicate: func(s string) bool {
				return strings.Contains(s, "a")
			},
		},
		{
			desc:   "returns false if no items match",
			input:  []string{"foo", "bar"},
			output: false,
			predicate: func(s string) bool {
				return strings.Contains(s, "z")
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if o := slicesx.Any(tC.input, tC.predicate); o != tC.output {
				t.Fail()
			}
		})
	}
}

func TestAll(t *testing.T) {
	testCases := []struct {
		predicate func(string) bool
		desc      string
		input     []string
		output    bool
	}{
		{
			desc:   "returns false if nil slice is passed",
			input:  nil,
			output: false,
			predicate: func(s string) bool {
				return true
			},
		},
		{
			desc:   "returns false if empty slice is passed",
			input:  []string{},
			output: false,
			predicate: func(s string) bool {
				return true
			},
		},
		{
			desc:   "returns true if all items match",
			input:  []string{"foo", "bar"},
			output: true,
			predicate: func(s string) bool {
				return len(s) > 0
			},
		},
		{
			desc:   "returns false if one item matches",
			input:  []string{"foo", "bar"},
			output: false,
			predicate: func(s string) bool {
				return strings.Contains(s, "a")
			},
		},
		{
			desc:   "returns false no item matches",
			input:  []string{"foo", "bar"},
			output: false,
			predicate: func(s string) bool {
				return strings.Contains(s, "z")
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if o := slicesx.All(tC.input, tC.predicate); o != tC.output {
				t.Fail()
			}
		})
	}
}
