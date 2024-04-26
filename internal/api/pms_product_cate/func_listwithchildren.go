package pms_product_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type listWithChildrenRequest struct{}

type listWithChildrenResponse struct{}

// ListWithChildren 查询所有一级分类及子分类
// @Summary 查询所有一级分类及子分类
// @Description 查询所有一级分类及子分类
// @Tags PmsProductCategoryController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listWithChildrenRequest true "请求信息"
// @Success 200 {object} code.Success{data=listWithChildrenResponse}
// @Failure 400 {object} code.Failure
// @Router /productCategory/list/withChildren [get]
func (h *handler) ListWithChildren(ctx *gin.Context) {
	api.Success(ctx, nil)
}
