///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package ums_admin_role_relation

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *UmsAdminRoleRelation {
	return new(UmsAdminRoleRelation)
}

func NewQueryBuilder() *umsAdminRoleRelationQueryBuilder {
	return new(umsAdminRoleRelationQueryBuilder)
}

func (t *UmsAdminRoleRelation) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type umsAdminRoleRelationQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *umsAdminRoleRelationQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *umsAdminRoleRelationQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&UmsAdminRoleRelation{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *umsAdminRoleRelationQueryBuilder) Update(db *gorm.DB, data *UmsAdminRoleRelation) (cnt int64, err error) {
	db = db.Model(&UmsAdminRoleRelation{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	ret := db.Updates(data)
	err = ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "update err")
	}
	return ret.RowsAffected, nil
}

func (qb *umsAdminRoleRelationQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&UmsAdminRoleRelation{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *umsAdminRoleRelationQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&UmsAdminRoleRelation{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *umsAdminRoleRelationQueryBuilder) First(db *gorm.DB) (*UmsAdminRoleRelation, error) {
	ret := &UmsAdminRoleRelation{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *umsAdminRoleRelationQueryBuilder) QueryOne(db *gorm.DB) (*UmsAdminRoleRelation, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *umsAdminRoleRelationQueryBuilder) QueryAll(db *gorm.DB) ([]*UmsAdminRoleRelation, error) {
	var ret []*UmsAdminRoleRelation
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *umsAdminRoleRelationQueryBuilder) Limit(limit int) *umsAdminRoleRelationQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *umsAdminRoleRelationQueryBuilder) Offset(offset int) *umsAdminRoleRelationQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *umsAdminRoleRelationQueryBuilder) WhereId(p mysql.Predicate, value int64) *umsAdminRoleRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *umsAdminRoleRelationQueryBuilder) WhereIdIn(value []int64) *umsAdminRoleRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *umsAdminRoleRelationQueryBuilder) WhereIdNotIn(value []int64) *umsAdminRoleRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsAdminRoleRelationQueryBuilder) OrderById(asc bool) *umsAdminRoleRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *umsAdminRoleRelationQueryBuilder) WhereAdminId(p mysql.Predicate, value int64) *umsAdminRoleRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "admin_id", p),
		value,
	})
	return qb
}

func (qb *umsAdminRoleRelationQueryBuilder) WhereAdminIdIn(value []int64) *umsAdminRoleRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "admin_id", "IN"),
		value,
	})
	return qb
}

func (qb *umsAdminRoleRelationQueryBuilder) WhereAdminIdNotIn(value []int64) *umsAdminRoleRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "admin_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsAdminRoleRelationQueryBuilder) OrderByAdminId(asc bool) *umsAdminRoleRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "admin_id "+order)
	return qb
}

func (qb *umsAdminRoleRelationQueryBuilder) WhereRoleId(p mysql.Predicate, value int64) *umsAdminRoleRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "role_id", p),
		value,
	})
	return qb
}

func (qb *umsAdminRoleRelationQueryBuilder) WhereRoleIdIn(value []int64) *umsAdminRoleRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "role_id", "IN"),
		value,
	})
	return qb
}

func (qb *umsAdminRoleRelationQueryBuilder) WhereRoleIdNotIn(value []int64) *umsAdminRoleRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "role_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsAdminRoleRelationQueryBuilder) OrderByRoleId(asc bool) *umsAdminRoleRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "role_id "+order)
	return qb
}
