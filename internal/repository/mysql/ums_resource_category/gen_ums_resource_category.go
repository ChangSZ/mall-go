///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package ums_resource_category

import (
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *UmsResourceCategory {
	return new(UmsResourceCategory)
}

func NewQueryBuilder() *umsResourceCategoryQueryBuilder {
	return new(umsResourceCategoryQueryBuilder)
}

func (t *UmsResourceCategory) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type umsResourceCategoryQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *umsResourceCategoryQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *umsResourceCategoryQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&UmsResourceCategory{})

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

func (qb *umsResourceCategoryQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&UmsResourceCategory{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *umsResourceCategoryQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&UmsResourceCategory{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *umsResourceCategoryQueryBuilder) First(db *gorm.DB) (*UmsResourceCategory, error) {
	ret := &UmsResourceCategory{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *umsResourceCategoryQueryBuilder) QueryOne(db *gorm.DB) (*UmsResourceCategory, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *umsResourceCategoryQueryBuilder) QueryAll(db *gorm.DB) ([]*UmsResourceCategory, error) {
	var ret []*UmsResourceCategory
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *umsResourceCategoryQueryBuilder) Limit(limit int) *umsResourceCategoryQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) Offset(offset int) *umsResourceCategoryQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) WhereId(p mysql.Predicate, value int64) *umsResourceCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) WhereIdIn(value []int64) *umsResourceCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) WhereIdNotIn(value []int64) *umsResourceCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) OrderById(asc bool) *umsResourceCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) WhereCreateTime(p mysql.Predicate, value time.Time) *umsResourceCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) WhereCreateTimeIn(value []time.Time) *umsResourceCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) WhereCreateTimeNotIn(value []time.Time) *umsResourceCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) OrderByCreateTime(asc bool) *umsResourceCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) WhereName(p mysql.Predicate, value string) *umsResourceCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) WhereNameIn(value []string) *umsResourceCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) WhereNameNotIn(value []string) *umsResourceCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) OrderByName(asc bool) *umsResourceCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) WhereSort(p mysql.Predicate, value int32) *umsResourceCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", p),
		value,
	})
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) WhereSortIn(value []int32) *umsResourceCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "IN"),
		value,
	})
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) WhereSortNotIn(value []int32) *umsResourceCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsResourceCategoryQueryBuilder) OrderBySort(asc bool) *umsResourceCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sort "+order)
	return qb
}
