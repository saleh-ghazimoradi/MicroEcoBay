gen-proto:
	protoc --proto_path=internal/gateway/grpc --go_out=. --go-grpc_out=. internal/gateway/grpc/order/*.proto

down:
	go run . down

http:
	go run . http

test:
	go test ./... -v