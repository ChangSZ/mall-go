package ums_resource

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type listAllRequest struct{}

type listAllResponse struct {
	List []UmsResource `json:",inline"`
}

// ListAll 查询所有后台资源
// @Summary 查询所有后台资源
// @Description 查询所有后台资源
// @Tags UmsResourceController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listAllRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]UmsResource}
// @Failure 400 {object} code.Failure
// @Router /resource/listAll [get]
func (h *handler) ListAll(ctx *gin.Context) {
	_ = new(listAllRequest)
	res := new(listAllResponse)
	list, err := h.umsResourceService.ListAll(ctx)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = make([]UmsResource, 0, len(list))
	for _, v := range list {
		res.List = append(res.List, UmsResource{
			Id:          v.Id,
			CreateTime:  v.CreateTime,
			Name:        v.Name,
			Url:         v.Url,
			Description: v.Description,
			CategoryId:  v.CategoryId,
		})
	}
	api.Success(ctx, res.List)
}
