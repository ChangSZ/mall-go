package sms_flash_promotion

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 根据活动名称分页查询
// @Summary 根据活动名称分页查询
// @Description 根据活动名称分页查询
// @Tags SmsFlashPromotionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /flash/list [get]
func (h *handler) List(ctx *gin.Context) {
	api.Success(ctx, nil)
}
