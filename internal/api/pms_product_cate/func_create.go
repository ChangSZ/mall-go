package pms_product_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type createRequest struct{}

type createResponse struct{}

// Create 添加商品分类
// @Summary 添加商品分类
// @Description 添加商品分类
// @Tags PmsProductCategoryController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} code.Success{data=createResponse}
// @Failure 400 {object} code.Failure
// @Router /productCategory/create [post]
func (h *handler) Create(ctx *gin.Context) {
	api.Success(ctx, nil)
}