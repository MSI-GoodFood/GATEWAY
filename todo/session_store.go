package todo

import (
	"github.com/go-redis/redis/v8"
)

type SessionStoreRedis struct {
	rdb *redis.Client
}

func NewSessionStoreRedis(db *redis.Client) *SessionStoreRedis {
	return &SessionStoreRedis{db}
}

// TODO: Implement SessionStore methods
