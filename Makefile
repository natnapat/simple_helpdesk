postgres:
	docker run --name postgres14.5 --network helpdesk-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecret -d postgres:14.5-alpine

createdb:
	docker exec -it postgres14.5 createdb --username=root --owner=root simple_helpdesk

dropdb:
	docker exec -it postgres14.5 dropdb simple_helpdesk

migrateup:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/simple_helpdesk?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/simple_helpdesk?sslmode=disable" -verbose down

.PHONNY: postgres createdb dropdb migrateup migratedown