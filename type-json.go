package sqlt

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

var ErrJSONUnsupportedType = errors.New("unsupported type")

// JSON wraps value stored in JSON format
type JSON string

// String return JSON string
func (j JSON) String() string { return string(j) }

// Unmarshal reads value into given target struct
func (j JSON) Unmarshal(target any) error {
	return json.Unmarshal([]byte(j), target)
}

// NativeSQL returns data in SQL native format
func (j JSON) NativeSQL() any {
	return j.String()
}

// Value is a sql.driver.Valuer interface implementation
func (j JSON) Value() (driver.Value, error) { return j.String(), nil }

// Scan is sql.Scanner interface implementation
func (j *JSON) Scan(src any) error {
	switch x := src.(type) {
	case string:
		*j = JSON(x)
		return nil
	case []byte:
		*j = JSON(string(x))
		return nil
	}
	return ErrJSONUnsupportedType
}
