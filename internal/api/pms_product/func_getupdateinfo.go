package pms_product

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type getUpdateInfoRequest struct{}

type getUpdateInfoResponse struct{}

// GetUpdateInfo 根据商品id获取商品编辑信息
// @Summary 根据商品id获取商品编辑信息
// @Description 根据商品id获取商品编辑信息
// @Tags PmsProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getUpdateInfoRequest true "请求信息"
// @Success 200 {object} code.Success{data=getUpdateInfoResponse}
// @Failure 400 {object} code.Failure
// @Router /product/updateInfo/{id} [post]
func (h *handler) GetUpdateInfo(ctx *gin.Context) {
	api.Success(ctx, nil)
}
