package ums_resource_cate

import (
	"github.com/gin-gonic/gin"
)

type updateRequest struct{}

type updateResponse struct{}

// Update 修改后台资源分类
// @Summary 修改后台资源分类
// @Description 修改后台资源分类
// @Tags UmsResourceController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateResponse}
// @Failure 400 {object} code.Failure
// @Router /resourceCategory/update/{id} [post]
func (h *handler) Update(ctx *gin.Context) {

}
