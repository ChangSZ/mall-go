package oms_portal_order_return_apply

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/oms_order_return_apply"

	"github.com/ChangSZ/golib/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, param dto.OmsOrderReturnApplyParam) (int64, error) {
	realApply := &oms_order_return_apply.OmsOrderReturnApply{}
	copy.AssignStruct(&param, realApply)
	realApply.Status = 0
	return realApply.Create(mysql.DB().GetDbW().WithContext(ctx))
}
