package pms_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateShowStatusRequest struct{}

type updateShowStatusResponse struct{}

// UpdateShowStatus 批量更新显示状态
// @Summary 批量更新显示状态
// @Description 批量更新显示状态
// @Tags PmsBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateShowStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateShowStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /brand/update/showStatus [post]
func (h *handler) UpdateShowStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
