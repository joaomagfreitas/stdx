package errorsx_test

import (
	"errors"
	"testing"

	"github.com/joaomagfreitas/stdx/errorsx"
)

func TestMustExample(t *testing.T) {
	div := func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, errors.New("can't divide by 0")
		}

		return a / b, nil
	}

	r := errorsx.Must(div(10, 5))
	if r != 2 {
		t.Fail()
	}

	defer func() { recover() }()
	r = errorsx.Must(div(10, 0))
	if r == 0 {
		t.Fail()
	}
}
