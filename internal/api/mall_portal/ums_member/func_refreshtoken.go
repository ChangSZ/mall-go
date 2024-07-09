package ums_member

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type refreshTokenRequest struct{}

type refreshTokenResponse struct {
	Token     string `json:"token"`
	TokenHead string `json:"tokenHead"`
}

// RefreshToken 刷新token
// @Summary 刷新token
// @Description 刷新token
// @Tags UmsMemberController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body refreshTokenRequest true "请求信息"
// @Success 200 {object} code.Success{data=refreshTokenResponse}
// @Failure 400 {object} code.Failure
// @Router /sso/refreshToken [get]
func (h *handler) RefreshToken(ctx *gin.Context) {
	_ = new(refreshTokenRequest)
	res := new(refreshTokenResponse)
	userInfo := core.GetUmsUserInfo(ctx)

	token, err := h.service.RefreshToken(ctx, userInfo.Token)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Token = token
	res.TokenHead = configs.Get().Jwt.TokenHead
	api.Success(ctx, res)
}
