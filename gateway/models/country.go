package models

import (
	"github.com/gofrs/uuid"
)

type Country struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label" binding:"required"`
	Code  string    `json:"code" binding:"required"`
}

type Countries []Country
