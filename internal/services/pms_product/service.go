package pms_product

import (
	"context"
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/cms_prefrence_area_product_relation"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/cms_subject_product_relation"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_member_price"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_attribute_value"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_full_reduction"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_ladder"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product_vertify_record"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_sku_stock"
	pms_sku_stock_svc "github.com/ChangSZ/mall-go/internal/services/pms_sku_stock"
	"github.com/ChangSZ/mall-go/pkg/copy"
	"github.com/ChangSZ/mall-go/pkg/log"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, param dto.PmsProductParam) (int64, error) {
	// 创建商品
	data := &pms_product.PmsProduct{}
	copy.AssignStruct(&param, data)
	productId, err := data.Create(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return 0, err
	}
	// 根据促销类型设置价格：会员价格、阶梯价格、满减价格
	// 会员价格
	if err := new(dao.PmsMemberPriceDao).InsertList(ctx, mysql.DB().GetDbW().WithContext(ctx),
		param.MemberPriceList, productId); err != nil {
		log.WithTrace(ctx).Errorf("创建产品出错(PmsMemberPrice): %v", err)
		return productId, err
	}
	// 阶梯价格
	if err := new(dao.PmsProductLadderDao).InsertList(ctx, mysql.DB().GetDbW().WithContext(ctx),
		param.ProductLadderList, productId); err != nil {
		log.WithTrace(ctx).Errorf("创建产品出错(PmsProductLadder): %v", err)
		return productId, err
	}
	// 满减价格
	if err := new(dao.PmsProductFullReductionDao).InsertList(ctx, mysql.DB().GetDbW().WithContext(ctx),
		param.ProductFullReductionList, productId); err != nil {
		log.WithTrace(ctx).Errorf("创建产品出错(PmsProductFullReduction): %v", err)
		return productId, err
	}
	// 处理sku的编码
	s.handleSkuStockCode(param.SkuStockList, productId)
	// 添加sku库存信息
	if err := new(dao.PmsSkuStockDao).InsertList(ctx, mysql.DB().GetDbW().WithContext(ctx),
		param.SkuStockList, productId); err != nil {
		log.WithTrace(ctx).Errorf("创建产品出错(PmsSkuStock): %v", err)
		return productId, err
	}
	// 添加商品参数,添加自定义商品规格
	if err := new(dao.PmsProductAttributeValueDao).InsertList(ctx, mysql.DB().GetDbW().WithContext(ctx),
		param.ProductAttributeValueList, productId); err != nil {
		log.WithTrace(ctx).Errorf("创建产品出错(PmsProductAttributeValue): %v", err)
		return productId, err
	}
	// 关联专题
	if err := new(dao.CmsSubjectProductRelationDao).InsertList(ctx, mysql.DB().GetDbW().WithContext(ctx),
		param.SubjectProductRelationList, productId); err != nil {
		log.WithTrace(ctx).Errorf("创建产品出错(CmsSubjectProductRelation): %v", err)
		return productId, err
	}
	// 关联优选
	if err := new(dao.CmsPrefrenceAreaProductRelationDao).InsertList(ctx, mysql.DB().GetDbW().WithContext(ctx),
		param.PrefrenceAreaProductRelationList, productId); err != nil {
		log.WithTrace(ctx).Errorf("创建产品出错(CmsPrefrenceAreaProductRelation): %v", err)
		return productId, err
	}
	return productId, nil
}

func (s *service) handleSkuStockCode(skuStockList []dto.PmsSkuStock, productId int64) {
	if len(skuStockList) == 0 {
		return
	}

	for i := range skuStockList {
		if skuStockList[i].SkuCode == "" {
			// 日期
			today := time.Now().Format("20060102")
			// 四位商品id
			productIDStr := fmt.Sprintf("%04d", productId)
			// 三位索引id
			indexIDStr := fmt.Sprintf("%03d", i+1)

			skuCode := today + productIDStr + indexIDStr
			skuStockList[i].SkuCode = skuCode
		}
	}
}

func (s *service) GetUpdateInfo(ctx context.Context, id int64) (*dto.PmsProductResult, error) {
	return new(dao.PmsProductDao).GetUpdateInfo(ctx, mysql.DB().GetDbR().WithContext(ctx), id)
}

