package main

import (
	"database/sql"
)

/*
	func FromStringToInt(queryParam string) (int, error) {
		return strconv.Atoi(queryParam)
	}
*/
func convertSQLNullInt64(n sql.NullInt64) *int {
	if n.Valid {
		num := int(n.Int64)
		return &num
	}
	return nil
}
