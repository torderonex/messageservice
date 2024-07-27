package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/torderonex/messageservice/internal/config"
	l "github.com/torderonex/messageservice/internal/logger"
	"github.com/torderonex/messageservice/internal/repo"
	"log"
	"strconv"
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
	store := repo.New(cfg)
	id, _ := store.Messages.SaveMessage(context.TODO(), "Hello, world!")
	logger.Info(strconv.Itoa(id))
	//init kafka

	//init services

	//init server

	//graceful shutdown
}
