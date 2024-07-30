package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/torderonex/messageservice/internal/broker"
	"github.com/torderonex/messageservice/internal/config"
	"github.com/torderonex/messageservice/internal/handler"
	l "github.com/torderonex/messageservice/internal/logger"
	"github.com/torderonex/messageservice/internal/repo"
	"github.com/torderonex/messageservice/internal/service"
	"github.com/torderonex/messageservice/pkg/server"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	// init cfg
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// init logger
	logger := l.MustCreate(cfg.Env)
	slog.SetDefault(logger)
	slog.Info(fmt.Sprintf("The server starts on port %s", cfg.HTTPServer.Port))
	// init repo
	store := repo.New(cfg)
	// init kafka
	kafka := broker.New(cfg)
	// init services
	services := service.New(store, kafka)
	// handler init
	h := handler.New(services)
	// init server
	srv := server.New(cfg.HTTPServer.Port, h.InitRoutes(), cfg.HTTPServer.Timeout)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// run server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	slog.Info("Server exiting")
}
