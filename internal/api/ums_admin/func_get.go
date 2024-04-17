package ums_admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/services/ums_user"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type getRequest struct {
	Id int64 `uri:"id" binding:"required"` // 用户ID
}

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
	req := new(getRequest)
	res := new(getResponse)

	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, err)
		return
	}

	admin, err := ums_user.New().GetItem(ctx, req.Id)
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
