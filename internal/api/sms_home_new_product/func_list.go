package sms_home_new_product

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 分页查询首页新品
// @Summary 分页查询首页新品
// @Description 分页查询首页新品
// @Tags SmsHomeNewProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /home/newProduct/list [get]
func (h *handler) List(ctx *gin.Context) {
	api.Success(ctx, nil)
}
