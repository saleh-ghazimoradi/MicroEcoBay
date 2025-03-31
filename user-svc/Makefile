gen-proto:
	protoc --proto_path=internal/gateway/grpc --go_out=. --go-grpc_out=. internal/gateway/grpc/order/*.proto


http:
	go run . http