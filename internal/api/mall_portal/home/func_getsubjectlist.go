package home

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type getSubjectListRequest struct{}

type getSubjectListResponse struct{}

// GetSubjectList 根据分类获取专题
// @Summary 根据分类获取专题
// @Description 根据分类获取专题
// @Tags HomeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getSubjectListRequest true "请求信息"
// @Success 200 {object} code.Success{data=getSubjectListResponse}
// @Failure 400 {object} code.Failure
// @Router /home/subjectList [get]
func (h *handler) GetSubjectList(ctx *gin.Context) {
	api.Success(ctx, nil)
}