func (s *service) Update(ctx context.Context, id int64, param dto.PmsProductParam) (int64, error) {
	// TODO: 后续可以考虑使用事务, 比如gorm的关联模式
	// 更新商品信息
	data := map[string]interface{}{
		"id":                            id,
		"brand_id":                      param.BrandId,
		"product_category_id":           param.ProductCategoryId,
		"feight_template_id":            param.FeightTemplateId,
		"product_attribute_category_id": param.ProductAttributeCategoryId,
		"name":                          param.Name,
		"pic":                           param.Pic,
		"product_sn":                    param.ProductSn,
		"delete_status":                 param.DeleteStatus,
		"publish_status":                param.PublishStatus,
		"new_status":                    param.NewStatus,
		"recommend_status":              param.RecommendStatus,
		"verify_status":                 param.VerifyStatus,
		"sort":                          param.Sort,
		"sale":                          param.Sale,
		"price":                         param.Price,
		"promotion_price":               param.PromotionPrice,
		"gift_growth":                   param.GiftGrowth,
		"gift_point":                    param.GiftPoint,
		"use_point_limit":               param.UsePointLimit,
		"sub_title":                     param.SubTitle,
		"description":                   param.Description,
		"original_price":                param.OriginalPrice,
		"stock":                         param.Stock,
		"low_stock":                     param.LowStock,
		"unit":                          param.Unit,
		"weight":                        param.Weight,
		"preview_status":                param.PreviewStatus,
		"service_ids":                   param.ServiceIds,
		"keywords":                      param.Keywords,
		"note":                          param.Note,
		"album_pics":                    param.AlbumPics,
		"detail_title":                  param.DetailTitle,
		"detail_desc":                   param.DetailDesc,
		"detail_html":                   param.DetailHtml,
		"detail_mobile_html":            param.DetailMobileHtml,
		"promotion_start_time":          param.PromotionStartTime,
		"promotion_end_time":            param.PromotionEndTime,
		"promotion_per_limit":           param.PromotionPerLimit,
		"promotion_type":                param.PromotionType,
		"brand_name":                    param.BrandName,
		"product_category_name":         param.ProductCategoryName,
	}
	qb := pms_product.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	cnt, err := qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return 0, err
	}

	// 会员价格
	{
		qb := pms_member_price.NewQueryBuilder()
		qb = qb.WhereProductId(mysql.EqualPredicate, id)
		if _, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
			log.WithTrace(ctx).Errorf("删除产品出错(PmsMemberPrice): %v", err)
			return cnt, err
		}
		if err := new(dao.PmsMemberPriceDao).InsertList(ctx, mysql.DB().GetDbW().WithContext(ctx),
			param.MemberPriceList, id); err != nil {
			log.WithTrace(ctx).Errorf("创建产品出错(PmsMemberPrice): %v", err)
			return cnt, err
		}
	}
	// 阶梯价格
	{
		qb := pms_product_ladder.NewQueryBuilder()
		qb = qb.WhereProductId(mysql.EqualPredicate, id)
		if _, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
			log.WithTrace(ctx).Errorf("删除产品出错(PmsProductLadder): %v", err)
			return cnt, err
		}
		if err := new(dao.PmsProductLadderDao).InsertList(ctx, mysql.DB().GetDbW().WithContext(ctx),
			param.ProductLadderList, id); err != nil {
			log.WithTrace(ctx).Errorf("创建产品出错(PmsProductLadder): %v", err)
			return cnt, err
		}
	}
	// 满减价格
	{
		qb := pms_product_full_reduction.NewQueryBuilder()
		qb = qb.WhereProductId(mysql.EqualPredicate, id)
		if _, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
			log.WithTrace(ctx).Errorf("删除产品出错(PmsProductFullReduction): %v", err)
			return cnt, err
		}
		if err := new(dao.PmsProductFullReductionDao).InsertList(ctx, mysql.DB().GetDbW().WithContext(ctx),
			param.ProductFullReductionList, id); err != nil {
			log.WithTrace(ctx).Errorf("创建产品出错(PmsProductFullReduction): %v", err)
			return cnt, err
		}
	}

	// 修改sku库存信息
	if err := s.handleUpdateSkuStockList(ctx, param.SkuStockList, id); err != nil {
		log.WithTrace(ctx).Errorf("更新sku库存信息出错: %v", err)
		return cnt, err
	}

	// 修改商品参数,添加自定义商品规格
	{
		qb := pms_product_attribute_value.NewQueryBuilder()
		qb = qb.WhereProductId(mysql.EqualPredicate, id)
		if _, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
			log.WithTrace(ctx).Errorf("删除产品出错(PmsProductAttributeValue): %v", err)
			return cnt, err
		}
		if err := new(dao.PmsProductAttributeValueDao).InsertList(ctx, mysql.DB().GetDbW().WithContext(ctx),
			param.ProductAttributeValueList, id); err != nil {
			log.WithTrace(ctx).Errorf("创建产品出错(PmsProductAttributeValue): %v", err)
			return cnt, err
		}
	}
	// 关联专题
	{
		qb := cms_subject_product_relation.NewQueryBuilder()
		qb = qb.WhereProductId(mysql.EqualPredicate, id)
		if _, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
			log.WithTrace(ctx).Errorf("删除产品出错(CmsSubjectProductRelation): %v", err)
			return cnt, err
		}
		if err := new(dao.CmsSubjectProductRelationDao).InsertList(ctx, mysql.DB().GetDbW().WithContext(ctx),
			param.SubjectProductRelationList, id); err != nil {
			log.WithTrace(ctx).Errorf("创建产品出错(CmsSubjectProductRelation): %v", err)
			return cnt, err
		}
	}
	// 关联优选
	{
		qb := cms_prefrence_area_product_relation.NewQueryBuilder()
		qb = qb.WhereProductId(mysql.EqualPredicate, id)
		if _, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
			log.WithTrace(ctx).Errorf("删除产品出错(CmsPrefrenceAreaProductRelation): %v", err)
			return cnt, err
		}
		if err := new(dao.CmsPrefrenceAreaProductRelationDao).InsertList(ctx, mysql.DB().GetDbW().WithContext(ctx),
			param.PrefrenceAreaProductRelationList, id); err != nil {
			log.WithTrace(ctx).Errorf("创建产品出错(CmsPrefrenceAreaProductRelation): %v", err)
			return cnt, err
		}
	}
	return cnt, nil
}

