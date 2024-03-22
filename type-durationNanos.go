package sqlt

import (
	"database/sql/driver"
	"errors"
	"strconv"
	"time"
)

var ErrDurationNanosUnsupportedType = errors.New("unsupported type")

// DurationNanos is type duration in nanoseconds
type DurationNanos int64

// Int64 returns nanoseconds count
func (u DurationNanos) Int64() int64 {
	return int64(u)
}

// Duration returns Go's duration type
func (u DurationNanos) Duration() time.Duration {
	return time.Duration(u.Int64()) * time.Nanosecond
}

// Value is a driver.Valuer interface implementation
func (u DurationNanos) Value() (driver.Value, error) { return int64(u), nil }

// NativeSQL returns data in SQL native format
func (u DurationNanos) NativeSQL() any {
	return u.Int64()
}

// Scan is sql.Scanner interface implementation
func (u *DurationNanos) Scan(src any) error {
	switch x := src.(type) {
	case int64:
		*u = DurationNanos(x)
		return nil
	case string:
		i, err := strconv.ParseInt(x, 10, 64)
		if err == nil {
			*u = DurationNanos(i)
		}
		return err
	case []byte:
		i, err := strconv.ParseInt(string(x), 10, 64)
		if err == nil {
			*u = DurationNanos(i)
		}
		return err
	}
	return ErrDurationNanosUnsupportedType
}

// String return string representation of time in UTC
func (u DurationNanos) String() string { return u.Duration().String() }
