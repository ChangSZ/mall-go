///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package pms_sku_stock

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
)

func NewModel() *PmsSkuStock {
	return new(PmsSkuStock)
}

func NewQueryBuilder() *pmsSkuStockQueryBuilder {
	return new(pmsSkuStockQueryBuilder)
}

func (t *PmsSkuStock) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type pmsSkuStockQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *pmsSkuStockQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *pmsSkuStockQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&PmsSkuStock{})

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

func (qb *pmsSkuStockQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&PmsSkuStock{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *pmsSkuStockQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&PmsSkuStock{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *pmsSkuStockQueryBuilder) First(db *gorm.DB) (*PmsSkuStock, error) {
	ret := &PmsSkuStock{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *pmsSkuStockQueryBuilder) QueryOne(db *gorm.DB) (*PmsSkuStock, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *pmsSkuStockQueryBuilder) QueryAll(db *gorm.DB) ([]*PmsSkuStock, error) {
	var ret []*PmsSkuStock
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *pmsSkuStockQueryBuilder) Limit(limit int) *pmsSkuStockQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *pmsSkuStockQueryBuilder) Offset(offset int) *pmsSkuStockQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereId(p mysql.Predicate, value int64) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereIdIn(value []int64) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereIdNotIn(value []int64) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) OrderById(asc bool) *pmsSkuStockQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereProductId(p mysql.Predicate, value int64) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", p),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereProductIdIn(value []int64) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereProductIdNotIn(value []int64) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) OrderByProductId(asc bool) *pmsSkuStockQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_id "+order)
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereSkuCode(p mysql.Predicate, value string) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sku_code", p),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereSkuCodeIn(value []string) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sku_code", "IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereSkuCodeNotIn(value []string) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sku_code", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) OrderBySkuCode(asc bool) *pmsSkuStockQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sku_code "+order)
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WherePrice(p mysql.Predicate, value float64) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "price", p),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WherePriceIn(value []float64) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "price", "IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WherePriceNotIn(value []float64) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "price", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) OrderByPrice(asc bool) *pmsSkuStockQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "price "+order)
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereStock(p mysql.Predicate, value int32) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "stock", p),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereStockIn(value []int32) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "stock", "IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereStockNotIn(value []int32) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "stock", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) OrderByStock(asc bool) *pmsSkuStockQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "stock "+order)
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereLowStock(p mysql.Predicate, value int32) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "low_stock", p),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereLowStockIn(value []int32) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "low_stock", "IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereLowStockNotIn(value []int32) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "low_stock", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) OrderByLowStock(asc bool) *pmsSkuStockQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "low_stock "+order)
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WherePic(p mysql.Predicate, value string) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pic", p),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WherePicIn(value []string) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pic", "IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WherePicNotIn(value []string) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pic", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) OrderByPic(asc bool) *pmsSkuStockQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "pic "+order)
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereSale(p mysql.Predicate, value int32) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sale", p),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereSaleIn(value []int32) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sale", "IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereSaleNotIn(value []int32) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sale", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) OrderBySale(asc bool) *pmsSkuStockQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sale "+order)
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WherePromotionPrice(p mysql.Predicate, value float64) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "promotion_price", p),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WherePromotionPriceIn(value []float64) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "promotion_price", "IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WherePromotionPriceNotIn(value []float64) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "promotion_price", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) OrderByPromotionPrice(asc bool) *pmsSkuStockQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "promotion_price "+order)
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereLockStock(p mysql.Predicate, value int32) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "lock_stock", p),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereLockStockIn(value []int32) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "lock_stock", "IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereLockStockNotIn(value []int32) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "lock_stock", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) OrderByLockStock(asc bool) *pmsSkuStockQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "lock_stock "+order)
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereSpData(p mysql.Predicate, value string) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sp_data", p),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereSpDataIn(value []string) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sp_data", "IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) WhereSpDataNotIn(value []string) *pmsSkuStockQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sp_data", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsSkuStockQueryBuilder) OrderBySpData(asc bool) *pmsSkuStockQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sp_data "+order)
	return qb
}
