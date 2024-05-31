package ums_member

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/ums_member"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Register 会员注册
	// @Tags UmsMemberController
	// @Router /sso/register [post]
	Register(*gin.Context)

	// Login 会员登录
	// @Tags UmsMemberController
	// @Router /sso/login [post]
	Login(*gin.Context)

	// Info 获取会员信息
	// @Tags UmsMemberController
	// @Router /sso/info [get]
	Info(*gin.Context)

	// GetAuthCode 获取验证码
	// @Tags UmsMemberController
	// @Router /sso/getAuthCode [get]
	GetAuthCode(*gin.Context)

	// UpdatePassword 会员修改密码
	// @Tags UmsMemberController
	// @Router /sso/updatePassword [post]
	UpdatePassword(*gin.Context)

	// RefreshToken 刷新token
	// @Tags UmsMemberController
	// @Router /sso/refreshToken [get]
	RefreshToken(*gin.Context)
}

type handler struct {
	service ums_member.Service
}

func New() Handler {
	return &handler{
		service: ums_member.New(),
	}
}

func (h *handler) i() {}
