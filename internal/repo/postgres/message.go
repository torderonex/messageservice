package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/torderonex/messageservice/internal/entity"
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

func (m *MessageRepo) ProcessMessage(ctx context.Context, id int) error {
	query := fmt.Sprintf("UPDATE %s SET is_processed = 't' WHERE id = $1", messagesTable)
	_, err := m.db.ExecContext(ctx, query, id)
	return err
}

func (m *MessageRepo) GetProcessedMessagesStats(ctx context.Context) ([]entity.Message, error) {
	var res []entity.Message
	query := fmt.Sprintf("SELECT * FROM %s WHERE is_processed = 't'", messagesTable)
	err := m.db.SelectContext(ctx, &res, query)
	return res, err
}
