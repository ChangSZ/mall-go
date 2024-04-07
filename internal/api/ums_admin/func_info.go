package ums_admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_menu"
	"github.com/ChangSZ/mall-go/internal/services/ums_user"
)

type infoRequest struct{}

type infoResponse struct {
	Username string             `json:"username"`
	Menus    []ums_menu.UmsMenu `json:"menus"`
	Icon     string             `json:"icon"`
	Roles    []string           `json:"roles"`
}

// Info 获取当前登录用户信息
// @Summary 获取当前登录用户信息
// @Description 获取当前登录用户信息
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body infoRequest true "请求信息"
// @Success 200 {object} infoResponse
// @Failure 400 {object} code.Failure
// @Router /admin/info [get]
func (h *handler) Info() core.HandlerFunc {
	return func(c core.Context) {
		res := new(infoResponse)
		userInfo := c.GetUmsUserInfo()
		umsAdmin, err := ums_user.DefalutService.GetAdminByUsername(c, userInfo.UserName)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UmsAdminGetUsernameError,
				code.Text(code.UmsAdminGetUsernameError)).WithError(err),
			)
			return
		}
		res.Username = umsAdmin.Username
		res.Icon = umsAdmin.Icon

		menuList, err := h.umsRoleService.GetMenuList(c, umsAdmin.Id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UmsAdminGetMenuListError,
				code.Text(code.UmsAdminGetMenuListError)).WithError(err),
			)
			return
		}
		res.Menus = menuList

		roleList, err := h.umsAdminService.GetRoleList(c, umsAdmin.Id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UmsAdminGetRoleListError,
				code.Text(code.UmsAdminGetRoleListError)).WithError(err),
			)
			return
		}
		if len(roleList) > 0 {
			res.Roles = make([]string, 0, len(roleList))
			for _, role := range roleList {
				res.Roles = append(res.Roles, role.Name)
			}
		}
		c.Payload(res)
	}
}
