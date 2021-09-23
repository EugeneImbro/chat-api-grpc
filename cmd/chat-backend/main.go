package main

import (
	"fmt"
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/EugeneImbro/chat-backend/internal/repository/postgres"
	"github.com/EugeneImbro/chat-backend/internal/server"
	service "github.com/EugeneImbro/chat-backend/internal/service/impl"
)

func main() {
	db, err := sqlx.Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		logrus.WithError(err).Fatal("database initialization error")
	}

	if err = db.Ping(); err != nil {
		logrus.WithError(err).Fatal("database is not available")
	}

	repo := postgres.New(db)
	us := service.NewUserService(repo)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		logrus.WithError(err).Fatal("listener initialization error")
	}
	s := grpc.NewServer()
	server.RegisterUserServiceServer(s, server.NewServer(us))
	logrus.Infoln("server initialized")

	if err := s.Serve(listener); err != nil {
		logrus.WithError(err).Fatal("failed to serve")
	}
}
