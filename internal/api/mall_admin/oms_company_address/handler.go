package oms_company_address

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_admin/oms_company_address"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// List 获取所有收货地址
	// @Tags OmsCompanyAddressController
	// @Router /companyAddress/list [get]
	List(*gin.Context)
}

type handler struct {
	service oms_company_address.Service
}

func New() Handler {
	return &handler{
		service: oms_company_address.New(),
	}
}

func (h *handler) i() {}
