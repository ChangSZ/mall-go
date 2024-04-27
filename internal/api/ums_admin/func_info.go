package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type infoRequest struct{}

type infoResponse struct {
	Username string        `json:"username"`
	Menus    []dto.UmsMenu `json:"menus"`
	Icon     string        `json:"icon"`
	Roles    []string      `json:"roles"`
}

// Info 获取当前登录用户信息
// @Summary 获取当前登录用户信息
// @Description 获取当前登录用户信息
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body infoRequest true "请求信息"
// @Success 200 {object} code.Success{data=infoResponse}
// @Failure 400 {object} code.Failure
// @Router /admin/info [get]
func (h *handler) Info(ctx *gin.Context) {
	_ = new(infoRequest)
	res := new(infoResponse)
	userInfo := core.GetUmsUserInfo(ctx)
	umsAdmin, err := h.umsAdminService.GetAdminByUsername(ctx, userInfo.UserName)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Username = umsAdmin.Username
	res.Icon = umsAdmin.Icon

	menuList, err := h.umsRoleService.GetMenuList(ctx, umsAdmin.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Menus = menuList

	roleList, err := h.umsAdminService.GetRoleList(ctx, umsAdmin.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if len(roleList) > 0 {
		res.Roles = make([]string, 0, len(roleList))
		for _, role := range roleList {
			res.Roles = append(res.Roles, role.Name)
		}
	}
	api.Success(ctx, res)
}
