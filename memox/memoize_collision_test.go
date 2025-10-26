package memox_test

import (
	"testing"

	"github.com/joaomagfreitas/stdx/memox"
)

func TestMemoizeCollisionTest(t *testing.T) {
	i := 0
	sum := func(n1, n2 int) {
		if i > 3 {
			t.Fatal()
		}

		i++
	}

	memo := memox.Memoize2(sum)

	memo(10, 10)
	memo(10, 20)
	memo(20, 10)
	memo(10, 20)

}
