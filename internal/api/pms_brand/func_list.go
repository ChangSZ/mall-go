package pms_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 根据品牌名称分页获取品牌列表
// @Summary 根据品牌名称分页获取品牌列表
// @Description 根据品牌名称分页获取品牌列表
// @Tags PmsBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /brand/list [get]
func (h *handler) List(ctx *gin.Context) {
	api.Success(ctx, nil)
}
