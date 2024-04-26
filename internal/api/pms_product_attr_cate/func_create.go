package pms_product_attr_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type createRequest struct{}

type createResponse struct{}

// Create 添加商品属性分类
// @Summary 添加商品属性分类
// @Description 添加商品属性分类
// @Tags PmsProductAttributeCategoryController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} code.Success{data=createResponse}
// @Failure 400 {object} code.Failure
// @Router /productAttribute/category/create [post]
func (h *handler) Create(ctx *gin.Context) {
	api.Success(ctx, nil)
}
