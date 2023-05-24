package mysqlt

import (
	"strings"
)

// Name performs table/column name escaping.
func Name(s string) string {
	if !strings.HasPrefix(s, "`") {
		s = "`" + s
	}
	if !strings.HasSuffix(s, "`") {
		s += "`"
	}
	return s
}
