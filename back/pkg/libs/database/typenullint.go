package database

import (
	"database/sql"
)


// NewNullInt - Create a NullBool
func NewNullInt(i int64) sql.NullInt64 {
	if(i == 0){
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64:  i,
		Valid: true,
	}
}
