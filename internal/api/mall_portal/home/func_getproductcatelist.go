package home

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type getProductCateListRequest struct{}

type getProductCateListResponse struct{}

// GetProductCateList 获取首页商品分类
// @Summary 获取首页商品分类
// @Description 获取首页商品分类
// @Tags HomeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getProductCateListRequest true "请求信息"
// @Success 200 {object} code.Success{data=getProductCateListResponse}
// @Failure 400 {object} code.Failure
// @Router /home/productCateList/{parentId} [get]
func (h *handler) GetProductCateList(ctx *gin.Context) {
	api.Success(ctx, nil)
}
