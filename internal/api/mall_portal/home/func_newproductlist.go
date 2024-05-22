package home

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type newProductListRequest struct{}

type newProductListResponse struct{}

// NewProductList 分页获取新品推荐商品
// @Summary 分页获取新品推荐商品
// @Description 分页获取新品推荐商品
// @Tags HomeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body newProductListRequest true "请求信息"
// @Success 200 {object} code.Success{data=newProductListResponse}
// @Failure 400 {object} code.Failure
// @Router /home/newProductList [get]
func (h *handler) NewProductList(ctx *gin.Context) {
	api.Success(ctx, nil)
}
