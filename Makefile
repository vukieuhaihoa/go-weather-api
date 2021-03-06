postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -d postgres:14.1-alpine
restartdb:
	docker restart postgres14
redis:
	dock run --name redis -p 6379:6379 -d redis
restartredis:
	docker restart redis
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
test:
	go test -v ./...
server:
	go run main.go
PHONY: postgres restartdb createdb dropdb migrateup migratedown sqlc test server