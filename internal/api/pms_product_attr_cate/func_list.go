package pms_product_attr_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 分页获取所有商品属性分类
// @Summary 分页获取所有商品属性分类
// @Description 分页获取所有商品属性分类
// @Tags PmsProductAttributeCategoryController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /productAttribute/category/list [get]
func (h *handler) List(ctx *gin.Context) {
	api.Success(ctx, nil)
}
