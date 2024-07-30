package service

import (
	"context"
	"github.com/torderonex/messageservice/internal/broker"
	"github.com/torderonex/messageservice/internal/repo"
	"log/slog"
)

type MessageService struct {
	repo   repo.Messages
	broker *broker.Broker
}

func newMessageService(repo repo.Messages, broker *broker.Broker) *MessageService {
	return &MessageService{
		repo:   repo,
		broker: broker,
	}
}

// SendMessage sends a message to the broker
func (m *MessageService) SendMessage(ctx context.Context, content string) (int, error) {
	slog.Info("Saving message", "content", content)

	messageID, err := m.repo.SaveMessage(ctx, content)
	if err != nil {
		slog.Error("Failed to save message", "error", err)
		return 0, err
	}

	slog.Info("Message saved", "messageID", messageID)

	err = m.broker.Send(messageID)
	if err != nil {
		slog.Error("Failed to send message to broker", "messageID", messageID, "error", err)
		return 0, err
	}

	slog.Info("Message sent to broker", "messageID", messageID)
	return messageID, nil
}

// TODO: fix infinite waiting for a message if query is empty
// ProcessMessages reads messages from the broker, processes them, and marks them as processed
func (m *MessageService) ProcessMessages(ctx context.Context) error {
	slog.Info("Starting to process messages from broker")

	messageChan, err := m.broker.Read()
	if err != nil {
		slog.Error("Failed to read messages from broker", "error", err)
		return err
	}

	for messageID := range messageChan {
		slog.Info("Processing message", "messageID", messageID)

		err = m.markMessageAsProcessed(ctx, messageID)
		if err != nil {
			slog.Error("Failed to mark message as processed", "messageID", messageID, "error", err)
			return err
		}

		slog.Info("Message marked as processed", "messageID", messageID)
	}

	slog.Info("Finished processing messages from broker")
	return nil
}

// markMessageAsProcessed marks a message as processed in the repository
func (m *MessageService) markMessageAsProcessed(ctx context.Context, messageID int) error {
	slog.Info("Marking message as processed", "messageID", messageID)

	err := m.repo.ProcessMessage(ctx, messageID)
	if err != nil {
		slog.Error("Failed to mark message as processed in repository", "messageID", messageID, "error", err)
		return err
	}

	slog.Info("Message marked as processed in repository", "messageID", messageID)
	return nil
}

// GetProcessedMessagesStats returns statistics about processed messages
func (m *MessageService) GetProcessedMessagesStats(ctx context.Context) (map[string]int, error) {
	slog.Info("Getting processed messages statistics")

	stats, err := m.repo.GetProcessedMessagesStats(ctx)
	if err != nil {
		slog.Error("Failed to get processed messages statistics", "error", err)
		return nil, err
	}

	slog.Info("Retrieved processed messages statistics", "stats", stats)
	return map[string]int{}, nil
}
