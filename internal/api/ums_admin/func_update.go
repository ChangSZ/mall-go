package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/api"
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
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
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
		api.Failed(ctx, err.Error())
		return
	}
	if cnt == 0 {
		api.Failed(ctx, "更新数量为0")
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
