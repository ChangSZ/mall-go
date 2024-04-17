package ums_admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type getRequest struct{}

type getResponse struct {
	UmsAdmin `json:",inline"`
}

// Get 获取指定用户信息
// @Summary 获取指定用户信息
// @Description 获取指定用户信息
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getRequest true "请求信息"
// @Success 200 {object} code.Success{data=getResponse}
// @Failure 400 {object} code.Failure
// @Router /admin/{id} [get]
func (h *handler) Get(ctx *gin.Context) {
	_ = new(getRequest)
	res := new(getResponse)
	uri := new(UmsAdminUri)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	admin, err := h.umsAdminService.GetItem(ctx, uri.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.UmsAdminGetItemError, err)
		return
	}
	res.UmsAdmin = UmsAdmin{
		ID:         admin.Id,
		Username:   admin.Username,
		Password:   admin.Password,
		Icon:       admin.Icon,
		Email:      admin.Email,
		NickName:   admin.NickName,
		Note:       admin.Note,
		CreateTime: admin.CreateTime,
		LoginTime:  admin.LoginTime,
		Status:     admin.Status,
	}
	api.Success(ctx, res)
}
