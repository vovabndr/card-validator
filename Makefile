server:
	go run main.go

test:
	go test -v -cover ./...

proto:
	rm -f pb/*.go
	protoc --proto_path=proto \
		--go_out=pb           --go_opt=paths=source_relative \
		--go-grpc_out=pb      --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		proto/*.proto

docker_pull:
	docker pull alpine:latest
	docker pull golang:1.22-alpine

.PHONY: server test proto docker_pull