package ums_admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type refreshTokenRequest struct{}

type refreshTokenResponse struct {
	Token     string `json:"token"`
	TokenHead string `json:"tokenHead"`
}

// RefreshToken 刷新token
// @Summary 刷新token
// @Description 刷新token
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body refreshTokenRequest true "请求信息"
// @Success 200 {object} refreshTokenResponse
// @Failure 400 {object} code.Failure
// @Router /admin/refreshToken [get]
func (h *handler) RefreshToken(ctx *gin.Context) {
	res := new(refreshTokenResponse)
	userInfo := core.GetUmsUserInfo(ctx)

	token, err := h.umsAdminService.RefreshToken(ctx, userInfo.Token)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.UmsAdminRefreshTokenError, err)
		return
	}
	res.Token = token
	res.TokenHead = configs.Get().Jwt.TokenHead
	api.Success(ctx, res)
}
