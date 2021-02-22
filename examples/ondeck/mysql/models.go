// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package ondeck

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type VenuesStatus string

const (
	VenuesStatusOpen   VenuesStatus = "open"
	VenuesStatusClosed VenuesStatus = "closed"
)

func (e *VenuesStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VenuesStatus(s)
	case string:
		*e = VenuesStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for VenuesStatus: %T", src)
	}
	return nil
}

type NullVenuesStatus struct {
	VenuesStatus VenuesStatus
	Valid        bool // Valid is true if VenuesStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVenuesStatus) Scan(value interface{}) error {
	if value == nil {
		ns.VenuesStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VenuesStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVenuesStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VenuesStatus), nil
}

type City struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// Venues are places where muisc happens
type Venue struct {
	ID int64 `json:"id"`
	// Venues can be either open or closed
	Status   VenuesStatus   `json:"status"`
	Statuses sql.NullString `json:"statuses"`
	// This value appears in public URLs
	Slug            string         `json:"slug"`
	Name            string         `json:"name"`
	City            string         `json:"city"`
	SpotifyPlaylist string         `json:"spotify_playlist"`
	SongkickID      sql.NullString `json:"songkick_id"`
	Tags            sql.NullString `json:"tags"`
	CreatedAt       time.Time      `json:"created_at"`
}