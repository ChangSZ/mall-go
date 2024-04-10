package generator_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) HandlerView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "generator_handler.html", nil)
}
