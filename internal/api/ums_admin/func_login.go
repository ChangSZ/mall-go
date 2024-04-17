package ums_admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
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
// @Success 200 {object} code.Success{data=loginResponse}
// @Failure 400 {object} code.Failure
// @Router /admin/login [post]
func (h *handler) Login(ctx *gin.Context) {
	req := new(loginRequest)
	res := new(loginResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	token, err := h.umsAdminService.Login(ctx, req.Username, req.Password)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.UmsAdminAuthorizedError, err)
		return
	}
	res.Token = token
	res.TokenHead = configs.Get().Jwt.TokenHead
	api.Success(ctx, res)
}
