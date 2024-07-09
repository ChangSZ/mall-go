package oms_company_address

import (
	"context"

	"github.com/ChangSZ/golib/copy"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/oms_company_address"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) List(ctx context.Context) ([]dto.OmsCompanyAddress, error) {
	qb := oms_company_address.NewQueryBuilder()
	list, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	listData := make([]dto.OmsCompanyAddress, 0, len(list))
	for _, v := range list {
		tmp := dto.OmsCompanyAddress{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, nil
}
