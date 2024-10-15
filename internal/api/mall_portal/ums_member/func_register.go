package ums_member

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type registerRequest struct {
	Username  string `form:"username" binding:"required"`
	Password  string `form:"password" binding:"required"`
	Telephone string `form:"telephone" binding:"required"`
	AuthCode  string `form:"authCode" binding:"required"`
}

type registerResponse struct{}

// Register 会员注册
// @Summary 会员注册
// @Description 会员注册
// @Tags UmsMemberController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body registerRequest true "请求信息"
// @Success 200 {object} code.Success{data=registerResponse}
// @Failure 400 {object} code.Failure
// @Router /sso/register [post]
func (h *handler) Register(ctx *gin.Context) {
	req := new(registerRequest)
	_ = new(registerResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	err := h.service.Register(ctx, req.Username, req.Password, req.Telephone, req.AuthCode)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	api.Success(ctx, "注册成功")
}
