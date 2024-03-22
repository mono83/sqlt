package sqlt

import (
	"database/sql/driver"
	"errors"
	"strconv"
	"time"
)

var ErrUnixMillisUnsupportedType = errors.New("unsupported type")

// UnixMillis is type containing timestamp in unix milliseconds
type UnixMillis int64

// Int64 returns unix milliseconds as int64
func (u UnixMillis) Int64() int64 {
	return int64(u)
}

// Time return unix milliseconds formatted as UTC time
func (u UnixMillis) Time() time.Time {
	return time.UnixMilli(u.Int64()).UTC()
}

// Value is a driver.Valuer interface implementation
func (u UnixMillis) Value() (driver.Value, error) { return int64(u), nil }

// NativeSQL returns data in SQL native format
func (u UnixMillis) NativeSQL() any {
	return u.Int64()
}

// Scan is sql.Scanner interface implementation
func (u *UnixMillis) Scan(src any) error {
	switch x := src.(type) {
	case int64:
		*u = UnixMillis(x)
		return nil
	case string:
		i, err := strconv.ParseInt(x, 10, 64)
		if err == nil {
			*u = UnixMillis(i)
		}
		return err
	case []byte:
		i, err := strconv.ParseInt(string(x), 10, 64)
		if err == nil {
			*u = UnixMillis(i)
		}
		return err
	}
	return ErrUnixMillisUnsupportedType
}

// String return string representation of time in UTC
func (u UnixMillis) String() string { return u.Time().UTC().String() }
