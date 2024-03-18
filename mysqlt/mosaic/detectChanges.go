package mosaic

import (
	"reflect"
)

func detectChanges(ex []existing, data [][]any) (update []uint64, del []uint64, insert [][]any) {
	// Determining what to update and what to insert
	for _, datum := range data {
		processed := false
		for _, e := range ex {
			if sliceEquals(datum, e.Data) {
				update = append(update, e.ID)
				processed = true
				break
			}
		}
		if !processed {
			insert = append(insert, datum)
		}
	}

	// Determining what to del
	for _, e := range ex {
		processed := false
		for _, datum := range data {
			if sliceEquals(datum, e.Data) {
				processed = true
				break
			}
		}

		if !processed {
			del = append(del, e.ID)
		}
	}

	return
}

func sliceEquals(a, b []any) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) > 0 {
		for i := range a {
			if !equals(a[i], b[i]) {
				return false
			}
		}
	}
	return true
}

func equals(a, b any) bool {
	// Type coercion
	if xa, ok := a.(string); ok {
		if xb, ok := b.([]byte); ok {
			return xa == string(xb)
		}
	}
	if xa, ok := a.(uint64); ok {
		if xb, ok := b.(int64); ok {
			return xa == uint64(xb)
		}
	}

	if !reflect.DeepEqual(a, b) {
		return false
	}

	return true
}

type existing struct {
	ID   uint64
	Data []any
}
