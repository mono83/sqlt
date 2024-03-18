package mosaic

import (
	"errors"
	"time"
)

func MapFromString(s string) []any               { return []any{s} }
func MapFromInt(i int) []any                     { return []any{i} }
func MapFromUInt(i uint) []any                   { return []any{i} }
func MapFromUInt64(i uint64) []any               { return []any{i} }
func MapFromTimeUnixSeconds(t time.Time) []any   { return []any{t.Unix()} }
func MapFromDurationNanos(d time.Duration) []any { return []any{d.Nanoseconds()} }

func MapToString(data []any) (*string, error) {
	if len(data) != 1 {
		return nil, errors.New("wrong column count")
	}
	datum := data[0]
	switch x := datum.(type) {
	case string:
		return &x, nil
	case *string:
		if x == nil {
			return nil, errors.New("nil value")
		}
		return x, nil
	case []byte:
		s := string(x)
		return &s, nil
	default:
		return nil, errors.New("unsupported type")
	}
}

func MapToUint64(data []any) (*uint64, error) {
	if len(data) != 1 {
		return nil, errors.New("wrong column count")
	}
	datum := data[0]
	switch x := datum.(type) {
	case int64:
		i := uint64(x)
		return &i, nil
	default:
		return nil, errors.New("unsupported type")
	}
}
