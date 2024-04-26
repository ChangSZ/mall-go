package pms_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type getItemRequest struct{}

type getItemResponse struct {
	PmsBrand `json:",inline"`
}

// GetItem 根据编号查询品牌信息
// @Summary 根据编号查询品牌信息
// @Description 根据编号查询品牌信息
// @Tags PmsBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getItemRequest true "请求信息"
// @Success 200 {object} code.Success{data=getItemResponse}
// @Failure 400 {object} code.Failure
// @Router /brand/{id} [get]
func (h *handler) GetItem(ctx *gin.Context) {
	_ = new(getItemRequest)
	res := new(getItemResponse)
	uri := new(PmsBrandUri)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	item, err := h.pmsBrandService.GetItem(ctx, uri.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.PmsBrand = PmsBrand{
		Id:                  item.Id,
		Name:                item.Name,
		FirstLetter:         item.FirstLetter,
		Sort:                item.Sort,
		FactoryStatus:       item.FactoryStatus,
		ShowStatus:          item.ShowStatus,
		ProductCount:        item.ProductCount,
		ProductCommentCount: item.ProductCommentCount,
		Logo:                item.Logo,
		BigPic:              item.BigPic,
		BrandStory:          item.BrandStory,
	}
	api.Success(ctx, res)
}
