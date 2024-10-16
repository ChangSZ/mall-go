package ums_admin

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type listRequest struct {
	Keyword  string `form:"keyword" binding:"omitempty"`
	PageSize int    `form:"pageSize,default=5" binding:"omitempty"`
	PageNum  int    `form:"pageNum,default=1" binding:"omitempty"`
}

type listResponse struct {
	*pagehelper.ListData[dto.UmsAdmin] `json:",inline"`
}

// List 根据用户名或姓名分页获取用户列表
// @Summary 根据用户名或姓名分页获取用户列表
// @Description 根据用户名或姓名分页获取用户列表
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request formData listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /admin/list [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	list, err := h.service.List(ctx, req.Keyword, req.PageSize, req.PageNum)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.ListData = list
	api.Success(ctx, res)
}
