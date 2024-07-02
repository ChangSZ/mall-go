///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package oms_order_item

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *OmsOrderItem {
	return new(OmsOrderItem)
}

func NewQueryBuilder() *omsOrderItemQueryBuilder {
	return new(omsOrderItemQueryBuilder)
}

func (t *OmsOrderItem) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type omsOrderItemQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *omsOrderItemQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	if qb.limit != 0 {
		ret = ret.Limit(qb.limit)
	}
	ret = ret.Offset(qb.offset)
	return ret
}

func (qb *omsOrderItemQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&OmsOrderItem{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	ret := db.Updates(m)
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "updates err")
	}
	return ret.RowsAffected, nil
}

func (qb *omsOrderItemQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&OmsOrderItem{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *omsOrderItemQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&OmsOrderItem{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *omsOrderItemQueryBuilder) First(db *gorm.DB) (*OmsOrderItem, error) {
	ret := &OmsOrderItem{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *omsOrderItemQueryBuilder) QueryOne(db *gorm.DB) (*OmsOrderItem, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *omsOrderItemQueryBuilder) QueryAll(db *gorm.DB) ([]*OmsOrderItem, error) {
	var ret []*OmsOrderItem
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *omsOrderItemQueryBuilder) Limit(limit int) *omsOrderItemQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *omsOrderItemQueryBuilder) Offset(offset int) *omsOrderItemQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereId(p mysql.Predicate, value int64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereIdIn(value []int64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereIdNotIn(value []int64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderById(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereOrderId(p mysql.Predicate, value int64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_id", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereOrderIdIn(value []int64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_id", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereOrderIdNotIn(value []int64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByOrderId(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "order_id "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereOrderSn(p mysql.Predicate, value string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_sn", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereOrderSnIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_sn", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereOrderSnNotIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_sn", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByOrderSn(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "order_sn "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductId(p mysql.Predicate, value int64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductIdIn(value []int64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductIdNotIn(value []int64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByProductId(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_id "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductPic(p mysql.Predicate, value string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_pic", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductPicIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_pic", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductPicNotIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_pic", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByProductPic(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_pic "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductName(p mysql.Predicate, value string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_name", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductNameIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_name", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductNameNotIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByProductName(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_name "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductBrand(p mysql.Predicate, value string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_brand", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductBrandIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_brand", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductBrandNotIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_brand", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByProductBrand(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_brand "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductSn(p mysql.Predicate, value string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_sn", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductSnIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_sn", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductSnNotIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_sn", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByProductSn(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_sn "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductPrice(p mysql.Predicate, value float64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_price", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductPriceIn(value []float64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_price", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductPriceNotIn(value []float64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_price", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByProductPrice(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_price "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductQuantity(p mysql.Predicate, value int32) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_quantity", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductQuantityIn(value []int32) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_quantity", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductQuantityNotIn(value []int32) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_quantity", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByProductQuantity(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_quantity "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductSkuId(p mysql.Predicate, value int64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_sku_id", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductSkuIdIn(value []int64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_sku_id", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductSkuIdNotIn(value []int64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_sku_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByProductSkuId(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_sku_id "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductSkuCode(p mysql.Predicate, value string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_sku_code", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductSkuCodeIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_sku_code", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductSkuCodeNotIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_sku_code", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByProductSkuCode(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_sku_code "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductCategoryId(p mysql.Predicate, value int64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_category_id", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductCategoryIdIn(value []int64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_category_id", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductCategoryIdNotIn(value []int64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_category_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByProductCategoryId(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_category_id "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WherePromotionName(p mysql.Predicate, value string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "promotion_name", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WherePromotionNameIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "promotion_name", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WherePromotionNameNotIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "promotion_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByPromotionName(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "promotion_name "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WherePromotionAmount(p mysql.Predicate, value float64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "promotion_amount", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WherePromotionAmountIn(value []float64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "promotion_amount", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WherePromotionAmountNotIn(value []float64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "promotion_amount", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByPromotionAmount(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "promotion_amount "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereCouponAmount(p mysql.Predicate, value float64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "coupon_amount", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereCouponAmountIn(value []float64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "coupon_amount", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereCouponAmountNotIn(value []float64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "coupon_amount", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByCouponAmount(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "coupon_amount "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereIntegrationAmount(p mysql.Predicate, value float64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "integration_amount", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereIntegrationAmountIn(value []float64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "integration_amount", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereIntegrationAmountNotIn(value []float64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "integration_amount", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByIntegrationAmount(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "integration_amount "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereRealAmount(p mysql.Predicate, value float64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "real_amount", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereRealAmountIn(value []float64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "real_amount", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereRealAmountNotIn(value []float64) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "real_amount", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByRealAmount(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "real_amount "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereGiftIntegration(p mysql.Predicate, value int32) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "gift_integration", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereGiftIntegrationIn(value []int32) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "gift_integration", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereGiftIntegrationNotIn(value []int32) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "gift_integration", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByGiftIntegration(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "gift_integration "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereGiftGrowth(p mysql.Predicate, value int32) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "gift_growth", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereGiftGrowthIn(value []int32) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "gift_growth", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereGiftGrowthNotIn(value []int32) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "gift_growth", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByGiftGrowth(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "gift_growth "+order)
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductAttr(p mysql.Predicate, value string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attr", p),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductAttrIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attr", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) WhereProductAttrNotIn(value []string) *omsOrderItemQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attr", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderItemQueryBuilder) OrderByProductAttr(asc bool) *omsOrderItemQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_attr "+order)
	return qb
}
