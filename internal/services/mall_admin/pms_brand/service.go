package pms_brand

import (
	"context"
	"strings"

	"github.com/ChangSZ/golib/copy"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_brand"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, param dto.PmsBrandParam) (int64, error) {
	data := &pms_brand.PmsBrand{
		Name:                param.Name,
		FirstLetter:         param.FirstLetter,
		Sort:                param.Sort,
		FactoryStatus:       param.FactoryStatus,
		ShowStatus:          param.ShowStatus,
		ProductCount:        param.ProductCount,
		ProductCommentCount: param.ProductCommentCount,
		Logo:                param.Logo,
		BigPic:              param.BigPic,
		BrandStory:          param.BrandStory,
	}
	// 如果创建时首字母为空，取名称的第一个为首字母
	if data.FirstLetter == "" && len(data.Name) > 0 {
		data.FirstLetter = data.Name[:1]
	}
	return data.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Update(ctx context.Context, id int64, param dto.PmsBrandParam) (int64, error) {
	// 如果创建时首字母为空，取名称的第一个为首字母
	if param.FirstLetter == "" && len(param.Name) > 0 {
		param.FirstLetter = param.Name[:1]
	}

	// 更新品牌时要更新商品中的品牌名称
	pmsProduct := map[string]interface{}{"brand_name": param.Name}
	pmsProductQb := pms_product.NewQueryBuilder()
	pmsProductQb = pmsProductQb.WhereBrandId(mysql.EqualPredicate, id)
	if _, err := pmsProductQb.Updates(mysql.DB().GetDbW().WithContext(ctx), pmsProduct); err != nil {
		return 0, err
	}

	data := map[string]interface{}{
		"id":                    id,
		"name":                  param.Name,
		"first_letter":          param.FirstLetter,
		"sort":                  param.Sort,
		"factory_status":        param.FactoryStatus,
		"show_status":           param.ShowStatus,
		"product_count":         param.ProductCount,
		"product_comment_count": param.ProductCommentCount,
		"logo":                  param.Logo,
		"big_pic":               param.BigPic,
		"brand_story":           param.BrandStory,
	}
	qb := pms_brand.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	return s.DeleteBatch(ctx, []int64{id})
}

func (s *service) DeleteBatch(ctx context.Context, ids []int64) (int64, error) {
	qb := pms_brand.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) ListAll(ctx context.Context) ([]dto.PmsBrand, error) {
	qb := pms_brand.NewQueryBuilder()
	list, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	listData := make([]dto.PmsBrand, 0, len(list))
	for _, v := range list {
		listData = append(listData, dto.PmsBrand{
			Id:                  v.Id,
			Name:                v.Name,
			FirstLetter:         v.FirstLetter,
			Sort:                v.Sort,
			FactoryStatus:       v.FactoryStatus,
			ShowStatus:          v.ShowStatus,
			ProductCount:        v.ProductCount,
			ProductCommentCount: v.ProductCommentCount,
			Logo:                v.Logo,
			BigPic:              v.BigPic,
			BrandStory:          v.BrandStory,
		})
	}
	return listData, nil
}

func (s *service) List(ctx context.Context, keyword string, showStatus int32, pageSize, pageNum int) (
	*pagehelper.ListData[dto.PmsBrand], error) {
	qb := pms_brand.NewQueryBuilder()
	if strings.TrimSpace(keyword) != "" {
		qb = qb.WhereName(mysql.LikePredicate, "%"+keyword+"%")
	}
	if showStatus != 0 {
		qb = qb.WhereShowStatus(mysql.EqualPredicate, showStatus)
	}

	res := pagehelper.New[dto.PmsBrand]()
	count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return res, err
	}

	offset := (pageNum - 1) * pageSize
	list, err := qb.
		Limit(pageSize).
		Offset(offset).
		OrderBySort(false).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return res, err
	}

	listData := make([]dto.PmsBrand, 0, len(list))
	for _, v := range list {
		tmp := dto.PmsBrand{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}

	res.Set(pageNum, pageSize, count, listData)
	return res, err
}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.PmsBrand, error) {
	qb := pms_brand.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	item, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return &dto.PmsBrand{
		Id:                  item.Id,
		Name:                item.Name,
		FirstLetter:         item.FirstLetter,
		Sort:                item.Sort,
		FactoryStatus:       item.FactoryStatus,
		ShowStatus:          item.ShowStatus,
		ProductCount:        item.ProductCount,
		ProductCommentCount: item.ProductCommentCount,
		Logo:                item.Logo,
		BigPic:              item.BigPic,
		BrandStory:          item.BrandStory,
	}, nil
}

func (s *service) UpdateShowStatus(ctx context.Context, ids []int64, showStatus int32) (int64, error) {
	data := map[string]interface{}{
		"show_status": showStatus,
	}
	qb := pms_brand.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) UpdateFactoryStatus(ctx context.Context, ids []int64, factoryStatus int32) (int64, error) {
	data := map[string]interface{}{
		"factory_status": factoryStatus,
	}
	qb := pms_brand.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}
