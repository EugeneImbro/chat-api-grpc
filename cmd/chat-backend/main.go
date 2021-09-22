package main

import (
	"fmt"
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/EugeneImbro/chat-backend/internal/repository/postgres"
	"github.com/EugeneImbro/chat-backend/internal/server"
	"github.com/EugeneImbro/chat-backend/internal/service"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.WithError(err).Fatal("config initialization error")
	}

	db, err := sqlx.Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		logrus.WithError(err).Fatal("database initialization error")
	}

	if err = db.Ping(); err != nil {
		logrus.WithError(err).Fatal("database is not available")
	}

	repo := postgres.New(db)
	us := service.NewUserService(repo)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", viper.GetString("port")))
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

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("chat-backend")
	return viper.ReadInConfig()
}
