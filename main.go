package main

import (
	"database/sql"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/vukieuhaihoa/go-weather-api/api"
	db "github.com/vukieuhaihoa/go-weather-api/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://admin:admin@localhost:5432/weather?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("can not connect to db:", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "",
		DB:       0,
	})

	store := db.NewStore(conn, redisClient)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("can not start server:", err)
	}

}
