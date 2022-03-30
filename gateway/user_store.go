package gateway

import (
	"context"
	"errors"
	"gateway/gateway/models"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserStorePG struct {
	db *pgxpool.Pool
}

func NewUserStorePG(db *pgxpool.Pool) *UserStorePG {
	return &UserStorePG{db}
}

func (store UserStorePG) Insert(user models.User) (*models.User, error) {
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

func (store UserStorePG) UpdateUser(id uuid.UUID, body models.UserUpdate) (*models.User, error) {
	var user models.User

	err := store.db.QueryRow(
		context.Background(),
		"UPDATE \"User\" "+
			"SET password = $2, active = $3, id_role = $4, id_country = $5 "+
			"WHERE id = $1",
		id,
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

func (store UserStorePG) FindById(id uuid.UUID) (*models.User, error) {
	var user models.User

	err := store.db.QueryRow(
		context.Background(),
		"SELECT * FROM \"User\" WHERE id = $1",
		id,
	).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.Active, &user.Role, &user.Country)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &user, nil
}

func (store UserStorePG) FindByEmail(email string) (*models.User, error) {
	var user models.User

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
