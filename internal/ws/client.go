package ws

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/b00lduck/rest-websocket-tester/internal/broker"
	"github.com/b00lduck/rest-websocket-tester/internal/log"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 5 * time.Second

	// Time allowed to read the pong message from the peer
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = 10 * time.Second
)

type client struct {
	logger  log.SugaredLogger
	hub     Hub
	broker  broker.Broker
	conn    ConnI
	send    chan []byte
	stop    bool
	quit    chan bool
	mutex   *sync.Mutex
	logfile string
}

func (c *client) Send() chan []byte {
	return c.send
}

func (c *client) Close() {
	close(c.send)
}

func (c *client) readPump() {
	defer func() {
		c.quit <- true
	}()

	for c.stop != true {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				c.logger.Infow("Client unexpectedly closed connection",
					"error", err)
			}
			c.quit <- true
			return
		}

		c.logger.Infow("Got message via websocket",
			"messageLen", len(msg))

		if c.logfile != "" {
			f, err := os.OpenFile(c.logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
			if err != nil {
				c.logger.Errorw("Error opening logfile", "error", err)
				return
			}
			defer f.Close()

			smg := fmt.Sprintf("---\n%s\n", string(msg))

			if _, err := f.WriteString(string(smg)); err != nil {
				c.logger.Errorw("Error writing logfile", "error", err)
				return
			}
		}
	}
}

func (c *client) writePump() {

	ticker := time.NewTicker(pingPeriod)
	defer func() {
		c.logger.Infow("Send PING")
		ticker.Stop()
		c.quit <- true
	}()

	for c.stop != true {
		select {
		case message, ok := <-c.send:
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				c.quit <- true
				return
			}
			c.sendMessage(message)

		case <-ticker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{42}); err != nil {
				return
			}
		}
	}
}

func (c *client) sendMessage(text []byte) {
	w, err := c.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		c.logger.Errorw("NextWriter error", "error", err)
		c.quit <- true
		return
	}

	_, err = w.Write(text)
	if err != nil {
		c.logger.Errorw("Write error", "error", err)
		c.quit <- true
		return
	}

	if err := w.Close(); err != nil {
		c.logger.Errorw("Close error", "error", err)
		c.quit <- true
		return
	}
}

func (c *client) Lock() {
	c.mutex.Lock()
}

func (c *client) Unlock() {
	c.mutex.Unlock()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  0,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handler(logger log.SugaredLogger, hub Hub, broker broker.Broker, w http.ResponseWriter, r *http.Request, logfile string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Errorw("error upgrading connection",
			"remoteAddr", r.RemoteAddr,
			"userAgent", r.UserAgent(),
			"error", err)
		return
	}

	client := &client{
		logger:  logger,
		hub:     hub,
		broker:  broker,
		conn:    conn,
		send:    make(chan []byte, 32),
		stop:    false,
		quit:    make(chan bool, 1),
		mutex:   &sync.Mutex{},
		logfile: logfile,
	}

	client.hub.Register(client)
	go client.writePump()
	go client.readPump()

	<-client.quit

	client.stop = true
	client.hub.Unregister(client)
	client.conn.Close()
}
