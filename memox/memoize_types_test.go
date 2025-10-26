package memox_test

import (
	"math/rand"
	"testing"

	"github.com/joaomagfreitas/stdx/memox"
)

func TestMemoizeSupportedTypes(t *testing.T) {
	rand := func(v any) int { return rand.Int() }
	ptr := 0

	testCases := []struct {
		arg1 any
		desc string
	}{
		{
			desc: "supports nil values",
			arg1: nil,
		},
		{
			desc: "supports byte values",
			arg1: byte(127),
		},
		{
			desc: "supports rune values",
			arg1: rune(127),
		},
		{
			desc: "supports string values",
			arg1: "foo.bar",
		},
		{
			desc: "supports false values",
			arg1: false,
		},
		{
			desc: "supports true values",
			arg1: true,
		},
		{
			desc: "supports uint8 values",
			arg1: uint8(16),
		},
		{
			desc: "supports uint16 values",
			arg1: uint16(16),
		},
		{
			desc: "supports uint32 values",
			arg1: uint32(16),
		},
		{
			desc: "supports uint64 values",
			arg1: uint64(16),
		},
		{
			desc: "supports uint values",
			arg1: uint(16),
		},
		{
			desc: "supports uintptr values",
			arg1: uintptr(16),
		},
		{
			desc: "supports int8 values",
			arg1: int8(16),
		},
		{
			desc: "supports int16 values",
			arg1: int16(16),
		},
		{
			desc: "supports int32 values",
			arg1: int32(16),
		},
		{
			desc: "supports int64 values",
			arg1: int64(16),
		},
		{
			desc: "supports int values",
			arg1: int(16),
		},
		{
			desc: "supports float32 values",
			arg1: float32(16),
		},
		{
			desc: "supports float64 values",
			arg1: float64(16),
		},
		{
			desc: "supports complex64 values",
			arg1: complex64(16),
		},
		{
			desc: "supports complex128 values",
			arg1: complex128(16),
		},
		{
			desc: "supports string slice values",
			arg1: []string{"foo", "bar"},
		},
		{
			desc: "supports byte slice values",
			arg1: []byte{0, 1, 2},
		},
		{
			desc: "supports int slice values",
			arg1: []int{0, 1, 2},
		},
		{
			desc: "supports empty struct values",
			arg1: struct{}{},
		},
		{
			desc: "supports complex struct values",
			arg1: struct {
				A2 string
				A3 []float32
				A1 int
				A4 bool
			}{
				A1: 32,
				A2: "foo.bar",
				A3: []float32{16, 32, 64},
				A4: false,
			},
		},
		{
			desc: "supports function values",
			arg1: rand,
		},
		{
			desc: "supports anonymous values",
			arg1: func() {},
		},
		{
			desc: "supports pointers",
			arg1: &ptr,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			memo := memox.Memoize1R1(rand)
			o1 := memo(tC.arg1)
			o2 := memo(tC.arg1)

			if o1 != o2 {
				t.Fail()
			}
		})
	}
}

func TestMemoizeUnsupportedTypes(t *testing.T) {
	rand := func(v any) int { return rand.Int() }

	testCases := []struct {
		arg1 any
		desc string
	}{
		{
			desc: "does not support unexported fields in structs",
			arg1: struct {
				a2 string
				a3 []float32
				a1 int
				a4 bool
			}{
				a1: 32,
				a2: "foo.bar",
				a3: []float32{16, 32, 64},
				a4: false,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer func() {
				recover()
			}()

			memo := memox.Memoize1R1(rand)
			o1 := memo(tC.arg1)
			o2 := memo(tC.arg1)

			if o1 == o2 {
				t.Fail()
			}
		})
	}
}
