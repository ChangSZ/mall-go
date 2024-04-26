package pms_brand

import (
	"context"
	"strings"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_brand"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, data *pms_brand.PmsBrand) (int64, error) {
	// 如果创建时首字母为空，取名称的第一个为首字母
	if data.FirstLetter == "" && len(data.Name) > 0 {
		data.FirstLetter = data.Name[:1]
	}
	return data.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Update(ctx context.Context, id int64, data *pms_brand.PmsBrand) (int64, error) {
	data.Id = id
	// 如果创建时首字母为空，取名称的第一个为首字母
	if data.FirstLetter == "" && len(data.Name) > 0 {
		data.FirstLetter = data.Name[:1]
	}

	// 更新品牌时要更新商品中的品牌名称
	pmsProduct := pms_product.NewModel()
	pmsProduct.BrandName = data.Name
	pmsProductQb := pms_product.NewQueryBuilder()
	pmsProductQb = pmsProductQb.WhereBrandId(mysql.EqualPredicate, id)
	if _, err := pmsProductQb.Update(mysql.DB().GetDbW().WithContext(ctx), pmsProduct); err != nil {
		return 0, err
	}

	qb := pms_brand.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	cnt, err := qb.Update(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	return s.DeleteBatch(ctx, []int64{id})
}

func (s *service) DeleteBatch(ctx context.Context, ids []int64) (int64, error) {
	qb := pms_brand.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	cnt, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil || cnt == 0 {
		return 0, err
	}
	err = qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func (s *service) ListAll(ctx context.Context) ([]*pms_brand.PmsBrand, error) {
	qb := pms_brand.NewQueryBuilder()
	return qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
}

func (s *service) List(ctx context.Context, keyword string, showStatus int32, pageSize, pageNum int) (
	[]*pms_brand.PmsBrand, int64, error) {
	qb := pms_brand.NewQueryBuilder()
	if strings.TrimSpace(keyword) != "" {
		qb = qb.WhereName(mysql.LikePredicate, "%"+keyword+"%")
	}
	if showStatus != 0 {
		qb = qb.WhereShowStatus(mysql.EqualPredicate, showStatus)
	}
	count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}
	offset := (pageNum - 1) * pageSize
	list, err := qb.
		Limit(pageSize).
		Offset(offset).
		OrderBySort(false).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	return list, count, err
}

func (s *service) GetItem(ctx context.Context, id int64) (*pms_brand.PmsBrand, error) {
	qb := pms_brand.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	return qb.First(mysql.DB().GetDbR().WithContext(ctx))
}

func (s *service) UpdateShowStatus(ctx context.Context, ids []int64, showStatus int32) (int64, error) {
	data := pms_brand.NewModel()
	data.ShowStatus = showStatus

	qb := pms_brand.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	cnt, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return 0, err
	}
	err = qb.Updates(mysql.DB().GetDbW().WithContext(ctx), map[string]interface{}{"show_status": showStatus})
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func (s *service) UpdateFactoryStatus(ctx context.Context, ids []int64, factoryStatus int32) (int64, error) {
	data := pms_brand.NewModel()
	data.FactoryStatus = factoryStatus

	qb := pms_brand.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	cnt, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return 0, err
	}
	err = qb.Updates(mysql.DB().GetDbW().WithContext(ctx), map[string]interface{}{"factory_status": factoryStatus})
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
