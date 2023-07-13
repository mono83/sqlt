package mysqlt

import (
	"database/sql"
	"github.com/mono83/sqlt/inspect"
	"strings"
)

// InspectColumns reads columns data for given table
func InspectColumns(db *sql.DB, table string) ([]inspect.Column, error) {
	rows, err := db.Query(
		"SELECT `TABLE_SCHEMA`,`TABLE_NAME`,`COLUMN_NAME`,`DATA_TYPE`,`COLUMN_TYPE`,`CHARACTER_MAXIMUM_LENGTH`,`NUMERIC_PRECISION`,`IS_NULLABLE`,`COLUMN_KEY`"+
			" FROM `information_schema`.`COLUMNS` WHERE `TABLE_SCHEMA`=DATABASE() AND `TABLE_NAME`=?",
		table,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []inspect.Column
	for rows.Next() {
		var c column
		err = rows.Scan(&c.Schema, &c.Table, &c.Column, &c.Type, &c.OriginalType, &c.CharLen, &c.NumericPrecision, &c.IsNullable, &c.IndexType)
		if err != nil {
			return nil, err
		}
		out = append(out, c.Covert())
	}
	return out, nil
}

type column struct {
	Schema           string
	Table            string
	Column           string
	Type             string
	OriginalType     string
	CharLen          *int
	NumericPrecision *int
	IsNullable       string
	IndexType        string
}

func (c column) Covert() inspect.Column {
	return inspect.Column{
		Database:     c.Schema,
		Table:        c.Table,
		Name:         c.Column,
		OriginalType: c.OriginalType,
		Type:         c.EvaluateType(),
		Size:         c.EvaluateSize(),
		Values:       c.EvaluateValues(),
		Indexed:      c.EvaluateIndexed(),
		Nullable:     c.EvaluateNullable(),
		Unsigned:     c.EvaluateUnsigned(),
	}
}

func (c column) EvaluateType() (t inspect.Type) {
	lower := strings.ToLower(c.Type)
	switch lower {
	case "bigint", "int", "mediumint", "smallint":
		t = inspect.Integer
	case "varchar", "char", "text", "tinytext", "mediumtext", "longtext":
		t = inspect.Text
	case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob":
		t = inspect.Binary
	case "enum":
		t = inspect.Enumeration
	case "set":
		t = inspect.Set
	case "double":
		t = inspect.Decimal
	}

	return
}

func (c column) EvaluateSize() int {
	if c.CharLen != nil {
		return *c.CharLen
	}
	if c.NumericPrecision != nil {
		return *c.NumericPrecision
	}
	return 0
}

func (c column) EvaluateIndexed() bool {
	return len(c.IndexType) > 0
}

func (c column) EvaluateNullable() bool {
	return c.IsNullable == "YES"
}

func (c column) EvaluateUnsigned() bool {
	return strings.HasSuffix(c.OriginalType, "unsigned")
}

func (c column) EvaluateValues() []string {
	if strings.HasPrefix(c.OriginalType, "enum") {
		return splitValues(c.OriginalType[5 : len(c.OriginalType)-1])
	} else if strings.HasPrefix(c.OriginalType, "set") {
		return splitValues(c.OriginalType[4 : len(c.OriginalType)-1])
	}

	return nil
}

func splitValues(s string) []string {
	chunks := strings.Split(s, ",")
	for i, j := range chunks {
		chunks[i] = j[1 : len(j)-1]
	}
	return chunks
}
