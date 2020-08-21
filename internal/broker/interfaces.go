package broker

//go:generate mockgen -source=interfaces.go -destination=mocks/broker.go -package=mock_broker

import (
	"github.com/b00lduck/rest-websocket-tester/internal/dto"
)

type Marshaller interface {
	Marshal(message []byte) ([]byte, error)
}

type Broadcaster interface {
	Broadcast(message dto.Message)
}

type Broker interface {
	Run()
	Message(message []byte)
}
