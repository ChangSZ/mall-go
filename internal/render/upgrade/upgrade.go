package upgrade

import (
	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"go.uber.org/zap"
)

type handler struct {
	db     mysql.Repo
	logger *zap.Logger
}

func New(logger *zap.Logger, db mysql.Repo) *handler {
	return &handler{
		logger: logger,
		db:     db,
	}
}
