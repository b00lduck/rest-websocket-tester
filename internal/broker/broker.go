package broker

import (
	"github.com/b00lduck/rest-websocket-tester/internal/dto"
	"github.com/b00lduck/rest-websocket-tester/internal/json"
	"github.com/b00lduck/rest-websocket-tester/internal/log"
)

type queuedMessage struct {
	message []byte
}

type broker struct {
	logger       log.SugaredLogger
	messageQueue chan queuedMessage
	broadcaster  Broadcaster
	marshaller   Marshaller
}

func NewBroker(logger log.SugaredLogger, broadcaster Broadcaster) *broker {
	return &broker{
		logger:       logger,
		messageQueue: make(chan queuedMessage),
		broadcaster:  broadcaster,
		marshaller:   json.NewMarshaller(logger),
	}
}

func (b *broker) Run() {
	for {
		message := <-b.messageQueue
		j, err := b.marshaller.Marshal(message.message)
		if err != nil {
			b.logger.Info("error marshalling json")
			continue
		}
		b.broadcaster.Broadcast(dto.Message{
			Text: j})
	}
}

func (b *broker) Message(message []byte) {
	b.messageQueue <- queuedMessage{message: message}
}
