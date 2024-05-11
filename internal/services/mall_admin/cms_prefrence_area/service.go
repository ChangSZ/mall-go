package cms_prefrence_area

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/cms_prefrence_area"
	"github.com/ChangSZ/mall-go/pkg/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) ListAll(ctx context.Context) ([]dto.CmsPrefrenceArea, error) {
	qb := cms_prefrence_area.NewQueryBuilder()
	list, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	listData := make([]dto.CmsPrefrenceArea, 0, len(list))
	for _, v := range list {
		tmp := dto.CmsPrefrenceArea{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, nil
}
