package todo

import "github.com/gofrs/uuid"

type UserStore interface {
	Insert(user User) (*User, error)
	FindById(id uuid.UUID) (*User, error)
	FindByEmail(email string) (*User, error)
}

type SessionStore interface {
	Add(userID uuid.UUID, token string) error
	Revoke(token string) error
	FindByToken(token string) (userID uuid.UUID, err error)
}

type TodoStore interface {
	Add(Todo) (*Todo, error)
	Delete(todoID uuid.UUID) error
	Toggle(todoID uuid.UUID, done bool) (*Todo, error)
	UpdateText(todoID uuid.UUID, text string) (*Todo, error)

	FindByID(todoID uuid.UUID) (*Todo, error)
	FindByUserID(userID uuid.UUID) ([]*Todo, error)
}
