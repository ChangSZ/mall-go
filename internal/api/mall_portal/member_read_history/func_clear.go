package member_read_history

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type clearRequest struct{}

type clearResponse struct {
	Count int64 `json:",inline"`
}

// Clear 清空浏览记录
// @Summary 清空浏览记录
// @Description 清空浏览记录
// @Tags MemberCollectionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body clearRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /member/readHistory/clear [post]
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
