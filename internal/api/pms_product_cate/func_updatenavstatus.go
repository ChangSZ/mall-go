package pms_product_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type updateNavStatusRequest struct{}

type updateNavStatusResponse struct{}

// UpdateNavStatus 修改导航栏显示状态
// @Summary 修改导航栏显示状态
// @Description 修改导航栏显示状态
// @Tags PmsProductCategoryController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateNavStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateNavStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /productCategory/update/navStatus[post]
func (h *handler) UpdateNavStatus(ctx *gin.Context) {
	api.Success(ctx, nil)
}
