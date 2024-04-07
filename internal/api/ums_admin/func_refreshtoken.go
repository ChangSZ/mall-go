package ums_admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/code"
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
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body refreshTokenRequest true "请求信息"
// @Success 200 {object} refreshTokenResponse
// @Failure 400 {object} code.Failure
// @Router /admin/refreshToken [get]
func (h *handler) RefreshToken() core.HandlerFunc {
	return func(c core.Context) {
		res := new(refreshTokenResponse)
		token := c.Request().Header.Get(configs.Get().Jwt.TokenHeader)

		token, err := h.umsAdminService.RefreshToken(c, token)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UmsAdminRefreshTokenError,
				code.Text(code.UmsAdminRefreshTokenError)).WithError(err),
			)
			return
		}
		res.Token = token
		res.TokenHead = configs.Get().Jwt.TokenHead
		c.Payload(res)
	}
}
