package ums_resource_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource_category"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	UmsResourceCateParam `json:",inline"`
}

type updateResponse struct {
	Count int64 `json:"count"`
}

// Update 修改后台资源分类
// @Summary 修改后台资源分类
// @Description 修改后台资源分类
// @Tags UmsResourceController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /resourceCategory/update/{id} [post]
func (h *handler) Update(ctx *gin.Context) {
	req := new(updateRequest)
	res := new(updateResponse)
	uri := new(UmsResourceCateUri)
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

	data := &ums_resource_category.UmsResourceCategory{
		Name: req.Name,
		Sort: req.Sort,
	}
	cnt, err := h.umsResourceCateService.Update(ctx, uri.Id, data)
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
