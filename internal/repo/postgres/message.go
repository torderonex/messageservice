package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	messagesTable = "messages"
)

type MessageRepo struct {
	db *sqlx.DB
}

func NewMessageRepo(db *sqlx.DB) *MessageRepo {
	return &MessageRepo{
		db,
	}
}

func (m *MessageRepo) SaveMessage(ctx context.Context, content string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (content) VALUES ($1) RETURNING id", messagesTable)
	row := m.db.QueryRowContext(ctx, query, content)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
