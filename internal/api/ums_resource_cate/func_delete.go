package ums_resource_cate

import (
	"github.com/gin-gonic/gin"
)

type deleteRequest struct{}

type deleteResponse struct{}

// Delete 根据ID删除后台资源分类
// @Summary 根据ID删除后台资源分类
// @Description 根据ID删除后台资源分类
// @Tags UmsResourceController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} deleteResponse
// @Failure 400 {object} code.Failure
// @Router /resourceCategory/delete/{id} [post]
func (h *handler) Delete(ctx *gin.Context) {

}
