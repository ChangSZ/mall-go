package ums_admin

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type registerRequest struct {
	dto.UmsAdminParam `json:",inline"`
}

type registerResponse struct {
	dto.UmsAdmin `json:",inline"`
}

// Register 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body registerRequest true "请求信息"
// @Success 200 {object} code.Success{data=registerResponse}
// @Failure 400 {object} code.Failure
// @Router /admin/register [post]
func (h *handler) Register(ctx *gin.Context) {
	req := new(registerRequest)
	res := new(registerResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	umsAdmin, err := h.service.Register(ctx, req.UmsAdminParam)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.UmsAdmin = *umsAdmin
	api.Success(ctx, res)
}
