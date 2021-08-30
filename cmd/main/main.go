package main

import (
	"github.com/EugeneImbro/chat-backend/internal/repository"
	"github.com/EugeneImbro/chat-backend/internal/server"
	"github.com/EugeneImbro/chat-backend/internal/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Config initialization error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.name"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Database initialization error: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	listener, err := net.Listen("tcp", ":"+viper.GetString("port"))
	if err != nil {
		log.Fatalf("Listener initialization error: %s", err.Error())
	}
	s := grpc.NewServer()
	server.RegisterUserServiceServer(s, server.NewServer(services))
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("chat-backend")
	return viper.ReadInConfig()
}
