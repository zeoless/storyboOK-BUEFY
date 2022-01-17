// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package querytest

import (
	"database/sql/driver"
	"fmt"
)

// this is the mood type
type FooMood string

const (
	FooMoodSad   FooMood = "sad"
	FooMoodOk    FooMood = "ok"
	FooMoodHappy FooMood = "happy"
)

func (e *FooMood) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FooMood(s)
	case string:
		*e = FooMood(s)
	default:
		return fmt.Errorf("unsupported scan type for FooMood: %T", src)
	}
	return nil
}

type NullFooMood struct {
	FooMood FooMood
	Valid   bool // Valid is true if FooMood is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFooMood) Scan(value interface{}) error {
	if value == nil {
		ns.FooMood, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.FooMood.Scan(value)
}

// Valu