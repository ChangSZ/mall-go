package ums_member

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type updatePasswordRequest struct {
	Telephone string `form:"telephone"`
	Password  string `form:"password"`
	AuthCode  string `form:"authCode"`
}

type updatePasswordResponse struct{}

// UpdatePassword 会员修改密码
// @Summary 会员修改密码
// @Description 会员修改密码
// @Tags UmsMemberController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updatePasswordRequest true "请求信息"
// @Success 200 {object} code.Success{data=string}
// @Failure 400 {object} code.Failure
// @Router /sso/updatePassword [post]
func (h *handler) UpdatePassword(ctx *gin.Context) {
	req := new(updatePasswordRequest)
	_ = new(updatePasswordResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	err := h.service.UpdatePassword(ctx, req.Telephone, req.Password, req.AuthCode)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	api.Success(ctx, "密码修改成功")
}
