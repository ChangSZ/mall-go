///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package ums_role_resource_relation

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
)

func NewModel() *UmsRoleResourceRelation {
	return new(UmsRoleResourceRelation)
}

func NewQueryBuilder() *umsRoleResourceRelationQueryBuilder {
	return new(umsRoleResourceRelationQueryBuilder)
}

func (t *UmsRoleResourceRelation) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type umsRoleResourceRelationQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *umsRoleResourceRelationQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *umsRoleResourceRelationQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&UmsRoleResourceRelation{})

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

func (qb *umsRoleResourceRelationQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&UmsRoleResourceRelation{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *umsRoleResourceRelationQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&UmsRoleResourceRelation{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *umsRoleResourceRelationQueryBuilder) First(db *gorm.DB) (*UmsRoleResourceRelation, error) {
	ret := &UmsRoleResourceRelation{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *umsRoleResourceRelationQueryBuilder) QueryOne(db *gorm.DB) (*UmsRoleResourceRelation, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *umsRoleResourceRelationQueryBuilder) QueryAll(db *gorm.DB) ([]*UmsRoleResourceRelation, error) {
	var ret []*UmsRoleResourceRelation
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *umsRoleResourceRelationQueryBuilder) Limit(limit int) *umsRoleResourceRelationQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *umsRoleResourceRelationQueryBuilder) Offset(offset int) *umsRoleResourceRelationQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *umsRoleResourceRelationQueryBuilder) WhereId(p mysql.Predicate, value int64) *umsRoleResourceRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *umsRoleResourceRelationQueryBuilder) WhereIdIn(value []int64) *umsRoleResourceRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *umsRoleResourceRelationQueryBuilder) WhereIdNotIn(value []int64) *umsRoleResourceRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsRoleResourceRelationQueryBuilder) OrderById(asc bool) *umsRoleResourceRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *umsRoleResourceRelationQueryBuilder) WhereRoleId(p mysql.Predicate, value int64) *umsRoleResourceRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "role_id", p),
		value,
	})
	return qb
}

func (qb *umsRoleResourceRelationQueryBuilder) WhereRoleIdIn(value []int64) *umsRoleResourceRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "role_id", "IN"),
		value,
	})
	return qb
}

func (qb *umsRoleResourceRelationQueryBuilder) WhereRoleIdNotIn(value []int64) *umsRoleResourceRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "role_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsRoleResourceRelationQueryBuilder) OrderByRoleId(asc bool) *umsRoleResourceRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "role_id "+order)
	return qb
}

func (qb *umsRoleResourceRelationQueryBuilder) WhereResourceId(p mysql.Predicate, value int64) *umsRoleResourceRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "resource_id", p),
		value,
	})
	return qb
}

func (qb *umsRoleResourceRelationQueryBuilder) WhereResourceIdIn(value []int64) *umsRoleResourceRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "resource_id", "IN"),
		value,
	})
	return qb
}

func (qb *umsRoleResourceRelationQueryBuilder) WhereResourceIdNotIn(value []int64) *umsRoleResourceRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "resource_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsRoleResourceRelationQueryBuilder) OrderByResourceId(asc bool) *umsRoleResourceRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "resource_id "+order)
	return qb
}
