package cron

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type executeRequest struct {
	Id string `uri:"id"` // HashID
}

type executeResponse struct {
	Id int `json:"id"` // ID
}

// Execute 手动执行单条任务
// @Summary 手动执行单条任务
// @Description 手动执行单条任务
// @Tags API.cron
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/cron/exec/{id} [patch]
// @Security LoginToken
func (h *handler) Execute(ctx *gin.Context) {
	req := new(executeRequest)
	res := new(executeResponse)
	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, err)
		return
	}

	ids, err := h.hashids.HashidsDecode(req.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.HashIdsDecodeError, err)
		return
	}

	err = h.cronService.Execute(ctx, cast.ToInt32(ids[0]))
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.CronExecuteError, err)
		return
	}

	res.Id = ids[0]
	api.ResponseOK(ctx, res)
}
