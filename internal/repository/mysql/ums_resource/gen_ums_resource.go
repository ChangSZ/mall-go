///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package ums_resource

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
)

func NewModel() *UmsResource {
	return new(UmsResource)
}

func NewQueryBuilder() *umsResourceQueryBuilder {
	return new(umsResourceQueryBuilder)
}

func (t *UmsResource) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type umsResourceQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *umsResourceQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *umsResourceQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&UmsResource{})

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

func (qb *umsResourceQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&UmsResource{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *umsResourceQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&UmsResource{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *umsResourceQueryBuilder) First(db *gorm.DB) (*UmsResource, error) {
	ret := &UmsResource{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *umsResourceQueryBuilder) QueryOne(db *gorm.DB) (*UmsResource, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *umsResourceQueryBuilder) QueryAll(db *gorm.DB) ([]*UmsResource, error) {
	var ret []*UmsResource
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *umsResourceQueryBuilder) Limit(limit int) *umsResourceQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *umsResourceQueryBuilder) Offset(offset int) *umsResourceQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *umsResourceQueryBuilder) WhereId(p mysql.Predicate, value int64) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) WhereIdIn(value []int64) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) WhereIdNotIn(value []int64) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) OrderById(asc bool) *umsResourceQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *umsResourceQueryBuilder) WhereCreateTime(p mysql.Predicate, value time.Time) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) WhereCreateTimeIn(value []time.Time) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) WhereCreateTimeNotIn(value []time.Time) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) OrderByCreateTime(asc bool) *umsResourceQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *umsResourceQueryBuilder) WhereName(p mysql.Predicate, value string) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) WhereNameIn(value []string) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) WhereNameNotIn(value []string) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) OrderByName(asc bool) *umsResourceQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *umsResourceQueryBuilder) WhereUrl(p mysql.Predicate, value string) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "url", p),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) WhereUrlIn(value []string) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "url", "IN"),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) WhereUrlNotIn(value []string) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "url", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) OrderByUrl(asc bool) *umsResourceQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "url "+order)
	return qb
}

func (qb *umsResourceQueryBuilder) WhereDescription(p mysql.Predicate, value string) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", p),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) WhereDescriptionIn(value []string) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", "IN"),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) WhereDescriptionNotIn(value []string) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) OrderByDescription(asc bool) *umsResourceQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "description "+order)
	return qb
}

func (qb *umsResourceQueryBuilder) WhereCategoryId(p mysql.Predicate, value int64) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", p),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) WhereCategoryIdIn(value []int64) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", "IN"),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) WhereCategoryIdNotIn(value []int64) *umsResourceQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsResourceQueryBuilder) OrderByCategoryId(asc bool) *umsResourceQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "category_id "+order)
	return qb
}
