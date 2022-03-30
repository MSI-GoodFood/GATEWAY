package todo

import "github.com/gofrs/uuid"

type User struct {
	ID uuid.UUID `json:"id"`
	Email string `json:"email" default:""`
}

type Users []User

type Account struct {
	User  *User   `json:"user"`
	Token string `json:"token"`
}

type UserToken struct {
	User uuid.UUID `json:"id"`
	Token string `json:"token"`
}
