package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	Active    bool      `json:"active" default:"true"`
	Role      uuid.UUID `json:"id_role"`
	Country   uuid.UUID `json:"id_country"`
}

type UserRole struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label" binding:"required"`
}

type UserUpdate struct {
	Password string    `json:"password"`
	Active   bool      `json:"active"`
	Role     uuid.UUID `json:"id_role"`
	Country  uuid.UUID `json:"id_country"`
}

type UserRoles []UserRole
type Users []User
