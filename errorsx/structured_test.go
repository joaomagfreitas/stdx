package errorsx_test

import (
	"errors"
	"testing"

	"github.com/joaomagfreitas/stdx/errorsx"
)

type fooError struct {
	Foo string
}

func (err fooError) Error() string {
	return err.Foo
}

func TestStructuredError(t *testing.T) {
	testCases := []struct {
		desc   string
		err    error
		output string
	}{
		{
			desc:   "returns empty string if nothing is provided",
			err:    errorsx.StructuredError{},
			output: "",
		},
		{
			desc:   "includes origin if provided",
			err:    errorsx.StructuredError{Origin: "foo"},
			output: "origin: foo",
		},
		{
			desc:   "includes category if provided",
			err:    errorsx.StructuredError{Category: "bar"},
			output: "category: bar",
		},
		{
			desc:   "includes operation if provided",
			err:    errorsx.StructuredError{Operation: "foo bar"},
			output: "operation: foo bar",
		},
		{
			desc:   "includes message if provided",
			err:    errorsx.StructuredError{Message: "baz"},
			output: "message: baz",
		},
		{
			desc:   "includes wrapped error if provided",
			err:    errorsx.StructuredError{Wrapped: errors.New("lorem")},
			output: "err: lorem",
		},
		{
			desc: "full example",
			err: errorsx.StructuredError{
				Origin:    "foo",
				Category:  "bar",
				Operation: "foo bar",
				Message:   "baz",
				Wrapped:   errors.New("lorem"),
			},
			output: "origin: foo, category: bar, operation: foo bar, message: baz, err: lorem",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if s := tC.err.Error(); s != tC.output {
				t.Fatal(s)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	err := fooError{Foo: "bar"}

	werr := errorsx.Wrap(err, "", "", "", "")
	if !errors.Is(werr, err) {
		t.Fatal(err)
	}

}
