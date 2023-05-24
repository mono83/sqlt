package mysqlt

import (
	"strings"
)

func drain[T any](in []T) []any {
	l := len(in)
	if l == 0 {
		return nil
	} else if l == 1 {
		return []any{in[0]}
	}

	out := make([]any, l)
	for i, j := range in {
		out[i] = j
	}
	return out
}

// In constructs SQL query chunk and arguments slice
func In(values ...any) (string, []any) {
	l := len(values)
	if l == 0 {
		return "", nil
	} else if l == 1 {
		return "=?", values
	}
	return " IN (?" + strings.Repeat(",?", l-1) + ")", values
}

// InString constructs SQL query chunk and arguments slice
func InString(values ...string) (string, []any) {
	l := len(values)
	if l == 0 {
		return "", nil
	} else if l == 1 {
		return "=?", []any{values[0]}
	}

	return " IN (?" + strings.Repeat(",?", l-1) + ")", drain(values)
}

// InUint64 constructs SQL query chunk and arguments slice
func InUint64(values ...uint64) (string, []any) {
	l := len(values)
	if l == 0 {
		return "", nil
	} else if l == 1 {
		return "=?", []any{values[0]}
	}

	return " IN (?" + strings.Repeat(",?", l-1) + ")", drain(values)
}
