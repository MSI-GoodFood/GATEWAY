package model

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

type UserUpdate struct {
	Password string    `json:"password"`
	Active   bool      `json:"active"`
	Role     uuid.UUID `json:"id_role"`
	Country  uuid.UUID `json:"id_country"`
}

type Users []User
