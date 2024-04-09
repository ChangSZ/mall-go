package socket

import (
	"net/http"
	"time"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/pkg/errors"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var _ Server = (*server)(nil)

type server struct {
	logger *zap.Logger
	db     mysql.Repo
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

func New(logger *zap.Logger, db mysql.Repo, w http.ResponseWriter, r *http.Request, responseHeader http.Header) (Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	if db == nil {
		return nil, errors.New("db required")
	}

	ws, err := upGrader.Upgrade(w, r, responseHeader)
	if err != nil {
		return nil, errors.Wrap(err, "ws error")
	}

	return &server{
		logger: logger,
		db:     db,
		socket: ws,
	}, nil
}

func (s *server) i() {}
