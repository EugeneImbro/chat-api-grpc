compile:
	protoc --go_opt=paths=source_relative --go_out=internal/server --go-grpc_opt=paths=source_relative --go-grpc_out=internal/server chat-backend.proto