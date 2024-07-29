package repo

import (
	"context"
	"github.com/torderonex/messageservice/internal/config"
	"github.com/torderonex/messageservice/internal/repo/postgres"
)

func New(config *config.Config) *Repository {
	pg := postgres.New(config.Postgres)
	return &Repository{
		Messages: postgres.NewMessageRepo(pg),
	}
}

type Repository struct {
	Messages
}

type Messages interface {
	SaveMessage(c context.Context, content string) (int, error)
}
