package ums_resource_cate

import (
	"github.com/gin-gonic/gin"
)

type createRequest struct{}

type createResponse struct{}

// Create 添加后台资源分类
// @Summary 添加后台资源分类
// @Description 添加后台资源分类
// @Tags UmsResourceController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} code.Success{data=createResponse}
// @Failure 400 {object} code.Failure
// @Router /resourceCategory/create [post]
func (h *handler) Create(ctx *gin.Context) {

}
