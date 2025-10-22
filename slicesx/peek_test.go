package slicesx_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/joaomagfreitas/stdx/slicesx"
)

func TestFirst(t *testing.T) {
	testCases := []struct {
		predicate func(string) bool
		desc      string
		output    string
		input     []string
		ok        bool
	}{
		{
			desc:   "returns empty if nil slice is passed",
			input:  nil,
			output: "",
			ok:     false,
			predicate: func(s string) bool {
				return true
			},
		},
		{
			desc:   "returns first matched element",
			input:  []string{"foo", "bar", "baz"},
			output: "bar",
			ok:     true,
			predicate: func(s string) bool {
				return strings.Contains(s, "a")
			},
		},
		{
			desc:   "returns empty because no item matches predicate",
			input:  []string{"foo", "bar", "baz"},
			output: "",
			ok:     false,
			predicate: func(s string) bool {
				return strings.Contains(s, "aole")
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if o, ok := slicesx.First(tC.input, tC.predicate); o != tC.output || ok != tC.ok {
				t.Fail()
			}
		})
	}
}

func TestLast(t *testing.T) {
	testCases := []struct {
		predicate func(string) bool
		desc      string
		output    string
		input     []string
		ok        bool
	}{
		{
			desc:   "returns empty if nil slice is passed",
			input:  nil,
			output: "",
			ok:     false,
			predicate: func(s string) bool {
				return true
			},
		},
		{
			desc:   "returns last matched element",
			input:  []string{"foo", "bar", "baz"},
			output: "baz",
			ok:     true,
			predicate: func(s string) bool {
				return strings.Contains(s, "a")
			},
		},
		{
			desc:   "returns empty because no item matches predicate",
			input:  []string{"foo", "bar", "baz"},
			output: "",
			ok:     false,
			predicate: func(s string) bool {
				return strings.Contains(s, "aole")
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if o, ok := slicesx.Last(tC.input, tC.predicate); o != tC.output || ok != tC.ok {
				fmt.Printf("o: %v\n", o)
				t.Fail()
			}
		})
	}
}
