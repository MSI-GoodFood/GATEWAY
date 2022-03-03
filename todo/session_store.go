package todo

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid"
)

type SessionStoreRedis struct {
	rdb *redis.Client
}

func NewSessionStoreRedis(db *redis.Client) *SessionStoreRedis {
	return &SessionStoreRedis{db}
}

func (store SessionStoreRedis) Add(userID uuid.UUID, token string) error {
	err := store.rdb.Set(
		context.Background(),
		token,
		userID.String(),
		0,
	).Err()

	if err != nil { return errors.New(err.Error()) }
	return nil
}

func (store SessionStoreRedis) Revoke(token string) error {
	err := store.rdb.Del(
		context.Background(),
		token,
	).Err()

	if err != nil { return errors.New(err.Error()) }
	return nil
}

func (store SessionStoreRedis) FindByToken(token string) (userID uuid.UUID, err error) {
	id, err := store.rdb.Get(
		context.Background(),
		token,
	).Result()

	if err != nil { return uuid.Nil, errors.New(err.Error()) }

	userUuid, err := uuid.FromString(id)

	if err != nil { return uuid.Nil, errors.New(err.Error()) }

	return userUuid, nil
}