func (s *service) handleUpdateSkuStockList(ctx context.Context, currSkuList []dto.PmsSkuStock, id int64) error {
	// 当前没有sku直接删除
	if len(currSkuList) == 0 {
		qb := pms_sku_stock.NewQueryBuilder()
		qb = qb.WhereProductId(mysql.EqualPredicate, id)
		if _, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
			return err
		}
		return nil
	}
	// 获取初始sku信息
	qb := pms_sku_stock.NewQueryBuilder()
	qb = qb.WhereProductId(mysql.EqualPredicate, id)
	oriStuList, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return err
	}
	// 获取新增sku信息
	var insertSkuList = make([]dto.PmsSkuStock, 0)
	// 获取需要更新的sku信息
	var updateSkuList = make([]dto.PmsSkuStock, 0)
	var updateSkuIds = make(map[int64]bool, 0)
	for _, v := range currSkuList {
		if v.Id == 0 {
			insertSkuList = append(insertSkuList, v)
		} else {
			updateSkuList = append(updateSkuList, v)
			updateSkuIds[v.Id] = true
		}
	}
	// 获取需要删除的sku信息
	var removeSkuIds = make([]int64, 0)
	for _, v := range oriStuList {
		if _, ok := updateSkuIds[v.Id]; !ok {
			removeSkuIds = append(removeSkuIds, v.Id)
		}
	}

	// 删除sku
	if len(removeSkuIds) > 0 {
		qb := pms_sku_stock.NewQueryBuilder()
		qb = qb.WhereIdIn(removeSkuIds)
		if _, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
			return err
		}
		return nil
	}

	// 新增sku
	if len(insertSkuList) > 0 {
		s.handleSkuStockCode(insertSkuList, id)
		if err := new(dao.PmsSkuStockDao).InsertList(ctx, mysql.DB().GetDbW().WithContext(ctx),
			insertSkuList, id); err != nil {
			return err
		}
	}

	// 修改sku
	if len(updateSkuList) > 0 {
		s.handleSkuStockCode(updateSkuList, id)
		svc := pms_sku_stock_svc.New()
		for _, v := range updateSkuList {
			if _, err := svc.Update(ctx, v.Id, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *service) List(ctx context.Context, queryParam dto.PmsProductQueryParam, pageSize, pageNum int) (
	[]dto.PmsProduct, int64, error) {
	qb := pms_product.NewQueryBuilder()
	qb = qb.WhereDeleteStatus(mysql.EqualPredicate, 0)

	if queryParam.PublishStatus != 0 {
		qb = qb.WherePublishStatus(mysql.EqualPredicate, queryParam.PublishStatus)
	}
	if queryParam.VerifyStatus != 0 {
		qb = qb.WhereVerifyStatus(mysql.EqualPredicate, queryParam.VerifyStatus)
	}
	if queryParam.Keyword != "" {
		qb = qb.WhereName(mysql.LikePredicate, "%"+queryParam.Keyword+"%")
	}
	if queryParam.ProductSn != "" {
		qb = qb.WhereProductSn(mysql.EqualPredicate, queryParam.ProductSn)
	}
	if queryParam.BrandId != 0 {
		qb = qb.WhereBrandId(mysql.EqualPredicate, queryParam.BrandId)
	}
	if queryParam.ProductCategoryId != 0 {
		qb = qb.WhereProductCategoryId(mysql.EqualPredicate, queryParam.ProductCategoryId)
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

	listData := make([]dto.PmsProduct, 0, len(list))
	for _, v := range list {
		data := dto.PmsProduct{}
		copy.AssignStruct(v, &data)
		listData = append(listData, data)
	}
	return listData, count, err
}

func (s *service) UpdateVerifyStatus(ctx context.Context,
	ids []int64, verifyStatus int32, detail string) (int64, error) {
	data := map[string]interface{}{
		"verify_status": verifyStatus,
	}
	qb := pms_product.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	cnt, err := qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return 0, err
	}
	// 修改完审核状态后插入审核记录
	recordList := make([]pms_product_vertify_record.PmsProductVertifyRecord, 0)
	for _, id := range ids {
		record := pms_product_vertify_record.PmsProductVertifyRecord{
			ProductId:  id,
			VertifyMan: "test",
			Status:     verifyStatus,
			Detail:     detail,
		}
		recordList = append(recordList, record)
	}
	return cnt, mysql.DB().GetDbW().WithContext(ctx).CreateInBatches(recordList, len(recordList)).Error
}

func (s *service) UpdatePublishStatus(ctx context.Context, ids []int64, publishStatus int32) (int64, error) {
	data := map[string]interface{}{
		"publish_status": publishStatus,
	}
	qb := pms_product.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) UpdateRecommendStatus(ctx context.Context, ids []int64, recommendStatus int32) (int64, error) {
	data := map[string]interface{}{
		"recommend_status": recommendStatus,
	}
	qb := pms_product.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) UpdateNewStatus(ctx context.Context, ids []int64, newStatus int32) (int64, error) {
	data := map[string]interface{}{
		"new_status": newStatus,
	}
	qb := pms_product.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) UpdateDeleteStatus(ctx context.Context, ids []int64, deleteStatus int32) (int64, error) {
	data := map[string]interface{}{
		"delete_status": deleteStatus,
	}
	qb := pms_product.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) SimpleList(ctx context.Context, keyword string) ([]dto.PmsProduct, error) {
	list, err := new(dao.PmsProductDao).ListByKeyword(ctx, mysql.DB().GetDbR().WithContext(ctx), keyword)
	if err != nil {
		return nil, err
	}
	listData := make([]dto.PmsProduct, 0, len(list))
	for _, v := range list {
		tmp := dto.PmsProduct{}
		copy.AssignStruct(&v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, nil
}
