package sqlitet

import (
	"bufio"
	"bytes"
	"regexp"
	"strings"
)

// FromMySQL converts query from MySQL format into SQLite.
//
// Experimental: this feature is experimental and not recommended
// to use due not final API.
func FromMySQL(s string) string {
	if strings.Contains(s, "CREATE TABLE") {
		s = FromMySQLCreateTable(s)
	}

	return s
}

var (
	mysqlEngineRegex         = regexp.MustCompile(`(?mi)engine\s?=\s?\w+`)
	mysqlDefaultCharsetRegex = regexp.MustCompile(`(?mi)default charset\s?=\s?\w+`)
)

// FromMySQLCreateTable converts SQL CREATE TABLE query from
// MySQL format into SQLite.
//
// Experimental: this feature is experimental and not recommended
// to use due not final API.
func FromMySQLCreateTable(s string) string {
	out := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if strings.Contains(line, "AUTO_INCREMENT") {
			line = strings.ReplaceAll(line, "AUTO_INCREMENT", "")
		}
		if strings.Contains(line, " unsigned ") {
			line = strings.ReplaceAll(line, " unsigned ", " ")
		}
		if strings.HasPrefix(line, "PRIMARY KEY ") && strings.HasSuffix(line, ",") {
			line = line[0 : len(line)-1]
		}
		if strings.HasPrefix(line, "KEY ") {
			line = ""
		}

		if len(line) > 0 {
			line = mysqlEngineRegex.ReplaceAllString(line, "")
		}
		if len(line) > 0 {
			line = mysqlDefaultCharsetRegex.ReplaceAllString(line, "")
		}

		if len(line) > 0 {
			out.WriteString(line)
			out.WriteString("\n")
		}
	}
	return out.String()
}
