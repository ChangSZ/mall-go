package ums_admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/pkg/validation"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Token     string `json:"token"`
	TokenHead string `json:"tokenHead"`
}

// Login 登录以后返回token
// @Summary 登录以后返回token
// @Description 登录以后返回token
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body loginRequest true "请求信息"
// @Success 200 {object} loginResponse
// @Failure 400 {object} code.Failure
// @Router /admin/login [post]
func (h *handler) Login() core.HandlerFunc {
	return func(c core.Context) {
		req := new(loginRequest)
		res := new(loginResponse)
		if err := c.ShouldBind(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}

		token, err := h.umsAdminService.Login(c, req.Username, req.Password)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UmsAdminAuthorizedError,
				code.Text(code.UmsAdminAuthorizedError)).WithError(err),
			)
			return
		}
		res.Token = token
		res.TokenHead = configs.Get().Jwt.TokenHead
		c.Payload(res)
	}
}

// String token = adminService.login(umsAdminLoginParam.getUsername(), umsAdminLoginParam.getPassword());
// if (token == null) {
// 	return CommonResult.validateFailed("用户名或密码错误");
// }
// Map<String, String> tokenMap = new HashMap<>();
// tokenMap.put("token", token);
// tokenMap.put("tokenHead", tokenHead);
