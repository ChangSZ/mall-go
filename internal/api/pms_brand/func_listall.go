package pms_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type listAllRequest struct{}

type listAllResponse struct{}

// ListAll 获取全部品牌列表
// @Summary 获取全部品牌列表
// @Description 获取全部品牌列表
// @Tags PmsBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listAllRequest true "请求信息"
// @Success 200 {object} code.Success{data=listAllResponse}
// @Failure 400 {object} code.Failure
// @Router /brand/listAll [get]
func (h *handler) ListAll(ctx *gin.Context) {
	api.Success(ctx, nil)
}
