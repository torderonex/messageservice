package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/torderonex/messageservice/internal/config"
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
	fmt.Println(cfg)
	//init logger

	//init repo

	//init kafka

	//init services

	//init server

	//graceful shutdown
}
