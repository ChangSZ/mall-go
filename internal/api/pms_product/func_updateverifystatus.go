package pms_product

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateVerifyStatusRequest struct{}

type updateVerifyStatusResponse struct{}

// UpdateVerifyStatus 批量修改审核状态
// @Summary 批量修改审核状态
// @Description 批量修改审核状态
// @Tags PmsProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateVerifyStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateVerifyStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /product/update/verifyStatus[post]
func (h *handler) UpdateVerifyStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
