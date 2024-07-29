package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/torderonex/messageservice/internal/broker"
	"github.com/torderonex/messageservice/internal/config"
	l "github.com/torderonex/messageservice/internal/logger"
	"github.com/torderonex/messageservice/internal/repo"
	"log"
	"log/slog"
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
	fmt.Println(cfg)
	//init logger
	logger := l.MustCreate(cfg.Env)
	slog.SetDefault(logger)
	slog.Info(fmt.Sprintf("The server starts on port %s", cfg.HTTPServer.Port))
	//init repo
	store := repo.New(cfg)
	id, _ := store.Messages.SaveMessage(context.TODO(), "Hello, world!")
	logger.Info(strconv.Itoa(id))
	//init kafka
	kafka := broker.New(cfg)
	err := kafka.Producer.Send(228)
	if err != nil {
		log.Fatal(err)
	}
	//init services

	//init server

	//graceful shutdown
}
