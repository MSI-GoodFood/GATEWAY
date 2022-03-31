package model

import "github.com/gofrs/uuid"

type Account struct {
	User  *User  `json:"user" binding:"required"`
	Token string `json:"token" binding:"required"`
}

type UserToken struct {
	User  uuid.UUID `json:"id" binding:"required"`
	Token string    `json:"token" binding:"required"`
}
