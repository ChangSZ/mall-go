package sysmessage

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/socket"
	"github.com/ChangSZ/mall-go/pkg/errors"

	"go.uber.org/zap"
)

var (
	err    error
	server socket.Server
)

type handler struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *handler {
	return &handler{
		logger: logger,
	}
}

func GetConn() (socket.Server, error) {
	if server != nil {
		return server, nil
	}

	return nil, errors.New("conn is nil")
}

func (h *handler) Connect() core.HandlerFunc {
	return func(ctx core.Context) {
		server, err = socket.New(h.logger, ctx.ResponseWriter(), ctx.Request(), nil)
		if err != nil {
			return
		}

		go server.OnMessage()
	}
}
