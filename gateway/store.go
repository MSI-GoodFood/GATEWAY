package gateway

import (
	"gateway/gateway/models"

	"github.com/gofrs/uuid"
)

type UserStore interface {
	Insert(user models.User) (*models.User, error)
	FindById(id uuid.UUID) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	UpdateUser(id uuid.UUID, body models.UserUpdate) (*models.User, error)
}

type SessionStore interface {
	Add(userID uuid.UUID, token string) error
	Revoke(token string) error
	FindByToken(token string) (userID uuid.UUID, err error)
}
