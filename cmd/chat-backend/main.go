package main

import (
	"net"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/EugeneImbro/chat-backend/internal/repository"
	"github.com/EugeneImbro/chat-backend/internal/repository/postgres"
	"github.com/EugeneImbro/chat-backend/internal/server"
	"github.com/EugeneImbro/chat-backend/internal/service"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.WithError(err).Fatal("config initialization error")
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.name"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.WithError(err).Fatal("database initialization error")
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	listener, err := net.Listen("tcp", ":"+viper.GetString("port"))
	if err != nil {
		logrus.WithError(err).Fatal("listener initialization error")
	}
	s := grpc.NewServer()
	server.RegisterUserServiceServer(s, server.NewServer(services))
	if err := s.Serve(listener); err != nil {
		logrus.WithError(err).Fatal("failed to serve")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("chat-backend")
	return viper.ReadInConfig()
}
