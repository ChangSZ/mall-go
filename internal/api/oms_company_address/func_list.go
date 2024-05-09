package oms_company_address

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 获取所有收货地址
// @Summary 获取所有收货地址
// @Description 获取所有收货地址
// @Tags OmsCompanyAddressController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /companyAddress/list [get]
func (h *handler) List(ctx *gin.Context) {
	api.Success(ctx, nil)
}
