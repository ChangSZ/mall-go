package admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/services/admin"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type modifyPersonalInfoRequest struct {
	Nickname string `form:"nickname"` // 昵称
	Mobile   string `form:"mobile"`   // 手机号
}

type modifyPersonalInfoResponse struct {
	Username string `json:"username"` // 用户账号
}

// ModifyPersonalInfo 修改个人信息
// @Summary 修改个人信息
// @Description 修改个人信息
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param nickname formData string true "昵称"
// @Param mobile formData string true "手机号"
// @Success 200 {object} modifyPersonalInfoResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/modify_personal_info [patch]
// @Security LoginToken
func (h *handler) ModifyPersonalInfo(ctx *gin.Context) {
	req := new(modifyPersonalInfoRequest)
	res := new(modifyPersonalInfoResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	modifyData := new(admin.ModifyData)
	modifyData.Nickname = req.Nickname
	modifyData.Mobile = req.Mobile

	if err := h.service.ModifyPersonalInfo(ctx, core.SessionUserInfo(ctx).UserID, modifyData); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminModifyPersonalInfoError, err)
		return
	}

	res.Username = core.SessionUserInfo(ctx).UserName
	api.ResponseOK(ctx, res)
}
