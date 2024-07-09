package ums_member_level

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type listRequest struct {
	DefaultStatus int32 `form:"defaultStatus"`
}

type listResponse struct {
	List []dto.UmsMemberLevel `json:",inline"`
}

// List 查询所有会员等级
// @Summary 查询所有会员等级
// @Description 查询所有会员等级
// @Tags UmsMemberLevelController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request formData listRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.UmsMemberLevel}
// @Failure 400 {object} code.Failure
// @Router /memberLevel/list [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}
	list, err := h.service.List(ctx, req.DefaultStatus)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}
