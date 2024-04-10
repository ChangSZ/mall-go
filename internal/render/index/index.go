package index

import (
	"github.com/gin-gonic/gin"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) Index(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}
