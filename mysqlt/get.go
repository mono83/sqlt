package mysqlt

import "github.com/mono83/sqlt"

// GetByID reads single entity from database
func GetByID[T any](getter sqlt.Getter, table string, id any) (*T, error) {
	var out T
	if err := getter.Get(&out, "SELECT * FROM "+Name(table)+" WHERE `id`=?", id); err != nil {
		return nil, err
	}
	return &out, nil
}

// MakeGetByID constructs function that reads single entity from database
func MakeGetByID[T any, I any](getter sqlt.Getter, table string) func(I) (*T, error) {
	return func(id I) (*T, error) {
		return GetByID[T](getter, table, id)
	}
}
