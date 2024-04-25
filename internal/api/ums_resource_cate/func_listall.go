package ums_resource_cate

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type listAllRequest struct{}

type listAllResponse struct {
	List []UmsResourceCate `json:",inline"`
}

// ListAll 查询所有后台资源分类
// @Summary 查询所有后台资源分类
// @Description 查询所有后台资源分类
// @Tags UmsResourceController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listAllRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]UmsResourceCate}
// @Failure 400 {object} code.Failure
// @Router /resourceCategory/listAll [get]
func (h *handler) ListAll(ctx *gin.Context) {
	_ = new(listAllRequest)
	res := new(listAllResponse)
	list, err := h.umsResourceCateService.ListAll(ctx)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	listData := make([]UmsResourceCate, 0, len(list))
	for _, v := range list {
		listData = append(listData, UmsResourceCate{
			Id:         v.Id,
			CreateTime: v.CreateTime,
			Name:       v.Name,
			Sort:       v.Sort,
		})
	}
	res.List = listData
	api.Success(ctx, res.List)
}
