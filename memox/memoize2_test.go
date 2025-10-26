package memox_test

import (
	"math/rand"
	"testing"

	"github.com/joaomagfreitas/stdx/memox"
)

func TestMemoize2(t *testing.T) {
	i := 0
	inc := func(n1, n2 int) { i += n1 + n2 }
	memo := memox.Memoize2(inc)

	memo(10, 20)
	memo(20, 20)
	memo(20, 10)
	memo(10, 20)

	if i != 100 {
		t.Fail()
	}
}

func TestMemoize2R1(t *testing.T) {
	rand := func(n1, n2 int) int { return rand.Intn(n1) + rand.Intn(n2) }
	memo := memox.Memoize2R1(rand)

	r := memo(1000, 2000)
	if o := memo(1000, 2000); o != r {
		t.Fail()
	}
}

func TestMemoize2R2(t *testing.T) {
	rand := func(n1, n2 int) (int, int) { return rand.Intn(n1), rand.Intn(n2) }
	memo := memox.Memoize2R2(rand)

	r1, r2 := memo(1000, 2000)
	if o1, o2 := memo(1000, 2000); o1 != r1 || o2 != r2 {
		t.Fail()
	}
}
