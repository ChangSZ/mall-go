package member_read_history

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type listRequest struct {
	PageNum  int64 `form:"pageNum,default=1" binding:"omitempty"`
	PageSize int64 `form:"pageSize,default=5" binding:"omitempty"`
}

type listResponse struct {
	*pagehelper.ListData[dto.MemberReadHistory] `json:",inline"`
}

// List 分页获取浏览记录
// @Summary 分页获取浏览记录
// @Description 分页获取浏览记录
// @Tags MemberCollectionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /member/readHistory/list [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	list, err := h.service.List(ctx, req.PageNum, req.PageSize)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.ListData = list
	api.Success(ctx, res)
}
