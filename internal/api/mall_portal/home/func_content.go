package home

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
)

type contentRequest struct{}

type contentResponse struct {
	dto.HomeContentResult `json:",inline"`
}

// Content 首页内容信息展示
// @Summary 首页内容信息展示
// @Description 首页内容信息展示
// @Tags HomeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body contentRequest true "请求信息"
// @Success 200 {object} code.Success{data=contentResponse}
// @Failure 400 {object} code.Failure
// @Router /home/content [get]
func (h *handler) Content(ctx *gin.Context) {
	_ = new(contentRequest)
	res := new(contentResponse)

	data, err := h.service.Content(ctx)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.HomeContentResult = *data
	api.Success(ctx, res)
}
