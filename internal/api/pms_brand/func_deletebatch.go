package pms_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type deleteBatchRequest struct{}

type deleteBatchResponse struct{}

// DeleteBatch 批量删除品牌
// @Summary 批量删除品牌
// @Description 批量删除品牌
// @Tags PmsBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteBatchRequest true "请求信息"
// @Success 200 {object} code.Success{data=deleteBatchResponse}
// @Failure 400 {object} code.Failure
// @Router /brand/delete/batch [post]
func (h *handler) DeleteBatch(ctx *gin.Context) {
	api.Success(ctx, nil)
}
