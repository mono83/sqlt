package sqlt

import (
	"errors"
	"strconv"
	"time"
)

var ErrUnsupportedTypeForUnixSeconds = errors.New("unsupported type")

// UnixSeconds is type containing timestamp in unix seconds
type UnixSeconds int64

// String returns string representation of time
func (u UnixSeconds) String() string {
	return u.Time().String()
}

// Int64 returns unix seconds as int64
func (u UnixSeconds) Int64() int64 {
	return int64(u)
}

// Time return unix seconds formatted as UTC time
func (u UnixSeconds) Time() time.Time {
	return time.Unix(u.Int64(), 0).UTC()
}

// NativeSQL returns data in SQL native format
func (u UnixSeconds) NativeSQL() interface{} {
	return u.Int64()
}

// Scan is sql.Scanner interface implementation
func (u *UnixSeconds) Scan(src interface{}) error {
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
	return ErrUnsupportedTypeForUnixSeconds
}
