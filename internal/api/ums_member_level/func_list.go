package ums_member_level

import (
	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 查询所有会员等级
// @Summary 查询所有会员等级
// @Description 查询所有会员等级
// @Tags UmsMemberLevelController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /memberLevel/list [get]
func (h *handler) List(ctx *gin.Context) {

}
