package db

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
)

type Store struct {
	*Queries
	db          *sql.DB
	RedisClient *redis.Client
}

func NewStore(db *sql.DB, redisClient *redis.Client) *Store {
	return &Store{
		db:          db,
		Queries:     New(db),
		RedisClient: redisClient,
	}
}
