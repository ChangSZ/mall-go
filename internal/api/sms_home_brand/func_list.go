package sms_home_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 分页查询推荐品牌
// @Summary 分页查询推荐品牌
// @Description 分页查询推荐品牌
// @Tags SmsHomeBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /home/brand/list [get]
func (h *handler) List(ctx *gin.Context) {
	api.Success(ctx, nil)
}
