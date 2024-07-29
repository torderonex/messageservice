package broker

import (
	"github.com/torderonex/messageservice/internal/broker/kafka"
	"github.com/torderonex/messageservice/internal/config"
)

type Producer interface {
	Send(ID int) error
	Close() error
}

type Consumer interface {
	Read() (<-chan int, error)
	Close() error
}

type Broker struct {
	Producer
	Consumer
}

func New(cfg *config.Config) *Broker {
	kafkaBroker := kafka.New(cfg)
	return &Broker{
		Consumer: kafkaBroker,
		Producer: kafkaBroker,
	}
}
