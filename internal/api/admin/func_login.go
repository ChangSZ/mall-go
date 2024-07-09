package admin

import (
	"encoding/json"
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/password"
	"github.com/ChangSZ/mall-go/internal/proposal"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/internal/services/admin"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type loginRequest struct {
	Username string `form:"username"` // 用户名
	Password string `form:"password"` // 密码
}

type loginResponse struct {
	Token string `json:"token"` // 用户身份标识
}

// Login 管理员登录
// @Summary 管理员登录
// @Description 管理员登录
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "MD5后的密码"
// @Success 200 {object} loginResponse
// @Failure 400 {object} code.Failure
// @Router /api/login [post]
// @Security LoginToken
func (h *handler) Login(ctx *gin.Context) {
	req := new(loginRequest)
	res := new(loginResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	searchOneData := new(admin.SearchOneData)
	searchOneData.Username = req.Username
	searchOneData.Password = password.GeneratePassword(req.Password)
	searchOneData.IsUsed = 1

	info, err := h.service.Detail(ctx, searchOneData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminLoginError, err)
		return
	}

	if info == nil {
		log.WithTrace(ctx).Error("未查询出符合条件的用户")
		api.Response(ctx, http.StatusBadRequest, code.AdminLoginError, "未查询出符合条件的用户")
		return
	}

	token := password.GenerateLoginToken(info.Id)

	// 用户信息
	sessionUserInfo := &proposal.SessionUserInfo{
		UserID:   info.Id,
		UserName: info.Username,
	}

	// 将用户信息记录到 Redis 中
	err = redis.Cache().Set(ctx, configs.RedisKeyPrefixLoginUser+token, string(sessionUserInfo.Marshal()), configs.LoginSessionTTL)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminLoginError, err)
		return
	}

	searchMenuData := new(admin.SearchMyMenuData)
	searchMenuData.AdminId = info.Id
	menu, err := h.service.MyMenu(ctx, searchMenuData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminLoginError, err)
		return
	}

	// 菜单栏信息
	menuJsonInfo, _ := json.Marshal(menu)

	// 将菜单栏信息记录到 Redis 中
	err = redis.Cache().Set(ctx, configs.RedisKeyPrefixLoginUser+token+":menu", string(menuJsonInfo), configs.LoginSessionTTL)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminLoginError, err)
		return
	}

	searchActionData := new(admin.SearchMyActionData)
	searchActionData.AdminId = info.Id
	action, err := h.service.MyAction(ctx, searchActionData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminLoginError, err)
		return
	}

	// 可访问接口信息
	actionJsonInfo, _ := json.Marshal(action)

	// 将可访问接口信息记录到 Redis 中
	err = redis.Cache().Set(ctx, configs.RedisKeyPrefixLoginUser+token+":action", string(actionJsonInfo), configs.LoginSessionTTL)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminLoginError, err)
		return
	}

	res.Token = token
	api.ResponseOK(ctx, res)
}
