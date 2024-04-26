package pms_product

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateDeleteStatusRequest struct{}

type updateDeleteStatusResponse struct{}

// UpdateDeleteStatus 批量修改删除状态
// @Summary 批量修改删除状态
// @Description 批量修改删除状态
// @Tags PmsProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateDeleteStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateDeleteStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /product/update/deleteStatus [post]
func (h *handler) UpdateDeleteStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
