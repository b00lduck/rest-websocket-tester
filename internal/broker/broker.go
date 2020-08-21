package broker

import (
	"github.com/b00lduck/rest-websocket-tester/internal/log"
)

type broker struct {
	logger       log.SugaredLogger
	messageQueue chan []byte
	broadcaster  Broadcaster
}

func NewBroker(logger log.SugaredLogger, broadcaster Broadcaster) *broker {
	return &broker{
		logger:       logger,
		messageQueue: make(chan []byte),
		broadcaster:  broadcaster,
	}
}

func (b *broker) Run() {
	for {
		b.broadcaster.Broadcast(<-b.messageQueue)
	}
}

func (b *broker) Message(message []byte) {
	b.messageQueue <- message
}
