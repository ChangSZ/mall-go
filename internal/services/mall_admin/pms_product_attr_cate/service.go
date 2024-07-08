package pms_product_attr_cate

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_attribute_category"

	"github.com/ChangSZ/golib/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, name string) (int64, error) {
	data := pms_product_attribute_category.NewModel()
	data.Name = name
	return data.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Update(ctx context.Context, id int64, name string) (int64, error) {
	data := map[string]interface{}{
		name: name,
	}
	qb := pms_product_attribute_category.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	qb := pms_product_attribute_category.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.PmsProductAttributeCategory, error) {
	qb := pms_product_attribute_category.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	data, err := qb.First(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	res := &dto.PmsProductAttributeCategory{}
	copy.AssignStruct(data, res)
	return res, nil
}

func (s *service) List(ctx context.Context, pageSize, pageNum int) ([]dto.PmsProductAttributeCategory, int64, error) {
	qb := pms_product_attribute_category.NewQueryBuilder()
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

	listData := make([]dto.PmsProductAttributeCategory, 0, len(list))
	for _, v := range list {
		tmp := dto.PmsProductAttributeCategory{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, count, err
}

func (s *service) ListWithAttr(ctx context.Context) ([]dto.PmsProductAttrCateItem, error) {
	list, err := new(dao.PmsProductAttrCateDao).ListWithAttr(ctx, mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	res := make([]dto.PmsProductAttrCateItem, 0, len(list))
	for _, v := range list {
		tmp := dto.PmsProductAttrCateItem{
			ProductAttributeList: make([]dto.PmsProductAttribute, len(v.ProductAttributeList)),
		}
		copy.AssignStruct(&v, &tmp)
		res = append(res, tmp)
	}
	return res, nil
}
