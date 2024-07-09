package socket

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"github.com/ChangSZ/mall-go/pkg/errors"
)

var _ Server = (*server)(nil)

type server struct {
	socket *websocket.Conn
}

type Server interface {
	i()

	// OnMessage 接收消息
	OnMessage()

	// OnSend 发送消息
	OnSend(message []byte) error

	// OnClose 关闭
	OnClose()
}

var upGrader = websocket.Upgrader{
	HandshakeTimeout: 5 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func New(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (Server, error) {
	ws, err := upGrader.Upgrade(w, r, responseHeader)
	if err != nil {
		return nil, errors.Wrap(err, "ws error")
	}

	return &server{
		socket: ws,
	}, nil
}

func (s *server) i() {}
