package cms_subject

import (
	"github.com/ChangSZ/mall-go/internal/services/cms_subject"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// ListAll 获取全部商品专题
	// @Tags CmsSubjectController
	// @Router /subject/listAll [get]
	ListAll(*gin.Context)

	// List 根据专题名称分页获取商品专题
	// @Tags CmsSubjectController
	// @Router /subject/list [get]
	List(*gin.Context)
}

type handler struct {
	cmsSubjectService cms_subject.Service
}

func New() Handler {
	return &handler{
		cmsSubjectService: cms_subject.New(),
	}
}

func (h *handler) i() {}
