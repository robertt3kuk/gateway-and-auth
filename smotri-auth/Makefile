proto:
	protoc internal/pb/*.proto --go-grpc_out=:. --go-grpc_opt=paths=source_relative
	protoc internal/pb/*.proto --go_out=:. --go_opt=paths=source_relative
server:
	go run cmd/main.go
local-up:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5433/postgres?sslmode=disable' up
local-down:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5433/postgres?sslmode=disable' down
post-up:
	docker run --name=auth_db -e POSTGRES_PASSWORD='qwerty' -p 5433:5432 -d --rm postgres
post-down:
	docker stop auth_db


