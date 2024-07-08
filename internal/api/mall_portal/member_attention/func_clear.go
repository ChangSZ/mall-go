package member_attention

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type clearRequest struct{}

type clearResponse struct {
	Count int64 `json:",inline"`
}

// Clear 清空当前用户品牌关注列表
// @Summary 清空当前用户品牌关注列表
// @Description 清空当前用户品牌关注列表
// @Tags MemberAttentionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body clearRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /member/attention/clear [post]
func (h *handler) Clear(ctx *gin.Context) {
	_ = new(clearRequest)
	res := new(clearResponse)
	count, err := h.service.Clear(ctx)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = count
	api.Success(ctx, res.Count)
}
