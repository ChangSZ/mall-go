///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package sms_flash_promotion_product_relation

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
)

func NewModel() *SmsFlashPromotionProductRelation {
	return new(SmsFlashPromotionProductRelation)
}

func NewQueryBuilder() *smsFlashPromotionProductRelationQueryBuilder {
	return new(smsFlashPromotionProductRelationQueryBuilder)
}

func (t *SmsFlashPromotionProductRelation) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type smsFlashPromotionProductRelationQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *smsFlashPromotionProductRelationQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&SmsFlashPromotionProductRelation{})

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

func (qb *smsFlashPromotionProductRelationQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&SmsFlashPromotionProductRelation{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&SmsFlashPromotionProductRelation{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) First(db *gorm.DB) (*SmsFlashPromotionProductRelation, error) {
	ret := &SmsFlashPromotionProductRelation{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) QueryOne(db *gorm.DB) (*SmsFlashPromotionProductRelation, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) QueryAll(db *gorm.DB) ([]*SmsFlashPromotionProductRelation, error) {
	var ret []*SmsFlashPromotionProductRelation
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) Limit(limit int) *smsFlashPromotionProductRelationQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) Offset(offset int) *smsFlashPromotionProductRelationQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereId(p mysql.Predicate, value int64) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereIdIn(value []int64) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereIdNotIn(value []int64) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) OrderById(asc bool) *smsFlashPromotionProductRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereFlashPromotionId(p mysql.Predicate, value int64) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_promotion_id", p),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereFlashPromotionIdIn(value []int64) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_promotion_id", "IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereFlashPromotionIdNotIn(value []int64) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_promotion_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) OrderByFlashPromotionId(asc bool) *smsFlashPromotionProductRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "flash_promotion_id "+order)
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereFlashPromotionSessionId(p mysql.Predicate, value int64) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_promotion_session_id", p),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereFlashPromotionSessionIdIn(value []int64) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_promotion_session_id", "IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereFlashPromotionSessionIdNotIn(value []int64) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_promotion_session_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) OrderByFlashPromotionSessionId(asc bool) *smsFlashPromotionProductRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "flash_promotion_session_id "+order)
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereProductId(p mysql.Predicate, value int64) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", p),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereProductIdIn(value []int64) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereProductIdNotIn(value []int64) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) OrderByProductId(asc bool) *smsFlashPromotionProductRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_id "+order)
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereFlashPromotionPrice(p mysql.Predicate, value float64) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_promotion_price", p),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereFlashPromotionPriceIn(value []float64) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_promotion_price", "IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereFlashPromotionPriceNotIn(value []float64) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_promotion_price", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) OrderByFlashPromotionPrice(asc bool) *smsFlashPromotionProductRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "flash_promotion_price "+order)
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereFlashPromotionCount(p mysql.Predicate, value int32) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_promotion_count", p),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereFlashPromotionCountIn(value []int32) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_promotion_count", "IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereFlashPromotionCountNotIn(value []int32) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_promotion_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) OrderByFlashPromotionCount(asc bool) *smsFlashPromotionProductRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "flash_promotion_count "+order)
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereFlashPromotionLimit(p mysql.Predicate, value int32) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_promotion_limit", p),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereFlashPromotionLimitIn(value []int32) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_promotion_limit", "IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereFlashPromotionLimitNotIn(value []int32) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_promotion_limit", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) OrderByFlashPromotionLimit(asc bool) *smsFlashPromotionProductRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "flash_promotion_limit "+order)
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereSort(p mysql.Predicate, value int32) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", p),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereSortIn(value []int32) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) WhereSortNotIn(value []int32) *smsFlashPromotionProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionProductRelationQueryBuilder) OrderBySort(asc bool) *smsFlashPromotionProductRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sort "+order)
	return qb
}
