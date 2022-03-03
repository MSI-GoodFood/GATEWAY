package todo

import (
	"context"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserStorePG struct {
	db *pgxpool.Pool
}

func (store UserStorePG) Insert(user User) (*User, error) {
	_, err := store.db.Exec(
		context.Background(),
		"INSERT INTO \"User\" (id, email) VALUES ($1, $2)",
		user.ID,
		user.Email,
	)

	if err != nil { return nil, errors.New(err.Error()) }
	return &user, nil
}

func (store UserStorePG) FindById(id uuid.UUID) (*User, error) {
	var user User

	err := store.db.QueryRow(
		context.Background(),
		"SELECT * FROM \"User\" WHERE id = $1",
		id,
	).Scan(&user.ID, &user.Email)

	if err != nil { return nil, errors.New(err.Error()) }
	return &user, nil
}

func (store UserStorePG) FindByEmail(email string) (*User, error) {
	var user User

	err := store.db.QueryRow(
		context.Background(),
		"SELECT * FROM \"User\" WHERE email = $1",
		email,
	).Scan(&user.ID, &user.Email)

	if err != nil { return nil, errors.New(err.Error()) }
	return &user, nil
}

func NewUserStorePG(db *pgxpool.Pool) *UserStorePG {
	return &UserStorePG{db}
}
