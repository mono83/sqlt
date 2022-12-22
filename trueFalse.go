package sqlt

import "database/sql/driver"

// TrueFalse wraps values stored in database as enum(true,false)
type TrueFalse bool

// Bool returns boolean representation
func (t TrueFalse) Bool() bool { return bool(t) }

// IsTrue returns boolean representation, is an alias to Bool
func (t TrueFalse) IsTrue() bool { return t.Bool() }

// NativeSQL returns data in SQL native format
func (t TrueFalse) NativeSQL() interface{} {
	return t.Bool()
}

// Value is a sql.driver.Valuer interface implementation
func (t TrueFalse) Value() (driver.Value, error) { return t.Bool(), nil }

// Scan is sql.Scanner interface implementation
func (t *TrueFalse) Scan(value interface{}) error {
	if value != nil {
		v, err := driver.Bool.ConvertValue(value)
		if err != nil {
			return err
		}

		*t = TrueFalse(v.(bool))
	}
	return nil
}

// String returns string representation of boolean value
func (t *TrueFalse) String() string {
	if t.Bool() {
		return "true"
	}
	return "false"
}
