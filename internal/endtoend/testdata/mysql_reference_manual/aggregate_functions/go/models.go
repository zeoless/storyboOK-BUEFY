
// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package aggregate_functions

import (
	"database/sql"
)

type Student struct {
	StudentName sql.NullString
	Score       sql.NullFloat64
}