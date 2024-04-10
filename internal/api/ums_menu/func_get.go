package ums_menu

import (
	"github.com/gin-gonic/gin"
)

type getRequest struct{}

type getResponse struct{}

// Get 根据ID获取菜单详情
// @Summary 根据ID获取菜单详情
// @Description 根据ID获取菜单详情
// @Tags UmsMenuController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getRequest true "请求信息"
// @Success 200 {object} getResponse
// @Failure 400 {object} code.Failure
// @Router /menu/{id} [get]
func (h *handler) Get(ctx *gin.Context) {

}
