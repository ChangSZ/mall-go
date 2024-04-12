package ums_menu

import (
	"github.com/gin-gonic/gin"
)

type createRequest struct{}

type createResponse struct{}

// Create 添加后台菜单
// @Summary 添加后台菜单
// @Description 添加后台菜单
// @Tags UmsMenuController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} code.Success{data=createResponse}
// @Failure 400 {object} code.Failure
// @Router /menu/create [post]
func (h *handler) Create(ctx *gin.Context) {

}
