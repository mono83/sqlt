package sqlt

import (
	"database/sql/driver"
	"errors"
	"strconv"
	"time"
)

var ErrUnixSecondsUnsupportedType = errors.New("unsupported type")

// UnixSeconds is type containing timestamp in unix seconds
type UnixSeconds int64

// Int64 returns unix seconds as int64
func (u UnixSeconds) Int64() int64 {
	return int64(u)
}

// Time return unix seconds formatted as UTC time
func (u UnixSeconds) Time() time.Time {
	return time.Unix(u.Int64(), 0).UTC()
}

// Value is a driver.Valuer interface implementation
func (u UnixSeconds) Value() (driver.Value, error) { return int64(u), nil }

// NativeSQL returns data in SQL native format
func (u UnixSeconds) NativeSQL() any {
	return u.Int64()
}

// Scan is sql.Scanner interface implementation
func (u *UnixSeconds) Scan(src any) error {
	switch x := src.(type) {
	case int64:
		*u = UnixSeconds(x)
		return nil
	case string:
		i, err := strconv.ParseInt(x, 10, 64)
		if err == nil {
			*u = UnixSeconds(i)
		}
		return err
	case []byte:
		i, err := strconv.ParseInt(string(x), 10, 64)
		if err == nil {
			*u = UnixSeconds(i)
		}
		return err
	}
	return ErrUnixSecondsUnsupportedType
}

// String return string representation of time in UTC
func (u UnixSeconds) String() string { return u.Time().UTC().String() }
