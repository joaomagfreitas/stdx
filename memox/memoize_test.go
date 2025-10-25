package memox_test

import (
	"testing"
	"time"

	"github.com/joaomagfreitas/stdx/memox"
)

func TestMemoize(t *testing.T) {
	i := 0
	inc := func() { i++ }
	memo := memox.Memoize(inc)

	memo()
	memo()

	if i != 1 {
		t.Fail()
	}
}

func TestMemoizeR1(t *testing.T) {
	rand := time.Now
	memo := memox.MemoizeR1(rand)

	r1 := memo()
	if r1 != memo() {
		t.Fail()
	}
}

func TestMemoizeR2(t *testing.T) {
	rand := func() (time.Time, time.Time) { return time.Now(), time.Now() }
	memo := memox.MemoizeR2(rand)

	r1, r2 := memo()
	if o1, o2 := memo(); o1 != r1 || o2 != r2 {
		t.Fail()
	}
}
