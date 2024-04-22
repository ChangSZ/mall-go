package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updatePasswordRequest struct {
	Username    string `json:"username" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type updatePasswordResponse struct {
	Status int64 `json:"status"`
}

// UpdatePassword 修改指定用户密码
// @Summary 修改指定用户密码
// @Description 修改指定用户密码
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updatePasswordRequest true "请求信息"
// @Success 200 {object} code.Success{data=updatePasswordResponse}
// @Failure 400 {object} code.Failure
// @Router /admin/updatePassword [post]
func (h *handler) UpdatePassword(ctx *gin.Context) {
	req := new(updatePasswordRequest)
	res := new(updatePasswordResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}
	status, err := h.umsAdminService.UpdatePassword(ctx, req.Username, req.OldPassword, req.NewPassword)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if status <= 0 {
		api.Failed(ctx, "密码不存在变更")
		return
	}
	res.Status = status
	api.Success(ctx, res.Status)
}
