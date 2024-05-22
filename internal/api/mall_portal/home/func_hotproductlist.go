package home

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type hotProductListRequest struct{}

type hotProductListResponse struct{}

// HotProductList 分页获取人气推荐商品
// @Summary 分页获取人气推荐商品
// @Description 分页获取人气推荐商品
// @Tags HomeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body hotProductListRequest true "请求信息"
// @Success 200 {object} code.Success{data=hotProductListResponse}
// @Failure 400 {object} code.Failure
// @Router /home/hotProductList [get]
func (h *handler) HotProductList(ctx *gin.Context) {
	api.Success(ctx, nil)
}
