package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type getRoleRequest struct{}

type getRoleResponse struct {
	List []dto.UmsRole `json:",inline"`
}

// GetRole 获取指定用户的角色
// @Summary 获取指定用户的角色
// @Description 获取指定用户的角色
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getRoleRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.UmsRole}
// @Failure 400 {object} code.Failure
// @Router /admin/role/{adminId} [get]
func (h *handler) GetRole(ctx *gin.Context) {
	_ = new(getRoleRequest)
	res := new(getRoleResponse)
	uri := new(dto.UmsAdminIdUri)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	roleList, err := h.umsAdminService.GetRoleList(ctx, uri.AdminId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}

	res.List = roleList
	api.Success(ctx, res.List)
}
