package home

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type contentRequest struct{}

type contentResponse struct{}

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
	api.Success(ctx, nil)
}
