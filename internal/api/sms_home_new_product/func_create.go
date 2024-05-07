package sms_home_new_product

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type createRequest struct{}

type createResponse struct{}

// Create 添加首页新品
// @Summary 添加首页新品
// @Description 添加首页新品
// @Tags SmsHomeNewProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} code.Success{data=createResponse}
// @Failure 400 {object} code.Failure
// @Router /home/newProduct [post]
func (h *handler) Create(ctx *gin.Context) {
	api.Success(ctx, nil)
}
