package generator_handler

import (
	"go.uber.org/zap"
)

type handler struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *handler {
	return &handler{logger: logger}
}
