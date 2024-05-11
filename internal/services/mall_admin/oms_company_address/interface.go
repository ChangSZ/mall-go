package oms_company_address

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 收货地址管理Service
type Service interface {
	i()

	/**
	 * 获取全部收货地址
	 */
	List(ctx context.Context) ([]dto.OmsCompanyAddress, error)
}
