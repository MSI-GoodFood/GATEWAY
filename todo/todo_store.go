package todo

import (
	"context"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TodoStorePG struct {
	db *pgxpool.Pool
}

func NewTodoStorePG(db *pgxpool.Pool) *TodoStorePG {
	return &TodoStorePG{db}
}

func (store TodoStorePG) Add(todo Todo) (*Todo, error) {
	_, err := store.db.Exec(
		context.Background(),
		"INSERT INTO \"Todo\" (id, text, user_id) VALUES ($1, $2, $3)",
		todo.ID,
		todo.Text,
		todo.UserId,
	)

	if err != nil { return nil, errors.New(err.Error()) }
	return &todo, nil
}

func (store TodoStorePG) Delete(todoID uuid.UUID) error {
	_, err := store.db.Exec(
		context.Background(),
		"DELETE FROM \"Todo\" WHERE id = $1",
		todoID,
	)

	if err != nil { return errors.New(err.Error()) }
	return nil
}

func (store TodoStorePG) Toggle(todoID uuid.UUID, done int) (*Todo, error) {
	var todo Todo

	_, err := store.db.Exec(
		context.Background(),
		"UPDATE \"Todo\" SET done=$1 WHERE id = $2",
		done,
		todoID,
	)

	if err != nil { return nil, errors.New(err.Error()) }

	err = store.db.QueryRow(
		context.Background(),
		"SELECT * FROM \"Todo\" WHERE id = $1",
		todoID,
	).Scan(&todo.ID, &todo.Text, &todo.Done, &todo.UserId)

	if err != nil { return nil, errors.New(err.Error()) }
	return &todo, nil
}

func (store TodoStorePG) UpdateText(todoID uuid.UUID, text string) (*Todo, error) {
	var todo Todo

	_, err := store.db.Exec(
		context.Background(),
		"UPDATE \"Todo\" SET text=$1 WHERE id = $2",
		text,
		todoID,
	)

	if err != nil { return nil, errors.New(err.Error()) }

	err = store.db.QueryRow(
		context.Background(),
		"SELECT * FROM \"Todo\" WHERE id = $1",
		todoID,
	).Scan(&todo.ID, &todo.Text, &todo.Done, &todo.UserId)

	if err != nil { return nil, errors.New(err.Error()) }
	return &todo, nil
}

func (store TodoStorePG) FindByID(todoID uuid.UUID) (*Todo, error) {
	var todo Todo

	err := store.db.QueryRow(
		context.Background(),
		"SELECT * FROM \"Todo\" WHERE id = $1",
		todoID,
	).Scan(&todo.ID, &todo.Text, &todo.Done, &todo.UserId)

	if err != nil { return nil, errors.New(err.Error()) }
	return &todo, nil
}

func (store TodoStorePG) FindByUserID(userID uuid.UUID) (*Todos, error) {
	var todos Todos

	rows, err := store.db.Query(
		context.Background(),
		"SELECT * FROM \"Todo\" WHERE user_id = $1",
		userID,
	)
	if err != nil { return nil, errors.New(err.Error()) }

	for rows.Next() {
		var todo Todo

		err := rows.Scan(&todo.ID, &todo.Text, &todo.Done, &todo.UserId)

		if err != nil { return nil, errors.New(err.Error()) }

		todos = append(todos, todo)
	}

	return &todos, nil
}

func (store TodoStorePG) FindById(id uuid.UUID) (*Todo, error) {
	var todo Todo

	err := store.db.QueryRow(
		context.Background(),
		"SELECT * FROM \"Todo\" WHERE id = $1",
		id,
	).Scan(&todo.ID, &todo.Text, &todo.Done, &todo.UserId)

	if err != nil { return nil, errors.New(err.Error()) }
	return &todo, nil
}
