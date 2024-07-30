package service

import (
	"context"
	"github.com/torderonex/messageservice/internal/broker"
	"github.com/torderonex/messageservice/internal/repo"
)

type Service struct {
	Message
}

func New(repo *repo.Repository, broker *broker.Broker) *Service {
	return &Service{
		Message: newMessageService(repo.Messages, broker),
	}
}

type Message interface {
	SendMessage(ctx context.Context, content string) (int, error)
	ProcessMessages(ctx context.Context) error
}
