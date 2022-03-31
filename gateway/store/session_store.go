package store

import (
	"context"
	"errors"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid"
)

type SessionStoreRedis struct {
	rdb *redis.Client
}

func NewSessionStoreRedis(db *redis.Client) *SessionStoreRedis {
	return &SessionStoreRedis{db}
}

func (store SessionStoreRedis) GetUserByToken(token string) (userID uuid.UUID, err error) {

	id, err := store.rdb.Get(
		context.Background(),
		strings.Join(strings.Fields(token), ""),
	).Result()

	if err != nil {
		return uuid.Nil, errors.New(err.Error())
	}

	userUuid, err := uuid.FromString(id)

	if err != nil {
		return uuid.Nil, errors.New(err.Error())
	}

	return userUuid, nil
}

func (store SessionStoreRedis) AddToken(idUser uuid.UUID, token string) error {
	err := store.rdb.Set(
		context.Background(),
		strings.Join(strings.Fields(token), ""),
		idUser.String(),
		0,
	).Err()

	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (store SessionStoreRedis) RevokeToken(token string) error {
	err := store.rdb.Del(
		context.Background(),
		strings.Join(strings.Fields(token), ""),
	).Err()

	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
