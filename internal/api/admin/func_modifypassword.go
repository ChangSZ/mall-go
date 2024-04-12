package admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/pkg/password"
	"github.com/ChangSZ/mall-go/internal/services/admin"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type modifyPasswordRequest struct {
	OldPassword string `form:"old_password"` // 旧密码
	NewPassword string `form:"new_password"` // 新密码
}

type modifyPasswordResponse struct {
	Username string `json:"username"` // 用户账号
}

// ModifyPassword 修改密码
// @Summary 修改密码
// @Description 修改密码
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param old_password formData string true "旧密码"
// @Param new_password formData string true "新密码"
// @Success 200 {object} modifyPasswordResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/modify_password [patch]
// @Security LoginToken
func (h *handler) ModifyPassword(ctx *gin.Context) {
	req := new(modifyPasswordRequest)
	res := new(modifyPasswordResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, err)
		return
	}

	searchOneData := new(admin.SearchOneData)
	searchOneData.Id = core.SessionUserInfo(ctx).UserID
	searchOneData.Password = password.GeneratePassword(req.OldPassword)
	searchOneData.IsUsed = 1

	info, err := h.adminService.Detail(ctx, searchOneData)
	if err != nil || info == nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminModifyPasswordError, err)
		return
	}

	if err := h.adminService.ModifyPassword(ctx, core.SessionUserInfo(ctx).UserID, req.NewPassword); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminModifyPasswordError, err)
		return
	}

	res.Username = core.SessionUserInfo(ctx).UserName
	api.ResponseOK(ctx, res)
}
