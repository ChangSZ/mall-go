package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateRoleRequest struct {
	AdminId int64   `json:"adminId"`
	RoleIds []int64 `json:"roleIds"`
}

type updateRoleResponse struct {
	Count int64 `json:"count"`
}

// UpdateRole 给用户分配角色
// @Summary 给用户分配角色
// @Description 给用户分配角色
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRoleRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /admin/role/update [post]
func (h *handler) UpdateRole(ctx *gin.Context) {
	req := new(updateRoleRequest)
	res := new(updateRoleResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}
	count, err := h.service.UpdateRole(ctx, req.AdminId, req.RoleIds)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if count <= 0 {
		api.Failed(ctx, "无变更")
		return
	}
	res.Count = count
	api.Success(ctx, res.Count)
}
