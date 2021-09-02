### Database migrations

```
migrate -path ./migrations -database postgresql://postgres:Pass@word@localhost:5432/postgres?sslmode=disable up
```

### Run local environment

```
docker-compose -f local-environment/postgresql.yml up -d
```

### gRPC

```
protoc --go_opt=paths=source_relative --go_out=internal/server --go-grpc_opt=paths=source_relative --go-grpc_out=internal/server chat-backend.proto --go-grpc_opt=require_unimplemented_servers=false
```