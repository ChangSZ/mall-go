package pms_product

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updatePublishStatusRequest struct{}

type updatePublishStatusResponse struct{}

// UpdatePublishStatus 批量上下架商品
// @Summary 批量上下架商品
// @Description 批量上下架商品
// @Tags PmsProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updatePublishStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updatePublishStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /product/update/verifyStatus[post]
func (h *handler) UpdatePublishStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
