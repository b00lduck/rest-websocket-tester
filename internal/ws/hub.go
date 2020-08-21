package ws

import "github.com/b00lduck/rest-websocket-tester/internal/log"

type hub struct {
	logger     log.SugaredLogger
	clients    map[Client]bool
	broadcast  chan []byte
	register   chan Client
	unregister chan Client
}

func NewHub(logger log.SugaredLogger) Hub {
	return &hub{
		logger:     logger,
		broadcast:  make(chan []byte),
		register:   make(chan Client),
		unregister: make(chan Client),
		clients:    make(map[Client]bool),
	}
}

func (h *hub) Broadcast(message []byte) {
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
			h.logger.Infow("Registered new client",
				"numClients", len(h.clients))

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.Close()
				h.logger.Infow("Deregistered client",
					"numClients", len(h.clients))
			}
		case message := <-h.broadcast:
			if len(h.clients) > 0 {
				h.logger.Infow("Broadcasting message",
					"messageLength", len(message),
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
