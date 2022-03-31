package store

import (
	"context"
	"errors"
	"gateway/gateway/model"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRoleStorePG struct {
	db *pgxpool.Pool
}

func NewUserRoleStorePG(db *pgxpool.Pool) *UserRoleStorePG {
	return &UserRoleStorePG{db}
}

func (store UserRoleStorePG) GetAllUserRole() (*model.UserRoles, error) {
	var userRoles model.UserRoles

	rows, err := store.db.Query(
		context.Background(),
		"SELECT * FROM \"UserRole\"",
	)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	for rows.Next() {
		var userRole model.UserRole

		err := rows.Scan(&userRole.ID, &userRole.Label)

		if err != nil { return nil, errors.New(err.Error()) }

		userRoles = append(userRoles, userRole)
	}

	return &userRoles, nil
}

func (store UserRoleStorePG) CreateUserRole(userRole model.UserRole) (*model.UserRole, error) {
	_, err := store.db.Exec(
		context.Background(),
		"INSERT "+
			"INTO \"UserRole\" "+
			"(id, label) "+
			"VALUES "+
			"($1, $2)",
		userRole.ID,
		userRole.Label,
	)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &userRole, nil
}

func (store UserRoleStorePG) UpdateUserRole(idUserRole uuid.UUID, body model.UserRoleUpdate) (*model.UserRole, error) {
	var userRole model.UserRole

	err := store.db.QueryRow(
		context.Background(),
		"UPDATE \"UserRole\" "+
			"SET label = $2 "+
			"WHERE id = $1",
		idUserRole,
		body.Label,
	).Scan(&userRole.ID, &userRole.Label)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &userRole, nil
}

func (store UserRoleStorePG) DeleteUserRole(idUserRole uuid.UUID) error {
	_, err := store.db.Exec(
		context.Background(),
		"DELETE FROM \"UserRole\" "+
			"WHERE id = $1",
		idUserRole,
	)

	if err != nil { return errors.New(err.Error()) }
	return nil
}
