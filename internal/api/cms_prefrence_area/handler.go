package cms_prefrence_area

import (
	"github.com/ChangSZ/mall-go/internal/services/cms_prefrence_area"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// ListAll 获取所有商品优选
	// @Tags CmsPrefrenceAreaController
	// @Router /prefrenceArea/listAll [get]
	ListAll(*gin.Context)
}

type handler struct {
	cmsPrefrenceAreaService cms_prefrence_area.Service
}

func New() Handler {
	return &handler{
		cmsPrefrenceAreaService: cms_prefrence_area.New(),
	}
}

func (h *handler) i() {}
