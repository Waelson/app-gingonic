// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package repository

import (
	"database/sql"
)

type Authors struct {
	ID   int64
	Name string
	Bio  sql.NullString
}
