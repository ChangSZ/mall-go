package pms_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_brand"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	PmsBrandParam `json:",inline"`
}

type updateResponse struct {
	Count int64 `json:",inline"`
}

// Update 更新品牌
// @Summary 更新品牌
// @Description 更新品牌
// @Tags PmsBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /brand/update/{id} [post]
func (h *handler) Update(ctx *gin.Context) {
	req := new(updateRequest)
	res := new(updateResponse)
	uri := new(PmsBrandUri)
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

	data := &pms_brand.PmsBrand{
		Name:                req.Name,
		FirstLetter:         req.FirstLetter,
		Sort:                req.Sort,
		FactoryStatus:       req.FactoryStatus,
		ShowStatus:          req.ShowStatus,
		ProductCount:        req.ProductCount,
		ProductCommentCount: req.ProductCommentCount,
		Logo:                req.Logo,
		BigPic:              req.BigPic,
		BrandStory:          req.BrandStory,
	}
	cnt, err := h.pmsBrandService.Update(ctx, uri.Id, data)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if cnt == 0 {
		api.Failed(ctx, "更新个数为0")
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
