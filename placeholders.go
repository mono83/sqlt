package sqlt

import "bytes"

// PlaceholdersString constructs placeholders string like ?,?,?
func PlaceholdersString(n int) string {
	switch n {
	case 1:
		return "?"
	case 2:
		return "?,?"
	case 3:
		return "?,?,?"
	case 4:
		return "?,?,?,?"
	case 5:
		return "?,?,?,?,?"
	case 6:
		return "?,?,?,?,?,?"
	case 7:
		return "?,?,?,?,?,?,?"
	case 8:
		return "?,?,?,?,?,?,?,?"
	case 9:
		return "?,?,?,?,?,?,?,?,?"
	default:
		if n < 1 {
			return ""
		}
		buf := bytes.NewBufferString("?")
		for i := 1; i < n; i++ {
			buf.WriteString(",?")
		}
		return buf.String()
	}
}
