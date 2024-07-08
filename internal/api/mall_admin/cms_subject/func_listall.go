package cms_subject

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type listAllRequest struct{}

type listAllResponse struct {
	List []dto.CmsSubject `json:",inline"`
}

// ListAll 获取全部商品专题
// @Summary 获取全部商品专题
// @Description 获取全部商品专题
// @Tags CmsSubjectController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listAllRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.CmsSubject}
// @Failure 400 {object} code.Failure
// @Router /subject/listAll [get]
func (h *handler) ListAll(ctx *gin.Context) {
	_ = new(listAllRequest)
	res := new(listAllResponse)
	list, err := h.service.ListAll(ctx)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}
