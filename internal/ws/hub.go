package ws

import (
	"github.com/b00lduck/rest-websocket-tester/internal/dto"
	"go.uber.org/zap"
)

type hub struct {
	logger     *zap.SugaredLogger
	clients    map[Client]bool
	broadcast  chan dto.Message
	register   chan Client
	unregister chan Client
}

func NewHub(logger *zap.SugaredLogger) Hub {
	return &hub{
		logger:     logger,
		broadcast:  make(chan dto.Message),
		register:   make(chan Client),
		unregister: make(chan Client),
		clients:    make(map[Client]bool),
	}
}

func (h *hub) Broadcast(message dto.Message) {
	h.broadcast <- message
}

func (h *hub) Unregister(c Client) {
	h.unregister <- c
}

func (h *hub) Register(c Client) {
	h.register <- c
}

func (h *hub) NumClients() int {
	return len(h.clients)
}

func (h *hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.logger.Info("Registered new client",
				"numClients", len(h.clients))

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.Close()
				h.logger.Info("Deregistered client",
					"numClients", len(h.clients))
			}
		case message := <-h.broadcast:
			if len(h.clients) > 0 {
				h.logger.Info("Broadcasting message",
					"messageLength", len(message.Text),
					"numClients", len(h.clients))
				for client := range h.clients {
					func() {
						client.Lock()
						defer client.Unlock()
						select {
						case client.Send() <- message:
						default:
							client.Close()
							delete(h.clients, client)
						}
					}()
				}
			}
		}
	}
}
