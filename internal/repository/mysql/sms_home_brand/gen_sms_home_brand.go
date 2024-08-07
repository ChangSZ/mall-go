///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package sms_home_brand

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
)

func NewModel() *SmsHomeBrand {
	return new(SmsHomeBrand)
}

func NewQueryBuilder() *smsHomeBrandQueryBuilder {
	return new(smsHomeBrandQueryBuilder)
}

func (t *SmsHomeBrand) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type smsHomeBrandQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *smsHomeBrandQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *smsHomeBrandQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&SmsHomeBrand{})

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

func (qb *smsHomeBrandQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&SmsHomeBrand{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *smsHomeBrandQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&SmsHomeBrand{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *smsHomeBrandQueryBuilder) First(db *gorm.DB) (*SmsHomeBrand, error) {
	ret := &SmsHomeBrand{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *smsHomeBrandQueryBuilder) QueryOne(db *gorm.DB) (*SmsHomeBrand, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *smsHomeBrandQueryBuilder) QueryAll(db *gorm.DB) ([]*SmsHomeBrand, error) {
	var ret []*SmsHomeBrand
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *smsHomeBrandQueryBuilder) Limit(limit int) *smsHomeBrandQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *smsHomeBrandQueryBuilder) Offset(offset int) *smsHomeBrandQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *smsHomeBrandQueryBuilder) WhereId(p mysql.Predicate, value int64) *smsHomeBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *smsHomeBrandQueryBuilder) WhereIdIn(value []int64) *smsHomeBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeBrandQueryBuilder) WhereIdNotIn(value []int64) *smsHomeBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeBrandQueryBuilder) OrderById(asc bool) *smsHomeBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *smsHomeBrandQueryBuilder) WhereBrandId(p mysql.Predicate, value int64) *smsHomeBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "brand_id", p),
		value,
	})
	return qb
}

func (qb *smsHomeBrandQueryBuilder) WhereBrandIdIn(value []int64) *smsHomeBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "brand_id", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeBrandQueryBuilder) WhereBrandIdNotIn(value []int64) *smsHomeBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "brand_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeBrandQueryBuilder) OrderByBrandId(asc bool) *smsHomeBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "brand_id "+order)
	return qb
}

func (qb *smsHomeBrandQueryBuilder) WhereBrandName(p mysql.Predicate, value string) *smsHomeBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "brand_name", p),
		value,
	})
	return qb
}

func (qb *smsHomeBrandQueryBuilder) WhereBrandNameIn(value []string) *smsHomeBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "brand_name", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeBrandQueryBuilder) WhereBrandNameNotIn(value []string) *smsHomeBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "brand_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeBrandQueryBuilder) OrderByBrandName(asc bool) *smsHomeBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "brand_name "+order)
	return qb
}

func (qb *smsHomeBrandQueryBuilder) WhereRecommendStatus(p mysql.Predicate, value int32) *smsHomeBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "recommend_status", p),
		value,
	})
	return qb
}

func (qb *smsHomeBrandQueryBuilder) WhereRecommendStatusIn(value []int32) *smsHomeBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "recommend_status", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeBrandQueryBuilder) WhereRecommendStatusNotIn(value []int32) *smsHomeBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "recommend_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeBrandQueryBuilder) OrderByRecommendStatus(asc bool) *smsHomeBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "recommend_status "+order)
	return qb
}

func (qb *smsHomeBrandQueryBuilder) WhereSort(p mysql.Predicate, value int32) *smsHomeBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", p),
		value,
	})
	return qb
}

func (qb *smsHomeBrandQueryBuilder) WhereSortIn(value []int32) *smsHomeBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeBrandQueryBuilder) WhereSortNotIn(value []int32) *smsHomeBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeBrandQueryBuilder) OrderBySort(asc bool) *smsHomeBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sort "+order)
	return qb
}
