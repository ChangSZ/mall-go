package ums_resource_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	dto.UmsResourceCateParam `json:",inline"`
}

type createResponse struct {
	Count int64 `json:"count"`
}

// Create 添加后台资源分类
// @Summary 添加后台资源分类
// @Description 添加后台资源分类
// @Tags UmsResourceController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /resourceCategory/create [post]
func (h *handler) Create(ctx *gin.Context) {
	req := new(createRequest)
	res := new(createResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.Create(ctx, req.UmsResourceCateParam)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
