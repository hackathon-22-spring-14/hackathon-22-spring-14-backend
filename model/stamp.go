package model

import "github.com/google/uuid"

type Stamp struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	ImageURL string    `json:"-" db:"image_url"`
	Image    string    `json:"image" db:"-"`
}
