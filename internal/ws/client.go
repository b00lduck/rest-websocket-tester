package ws

import (
	"net/http"
	"sync"
	"time"

	"github.com/b00lduck/rest-websocket-tester/internal/broker"
	"github.com/b00lduck/rest-websocket-tester/internal/dto"
	"github.com/gorilla/websocket"
	"github.com/tarent/logrus"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the message form the peer
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

type client struct {
	hub    Hub
	broker broker.Broker
	conn   ConnI
	send   chan dto.Message
	stop   bool
	quit   chan bool
	mutex  *sync.Mutex
}

func (c *client) Send() chan dto.Message {
	return c.send
}

func (c *client) Close() {
	close(c.send)
}

func (c *client) readPump() {
	defer func() {
		c.quit <- true
	}()

	//c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for c.stop != true {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				logrus.WithField("err", err).Info("Client unexpectedly closed connection")
			}
			c.quit <- true
			return
		}

		logrus.Info("Got message:", msg)
	}

}

func (c *client) writePump() {

	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.quit <- true
	}()

	for c.stop != true {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				c.quit <- true
				return
			}
			c.sendMessage(message.Text)

		case <-ticker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *client) sendMessage(text []byte) {
	w, err := c.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		logrus.WithField("err", err).Info("NextWriter error")
		c.quit <- true
		return
	}

	_, err = w.Write(text)
	if err != nil {
		logrus.WithField("err", err).Info("Write error")
		c.quit <- true
		return
	}

	if err := w.Close(); err != nil {
		logrus.WithField("err", err).Info("Close error")
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

func Handler(hub Hub, broker broker.Broker, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.WithField("err", err).
			WithField("remoteAddr", r.RemoteAddr).
			WithField("userAgent", r.UserAgent()).
			Error("error upgrading connection")
		return
	}

	client := &client{
		hub:    hub,
		broker: broker,
		conn:   conn,
		send:   make(chan dto.Message, 32),
		stop:   false,
		quit:   make(chan bool, 1),
		mutex:  &sync.Mutex{}}

	client.hub.Register(client)
	go client.writePump()
	go client.readPump()

	<-client.quit

	client.stop = true
	client.hub.Unregister(client)
	client.conn.Close()
}
