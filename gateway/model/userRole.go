package model

import (
	"github.com/gofrs/uuid"
)

type UserRole struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label" binding:"required"`
}

type UserRoleUpdate struct {
	Label string    `json:"label" binding:"required"`
}

type UserRoles []UserRole
