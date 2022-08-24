package sqlt

import (
	"database/sql"
	"reflect"
	"strconv"
	"strings"
)

var (
	typeSQLRawBytes  = reflect.TypeOf(sql.RawBytes{})
	typeSQLNullInt64 = reflect.TypeOf(sql.NullInt64{})
)

// StdConvert performs conversion from interface{} to defined
// in sql.ColumnType type using standard ruleset.
func StdConvert(src interface{}, def *sql.ColumnType) interface{} {
	if def != nil && def.ScanType() != nil {
		t := def.ScanType()
		switch source := src.(type) {
		case []uint8:
			switch t.Kind() {
			case reflect.Uint64:
				x, _ := strconv.ParseUint(string(source), 10, 64)
				return x
			case reflect.Uint32:
				x, _ := strconv.ParseUint(string(source), 10, 32)
				return x
			case reflect.Struct:
				switch t {
				case typeSQLNullInt64:
					if src != nil {
						x, _ := strconv.ParseInt(string(source), 10, 64)
						return &x
					}
				}
			case reflect.Slice:
				switch t {
				case typeSQLRawBytes:
					switch strings.ToLower(def.DatabaseTypeName()) {
					case "char", "varchar", "text":
						return string(source)
					}
				}
			}
		}
	}
	return src
}
