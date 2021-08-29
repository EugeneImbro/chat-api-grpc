package main

import (
	"github.com/EugeneImbro/chat-backend/internal/repository"
	"github.com/EugeneImbro/chat-backend/internal/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Config initialization error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:      viper.GetString("db.port"),
		DBName:    viper.GetString("db.name"),
		Username:  viper.GetString("db.username"),
		Password:  viper.GetString("db.password"),
		SSLMode:   viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Database initialization error: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	//todo grpc
	_ = services

	server := &http.Server{Addr: ":" + viper.GetString("port")}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Start server error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("chat-backend")
	return viper.ReadInConfig()
}
