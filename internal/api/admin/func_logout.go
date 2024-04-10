package admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/gin-gonic/gin"
)

type logoutResponse struct {
	Username string `json:"username"` // 用户账号
}

// Logout 管理员登出
// @Summary 管理员登出
// @Description 管理员登出
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} logoutResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/logout [post]
// @Security LoginToken
func (h *handler) Logout(ctx *gin.Context) {
	res := new(logoutResponse)
	res.Username = ctx.SessionUserInfo().UserName

	if !redis.Cache().Del(ctx, configs.RedisKeyPrefixLoginUser+ctx.GetHeader(configs.HeaderLoginToken)) {
		log.WithTrace(ctx).Error("cache del err")
		api.Response(ctx, http.StatusBadRequest, code.AdminLogOutError, "cache del err")
		return
	}

	api.ResponseOK(ctx, res)
}
