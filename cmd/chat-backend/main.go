package main

import (
	"fmt"
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/EugeneImbro/chat-backend/internal/repository"
	"github.com/EugeneImbro/chat-backend/internal/server"
	"github.com/EugeneImbro/chat-backend/internal/service"
)

//todo remove
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func main() {
	if err := initConfig(); err != nil {
		logrus.WithError(err).Fatal("config initialization error")
	}

	cfg := Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.name"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	//todo use connection-string
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.DBName, cfg.Username, cfg.Password, cfg.SSLMode))
	if err != nil {
		logrus.WithError(err).Fatal("database initialization error")
	}

	if err = db.Ping(); err != nil {
		logrus.WithError(err).Fatal("db is not available")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s"+viper.GetString("port")))
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
