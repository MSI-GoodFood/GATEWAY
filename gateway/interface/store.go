package _interface

import (
	"gateway/gateway/model"

	"github.com/gofrs/uuid"
)

type UserStore interface {
	GetUserById(idUser uuid.UUID) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user model.User) (*model.User, error)
	UpdateUser(idUser uuid.UUID, body model.UserUpdate) (*model.User, error)
}

type SessionStore interface {
	GetUserByToken(token string) (userID uuid.UUID, err error)
	AddToken(idUser uuid.UUID, token string) error
	RevokeToken(token string) error
}
