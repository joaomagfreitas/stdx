package memox

import (
	"bytes"
	"encoding/gob"
	"math"
)

const (
	offset64 = uint64(14695981039346656037)
	prime64  = uint64(1099511628211)
)

// Computes the FNV-1a hash of an arbitrary value. Works with all builtin types, otherwise it uses `gob` encoder
// to translate the value to binary.
//
// This function panics if the value can't be encoded to binary.
func fnv64a(val any) uint64 {
	h := offset64

	switch v := val.(type) {
	case nil:
		h = h * prime64
	case bool:
		if v {
			h = (h ^ 1) * prime64
		} else {
			h = h * prime64
		}

	case int:
		h = (h ^ uint64(v)) * prime64
	case int8:
		h = (h ^ uint64(v)) * prime64
	case int16:
		h = (h ^ uint64(v)) * prime64
	case int32:
		h = (h ^ uint64(v)) * prime64
	case int64:
		h = (h ^ uint64(v)) * prime64

	case uint:
		h = (h ^ uint64(v)) * prime64
	case uint8:
		h = (h ^ uint64(v)) * prime64
	case uint16:
		h = (h ^ uint64(v)) * prime64
	case uint32:
		h = (h ^ uint64(v)) * prime64
	case uint64:
		h = (h ^ v) * prime64
	case uintptr:
		h = (h ^ uint64(v)) * prime64

	case float32:
		h = (h ^ uint64(math.Float32bits(v))) * prime64
	case float64:
		h = (h ^ math.Float64bits(v)) * prime64

	case complex64:
		h = (h ^ uint64(math.Float32bits(real(v)))) * prime64
		h = (h ^ uint64(math.Float32bits(imag(v)))) * prime64
	case complex128:
		h = (h ^ math.Float64bits(real(v))) * prime64
		h = (h ^ math.Float64bits(imag(v))) * prime64

	case string:
		for i := 0; i < len(v); i++ {
			h = (h ^ uint64(v[i])) * prime64
		}

	case []byte:
		for i := 0; i < len(v); i++ {
			h = (h ^ uint64(v[i])) * prime64
		}

	default:
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)

		err := enc.Encode(v)
		if err != nil {
			panic(err)
		}

		return fnv64a(buf.Bytes())
	}

	return h
}

func check(v any) {
	switch v.(type) {
	case int:
		println("int")
	case uint:
		println("uint")
	default:
		println("none")
	}
}
