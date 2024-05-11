package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type listMenuRequest struct{}

type listMenuResponse struct {
	List []dto.UmsMenu `json:",inline"`
}

// ListMenu 获取角色相关菜单
// @Summary 获取角色相关菜单
// @Description 获取角色相关菜单
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listMenuRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.UmsMenu}
// @Failure 400 {object} code.Failure
// @Router /role/listMenu/{roleId} [get]
func (h *handler) ListMenu(ctx *gin.Context) {
	_ = new(listMenuRequest)
	res := new(listMenuResponse)
	uri := new(dto.UmsRoleIdUri)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, err := h.service.ListMenu(ctx, uri.RoleId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}
