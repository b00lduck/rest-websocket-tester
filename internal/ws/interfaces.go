package ws

//go:generate mockgen -source=interfaces.go -destination=mocks/ws.go -package=mock_ws

import (
	"io"
	"net"
	"time"

	"github.com/b00lduck/rest-websocket-tester/internal/dto"
)

type Client interface {
	Send() chan dto.Message
	Close()
	Lock()
	Unlock()
}

type Hub interface {
	Run()
	Unregister(c Client)
	Register(c Client)
	Broadcast(message dto.Message)
	NumClients() int
}

type ConnI interface {
	RemoteAddr() net.Addr
	Close() error
	SetReadLimit(int64)
	SetReadDeadline(time.Time) error
	SetWriteDeadline(time.Time) error
	SetPongHandler(func(string) error)
	ReadMessage() (int, []byte, error)
	WriteMessage(int, []byte) error
	NextWriter(int) (io.WriteCloser, error)
}
