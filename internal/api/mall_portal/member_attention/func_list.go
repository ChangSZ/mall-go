package member_attention

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type listRequest struct {
	PageNum  int64 `form:"pageNum,default=1" binding:"omitempty"`
	PageSize int64 `form:"pageSize,default=5" binding:"omitempty"`
}

type listResponse struct {
	PageNum   int64                      `json:"pageNum"`
	PageSize  int64                      `json:"pageSize"`
	TotalPage int64                      `json:"totalPage"`
	Total     int64                      `json:"total"`
	List      []dto.MemberBrandAttention `json:"list"`
}

// List 分页查询当前用户品牌关注列表
// @Summary 分页查询当前用户品牌关注列表
// @Description 分页查询当前用户品牌关注列表
// @Tags MemberAttentionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /member/attention/list [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, total, err := h.service.List(ctx, req.PageNum, req.PageSize)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.PageNum = req.PageNum
	res.PageSize = req.PageSize
	totalPage := total / int64(req.PageSize)
	if total%int64(req.PageSize) > 0 {
		totalPage += 1
	}
	res.TotalPage = totalPage
	res.Total = total
	res.List = list
	api.Success(ctx, res)
}
