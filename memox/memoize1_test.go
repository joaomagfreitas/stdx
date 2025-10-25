package memox_test

import (
	"math/rand"
	"testing"

	"github.com/joaomagfreitas/stdx/memox"
)

func TestMemoize1(t *testing.T) {
	i := 0
	inc := func(n int) { i += n }
	memo := memox.Memoize1(inc)

	memo(10)
	memo(20)
	memo(10)

	if i != 30 {
		t.Fail()
	}
}

func TestMemoize1R1(t *testing.T) {
	rand := func(n int) int { return rand.Intn(n) }
	memo := memox.Memoize1R1(rand)

	r := memo(1000)
	if o := memo(1000); o != r {
		t.Fail()
	}
}

func TestMemoize1R2(t *testing.T) {
	rand := func(n int) (int, int) { return rand.Intn(n), rand.Intn(n + n) }
	memo := memox.Memoize1R2(rand)

	r1, r2 := memo(1000)
	if o1, o2 := memo(1000); o1 != r1 || o2 != r2 {
		t.Fail()
	}
}
