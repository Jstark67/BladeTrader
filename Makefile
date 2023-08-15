postgres: 
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=75Cheetah -d postgres:15-alpine


createdb:
	docker exec -it postgres createdb --username=root --owner=root backend_bank

dropdb:
	docker exec -it postgres dropdb backend_bank

migrateup:
	migrate -path Db/migration -database "postgresql://root:75Cheetah@localhost:5432/backend_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path Db/migration -database "postgresql://root:75Cheetah@localhost:5432/backend_bank?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

sqlc:
	sqlc generate
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test