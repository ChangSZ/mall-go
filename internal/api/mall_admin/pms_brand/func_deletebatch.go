package pms_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type deleteBatchRequest struct {
	Ids []int64 `form:"ids" binding:"required"`
}

type deleteBatchResponse struct {
	Count int64 `json:",inline"`
}

// DeleteBatch 批量删除品牌
// @Summary 批量删除品牌
// @Description 批量删除品牌
// @Tags PmsBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteBatchRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /brand/delete/batch [post]
func (h *handler) DeleteBatch(ctx *gin.Context) {
	req := new(deleteBatchRequest)
	res := new(deleteBatchResponse)
	cnt, err := h.service.DeleteBatch(ctx, req.Ids)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
