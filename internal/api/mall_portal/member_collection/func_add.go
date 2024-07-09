package member_collection

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type addRequest struct {
	dto.MemberProductCollection `json:",inline"`
}

type addResponse struct {
	Count int64 `json:",inline"`
}

// Add 添加商品收藏
// @Summary 添加商品收藏
// @Description 添加商品收藏
// @Tags MemberCollectionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body addRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /member/productCollection/add [post]
func (h *handler) Add(ctx *gin.Context) {
	req := new(addRequest)
	res := new(addResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	count, err := h.service.Add(ctx, req.MemberProductCollection)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = count
	api.Success(ctx, res.Count)
}
