package database

import (
	"database/sql"
)

// NewNullString crea un string que puede ser nulo
func NewNullString(s string) sql.NullString {
    if len(s) == 0 {
        return sql.NullString{}
    }
    return sql.NullString{
         String: s,
         Valid: true,
    }
}