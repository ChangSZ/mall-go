package pms_product

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateNewStatusRequest struct{}

type updateNewStatusResponse struct{}

// UpdateNewStatus 批量设为新品
// @Summary 批量设为新品
// @Description 批量设为新品
// @Tags PmsProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateNewStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateNewStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /product/update/newStatus [post]
func (h *handler) UpdateNewStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
