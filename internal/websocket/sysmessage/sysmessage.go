package sysmessage

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/repository/socket"
)

var (
	err    error
	server socket.Server
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func GetConn() (socket.Server, error) {
	if server != nil {
		return server, nil
	}

	return nil, errors.New("conn is nil")
}

func (h *handler) Connect(ctx *gin.Context) {
	server, err = socket.New(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}

	go server.OnMessage()
}
