package pms_product_attr

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 根据分类查询属性列表或参数列表
// @Summary 根据分类查询属性列表或参数列表
// @Description 根据分类查询属性列表或参数列表
// @Tags PmsProductAttributeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /productAttribute/list/{cid} [get]
func (h *handler) List(ctx *gin.Context) {
	api.Success(ctx, nil)
}
