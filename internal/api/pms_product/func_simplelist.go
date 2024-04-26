package pms_product

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type simpleListRequest struct{}

type simpleListResponse struct{}

// SimpleList 根据商品名称或货号模糊查询
// @Summary 根据商品名称或货号模糊查询
// @Description 根据商品名称或货号模糊查询
// @Tags PmsProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body simpleListRequest true "请求信息"
// @Success 200 {object} code.Success{data=simpleListResponse}
// @Failure 400 {object} code.Failure
// @Router /product/simpleList [get]
func (h *handler) SimpleList(ctx *gin.Context) {
	api.Success(ctx, nil)
}
