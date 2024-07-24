package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/torderonex/messageservice/internal/config"
	l "github.com/torderonex/messageservice/internal/logger"
	"log"
)

func init() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	//init cfg
	cfg := config.MustLoad()
	//init logger
	logger := l.MustCreate(cfg.Env)
	logger.Info(fmt.Sprintf("The server starts on port %s", cfg.HTTPServer.Port))
	//init repo

	//init kafka

	//init services

	//init server

	//graceful shutdown
}
