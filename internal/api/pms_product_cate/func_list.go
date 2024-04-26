package pms_product_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 分页查询商品分类
// @Summary 分页查询商品分类
// @Description 分页查询商品分类
// @Tags PmsProductCategoryController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /productCategory/{parentId} [get]
func (h *handler) List(ctx *gin.Context) {
	api.Success(ctx, nil)
}
