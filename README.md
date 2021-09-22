#Chat Backend
API (GRPC) for simple chat application. Users, Channels, Messages

####Database migrations

```
migrate -path ./migrations -database postgresql://postgres:Pass@word@localhost:5432/postgres?sslmode=disable up
```

#### Run local environment

```
docker-compose -f local-environment/postgresql.yml up -d
```