package pms_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateFactoryStatusRequest struct{}

type updateFactoryStatusResponse struct{}

// UpdateFactoryStatus 批量更新厂家制造商状态
// @Summary 批量更新厂家制造商状态
// @Description 批量更新厂家制造商状态
// @Tags PmsBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateFactoryStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateFactoryStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /brand/update/factoryStatus [post]
func (h *handler) UpdateFactoryStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
