postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -d postgres:14.1-alpine
createdb:
	docker exec -it postgres14 createdb --username=admin --owner=admin weather
dropdb:
	docker exec -it postgres14 dropdb -U admin weather
migrateup:
	migrate -path db/migration -database "postgres://admin:admin@localhost:5432/weather?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://admin:admin@localhost:5432/weather?sslmode=disable" -verbose down
sqlc:
	sqlc generate

PHONY: postgres createdb dropdb migrateup migratedown sqlc