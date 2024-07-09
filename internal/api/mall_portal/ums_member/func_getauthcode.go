package ums_member

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type getAuthCodeRequest struct {
	Telephone string `form:"telephone"`
}

type getAuthCodeResponse struct {
	AuthCode string `json:",inline"`
}

// GetAuthCode 获取验证码
// @Summary 获取验证码
// @Description 获取验证码
// @Tags UmsMemberController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getAuthCodeRequest true "请求信息"
// @Success 200 {object} code.Success{data=string}
// @Failure 400 {object} code.Failure
// @Router /sso/getAuthCode [get]
func (h *handler) GetAuthCode(ctx *gin.Context) {
	req := new(getAuthCodeRequest)
	res := new(getAuthCodeResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	authCode := h.service.GenerateAuthCode(ctx, req.Telephone)
	res.AuthCode = authCode
	api.Success(ctx, res.AuthCode)
}
