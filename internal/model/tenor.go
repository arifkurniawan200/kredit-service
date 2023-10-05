package models

import "time"

// Tenor adalah model struct untuk tabel tenor
type Tenor struct {
	ID        int        `json:"id"`
	Tenor     int        `json:"tenor"`
	Value     int        `json:"value"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
