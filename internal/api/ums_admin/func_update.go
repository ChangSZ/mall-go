package ums_admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	UmsAdmin `json:",inline"`
}

type updateResponse struct {
	Count int64
}

// Update 修改指定用户信息
// @Summary 修改指定用户信息
// @Description 修改指定用户信息
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateResponse.Count}
// @Failure 400 {object} code.Failure
// @Router /admin/update/{id} [post]
func (h *handler) Update(ctx *gin.Context) {
	req := new(updateRequest)
	res := new(updateResponse)
	uri := new(UmsAdminUri)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	data := &ums_admin.UmsAdmin{
		Id:         req.ID,
		Username:   req.Username,
		Password:   req.Password,
		Icon:       req.Icon,
		Email:      req.Email,
		NickName:   req.NickName,
		Note:       req.Note,
		CreateTime: req.CreateTime,
		LoginTime:  req.LoginTime,
		Status:     req.Status,
	}
	cnt, err := h.umsAdminService.Update(ctx, uri.Id, data)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.UmsAdminUpdateError, err)
		return
	}
	if cnt == 0 {
		api.Response(ctx, http.StatusBadRequest, code.UmsAdminUpdateError, "更新数量为0")
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
