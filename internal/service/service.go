package service

import (
	"github.com/torderonex/messageservice/internal/repo"
)

type Service struct {
}

func New(repo *repo.Repository) *Service {

	return &Service{}
}
