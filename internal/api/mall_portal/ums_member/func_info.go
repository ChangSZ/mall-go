package ums_member

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type infoRequest struct{}

type infoResponse struct {
	*dto.UmsMember `json:",inline"`
}

// Info 获取会员信息
// @Summary 获取会员信息
// @Description 获取会员信息
// @Tags UmsMemberController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body infoRequest true "请求信息"
// @Success 200 {object} code.Success{data=infoResponse}
// @Failure 400 {object} code.Failure
// @Router /sso/info [get]
func (h *handler) Info(ctx *gin.Context) {
	_ = new(infoRequest)
	res := new(infoResponse)
	member, err := h.service.GetCurrentMember(ctx)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.UmsMember = member
	api.Success(ctx, res)
}
