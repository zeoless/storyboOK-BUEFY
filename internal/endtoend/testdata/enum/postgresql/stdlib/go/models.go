// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package querytest

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type Size string

const (
	SizeXSmall Size = "x-small"
	SizeSmall  Size = "small"
	SizeMedium Size = "medium"
	SizeLarge  Size = "large"
	SizeXLarge Size = "x-large"
)

func (e *Size) Scan(src interface{}) error