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
	mysqlRowFormat           = regexp.MustCompile(`(?mi)row_format\s?=\s?\w+`)
	mysqlEnumRegex           = regexp.MustCompile(`(?Ui)enum \(.*\)`)
	mysqlCommaRegex          = regexp.MustCompile(`(?Ui),\s+\)`)
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
			line = strings.ReplaceAll(line, "AUTO_INCREMENT", "PRIMARY KEY")
		}
		if strings.Contains(line, " unsigned ") {
			line = strings.ReplaceAll(line, " unsigned ", " ")
		}
		if strings.HasPrefix(line, "PRIMARY KEY ") && strings.HasSuffix(line, ",") {
			line = ""
		}
		if strings.HasPrefix(line, "KEY ") {
			line = ""
		}

		if strings.Contains(line, "bigint(20)") {
			line = strings.ReplaceAll(line, "bigint(20)", "INTEGER")
		}
		if strings.Contains(line, "BIGINT(20)") {
			line = strings.ReplaceAll(line, "BIGINT(20)", "INTEGER")
		}

		if len(line) > 0 {
			line = mysqlEngineRegex.ReplaceAllString(line, "")
		}
		if len(line) > 0 {
			line = mysqlDefaultCharsetRegex.ReplaceAllString(line, "")
		}
		if len(line) > 0 {
			line = mysqlRowFormat.ReplaceAllString(line, "")
		}
		if len(line) > 0 {
			line = mysqlEnumRegex.ReplaceAllString(line, "text")
		}

		if len(line) > 0 {
			out.WriteString(line)
			out.WriteString("\n")
		}
	}
	s = out.String()
	s = mysqlCommaRegex.ReplaceAllString(s, ")")
	return s
}
