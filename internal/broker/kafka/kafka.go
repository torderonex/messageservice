package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/torderonex/messageservice/internal/config"
	"log"
)

type Broker struct {
	cfg      *config.Config
	producer sarama.SyncProducer
	consumer sarama.Consumer
}

type messageKafka struct {
	ID int
}

func New(cfg *config.Config) *Broker {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	brokers := []string{cfg.Kafka.Host}

	producer, err := sarama.NewSyncProducer(brokers, kafkaConfig)
	if err != nil {
		log.Fatal(err)
	}

	consumer, err := sarama.NewConsumer(brokers, kafkaConfig)
	if err != nil {
		log.Fatal(err)
	}

	return &Broker{cfg: cfg, producer: producer, consumer: consumer}
}

func (con *Broker) Send(ID int) error {
	message := messageKafka{
		ID: ID,
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: con.cfg.Kafka.Topic,
		Value: sarama.ByteEncoder(bytes),
	}

	_, _, err = con.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

func (con *Broker) Read() (<-chan int, error) {
	consumer, err := con.consumer.ConsumePartition(con.cfg.Kafka.Topic, 0, sarama.OffsetOldest)
	if err != nil {
		return nil, err
	}

	ch := make(chan int)
	go func() {
		defer close(ch)
		for msg := range consumer.Messages() {
			var notification messageKafka
			err := json.Unmarshal(msg.Value, &notification)
			if err != nil {
				log.Println("Error unmarshalling message:", err)
				continue
			}
			ch <- notification.ID
		}
	}()

	return ch, nil
}

func (con *Broker) Close() error {
	if err := con.producer.Close(); err != nil {
		return err
	}
	if err := con.consumer.Close(); err != nil {
		return err
	}
	return nil
}
