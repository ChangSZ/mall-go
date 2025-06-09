package cms_subject

import (
	"context"

	"github.com/ChangSZ/golib/copy"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/cms_subject"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) ListAll(ctx context.Context) ([]dto.CmsSubject, error) {
	qb := cms_subject.NewQueryBuilder()
	list, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	listData := make([]dto.CmsSubject, 0, len(list))
	for _, v := range list {
		tmp := dto.CmsSubject{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, nil
}

func (s *service) List(ctx context.Context, keyword string, pageSize, pageNum int) (
	[]dto.CmsSubject, int64, error) {
	qb := cms_subject.NewQueryBuilder()
	if keyword != "" {
		qb = qb.WhereTitle(mysql.LikePredicate, "%"+keyword+"%")
	}

	count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}

	offset := (pageNum - 1) * pageSize
	list, err := qb.
		Limit(pageSize).
		Offset(offset).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}

	listData := make([]dto.CmsSubject, 0, len(list))
	for _, v := range list {
		tmp := dto.CmsSubject{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, count, err
}
