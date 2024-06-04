package home

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type getSubjectListRequest struct {
	CateId   int64 `form:"cateId" binding:"omitempty"`
	PageNum  int   `form:"pageNum,default=1" binding:"omitempty"`
	PageSize int   `form:"pageSize,default=4" binding:"omitempty"`
}

type getSubjectListResponse struct {
	List []dto.CmsSubject `json:",inline"`
}

// GetSubjectList 根据分类获取专题
// @Summary 根据分类获取专题
// @Description 根据分类获取专题
// @Tags HomeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getSubjectListRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.CmsSubject}
// @Failure 400 {object} code.Failure
// @Router /home/subjectList [get]
func (h *handler) GetSubjectList(ctx *gin.Context) {
	req := new(getSubjectListRequest)
	res := new(getSubjectListResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, err := h.service.GetSubjectList(ctx, req.CateId, req.PageNum, req.PageSize)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}
