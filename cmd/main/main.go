package main

import (
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Config initializatin error: %s", err.Error())
	}

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
