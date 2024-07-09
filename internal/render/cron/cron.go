package cron

import (
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "cron_task_add.html", nil)
}

func (h *handler) Edit(ctx *gin.Context) {
	type editRequest struct {
		Id string `uri:"id"` // 主键ID
	}

	type editResponse struct {
		HashID string `json:"hash_id"` // hashID
	}

	req := new(editRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	obj := new(editResponse)
	obj.HashID = req.Id
	ctx.HTML(http.StatusOK, "cron_task_edit.html", obj)
}

func (h *handler) List(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "cron_task_list.html", nil)
}
