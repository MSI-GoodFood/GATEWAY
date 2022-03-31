package store

import (
	"context"
	"errors"
	"github.com/MSI-GoodFood/GATEWAY/gateway/model"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserStorePG struct {
	db *pgxpool.Pool
}

func NewUserStorePG(db *pgxpool.Pool) *UserStorePG {
	return &UserStorePG{db}
}

func (store UserStorePG) GetUserById(idUser uuid.UUID) (*model.User, error) {
	var user model.User

	err := store.db.QueryRow(
		context.Background(),
		"SELECT * FROM \"User\" WHERE id = $1",
		idUser,
	).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.Active, &user.Role, &user.Country)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &user, nil
}

func (store UserStorePG) GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	err := store.db.QueryRow(
		context.Background(),
		"SELECT * FROM \"User\" WHERE email = $1",
		email,
	).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.Active, &user.Role, &user.Country)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &user, nil
}

func (store UserStorePG) CreateUser(user model.User) (*model.User, error) {
	_, err := store.db.Exec(
		context.Background(),
		"INSERT "+
			"INTO \"User\" "+
			"(id, email, password, created_at, active, id_role, id_country) "+
			"VALUES "+
			"($1, $2, $3, $4, $5, $6, $7)",
		user.ID,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.Active,
		user.Role,
		user.Country,
	)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &user, nil
}

func (store UserStorePG) UpdateUser(idUser uuid.UUID, body model.UserUpdate) (*model.User, error) {
	var user model.User

	err := store.db.QueryRow(
		context.Background(),
		"UPDATE \"User\" "+
			"SET password = $2, active = $3, id_role = $4, id_country = $5 "+
			"WHERE id = $1",
		idUser,
		body.Password,
		body.Active,
		body.Role,
		body.Country,
	).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.Active, &user.Role, &user.Country)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &user, nil
}
