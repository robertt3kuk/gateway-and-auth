proto:
	protoc pkg/**/pb/*.proto --go-grpc_out=:. 
	protoc pkg/**/pb/*.proto --go_out=:. 
server:
	go run cmd/main.go
