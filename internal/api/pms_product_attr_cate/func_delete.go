package pms_product_attr_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct{}

type deleteResponse struct{}

// Delete 删除单个商品属性分类
// @Summary 删除单个商品属性分类
// @Description 删除单个商品属性分类
// @Tags PmsProductAttributeCategoryController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} code.Success{data=deleteResponse}
// @Failure 400 {object} code.Failure
// @Router /productAttribute/category/delete/{id} [post]
func (h *handler) Delete(ctx *gin.Context) {
	api.Success(ctx, nil)
}
