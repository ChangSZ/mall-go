package pms_brand

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type listAllRequest struct{}

type listAllResponse struct {
	List []PmsBrand `json:",inline"`
}

// ListAll 获取全部品牌列表
// @Summary 获取全部品牌列表
// @Description 获取全部品牌列表
// @Tags PmsBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listAllRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]PmsBrand}
// @Failure 400 {object} code.Failure
// @Router /brand/listAll [get]
func (h *handler) ListAll(ctx *gin.Context) {
	_ = new(listAllRequest)
	res := new(listAllResponse)
	list, err := h.pmsBrandService.ListAll(ctx)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	listData := make([]PmsBrand, 0, len(list))
	for _, v := range list {
		listData = append(listData, PmsBrand{
			Id:                  v.Id,
			Name:                v.Name,
			FirstLetter:         v.FirstLetter,
			Sort:                v.Sort,
			FactoryStatus:       v.FactoryStatus,
			ShowStatus:          v.ShowStatus,
			ProductCount:        v.ProductCount,
			ProductCommentCount: v.ProductCommentCount,
			Logo:                v.Logo,
			BigPic:              v.BigPic,
			BrandStory:          v.BrandStory,
		})
	}
	res.List = listData
	api.Success(ctx, res.List)
}
