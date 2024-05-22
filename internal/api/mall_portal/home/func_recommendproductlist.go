package home

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type recommendProductListRequest struct{}

type recommendProductListResponse struct{}

// RecommendProductList 分页获取推荐商品
// @Summary 分页获取推荐商品
// @Description 分页获取推荐商品
// @Tags HomeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body recommendProductListRequest true "请求信息"
// @Success 200 {object} code.Success{data=recommendProductListResponse}
// @Failure 400 {object} code.Failure
// @Router /home/recommendProductList [get]
func (h *handler) RecommendProductList(ctx *gin.Context) {
	api.Success(ctx, nil)
}
