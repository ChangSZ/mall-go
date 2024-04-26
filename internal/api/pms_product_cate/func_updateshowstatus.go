package pms_product_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateShowStatusRequest struct{}

type updateShowStatusResponse struct{}

// UpdateShowStatus 修改显示状态
// @Summary 修改显示状态
// @Description 修改显示状态
// @Tags PmsProductCategoryController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateShowStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateShowStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /productCategory/update/showStatus[post]
func (h *handler) UpdateShowStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
