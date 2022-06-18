package model

import "github.com/google/uuid"

type Stamp struct {
	ID     uuid.UUID
	Name   string
	Image  string
	UserID string
}
