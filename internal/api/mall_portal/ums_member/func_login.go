package ums_member

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type loginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type loginResponse struct {
	Token     string `json:"token"`
	TokenHead string `json:"tokenHead"`
}

// Login 会员登录
// @Summary 会员登录
// @Description 会员登录
// @Tags UmsMemberController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body loginRequest true "请求信息"
// @Success 200 {object} code.Success{data=loginResponse}
// @Failure 400 {object} code.Failure
// @Router /sso/login [post]
func (h *handler) Login(ctx *gin.Context) {
	req := new(loginRequest)
	res := new(loginResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	token, err := h.service.Login(ctx, req.Username, req.Password)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Token = token
	res.TokenHead = configs.Get().Jwt.TokenHead
	api.Success(ctx, res)
}
