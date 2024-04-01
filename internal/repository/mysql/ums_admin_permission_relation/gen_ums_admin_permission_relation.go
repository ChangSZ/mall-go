///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package ums_admin_permission_relation

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *UmsAdminPermissionRelation {
	return new(UmsAdminPermissionRelation)
}

func NewQueryBuilder() *umsAdminPermissionRelationQueryBuilder {
	return new(umsAdminPermissionRelationQueryBuilder)
}

func (t *UmsAdminPermissionRelation) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type umsAdminPermissionRelationQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *umsAdminPermissionRelationQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *umsAdminPermissionRelationQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&UmsAdminPermissionRelation{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *umsAdminPermissionRelationQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&UmsAdminPermissionRelation{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *umsAdminPermissionRelationQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&UmsAdminPermissionRelation{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *umsAdminPermissionRelationQueryBuilder) First(db *gorm.DB) (*UmsAdminPermissionRelation, error) {
	ret := &UmsAdminPermissionRelation{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *umsAdminPermissionRelationQueryBuilder) QueryOne(db *gorm.DB) (*UmsAdminPermissionRelation, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *umsAdminPermissionRelationQueryBuilder) QueryAll(db *gorm.DB) ([]*UmsAdminPermissionRelation, error) {
	var ret []*UmsAdminPermissionRelation
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *umsAdminPermissionRelationQueryBuilder) Limit(limit int) *umsAdminPermissionRelationQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) Offset(offset int) *umsAdminPermissionRelationQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) WhereId(p mysql.Predicate, value int64) *umsAdminPermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) WhereIdIn(value []int64) *umsAdminPermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) WhereIdNotIn(value []int64) *umsAdminPermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) OrderById(asc bool) *umsAdminPermissionRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) WhereAdminId(p mysql.Predicate, value int64) *umsAdminPermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "admin_id", p),
		value,
	})
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) WhereAdminIdIn(value []int64) *umsAdminPermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "admin_id", "IN"),
		value,
	})
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) WhereAdminIdNotIn(value []int64) *umsAdminPermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "admin_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) OrderByAdminId(asc bool) *umsAdminPermissionRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "admin_id "+order)
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) WherePermissionId(p mysql.Predicate, value int64) *umsAdminPermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "permission_id", p),
		value,
	})
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) WherePermissionIdIn(value []int64) *umsAdminPermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "permission_id", "IN"),
		value,
	})
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) WherePermissionIdNotIn(value []int64) *umsAdminPermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "permission_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) OrderByPermissionId(asc bool) *umsAdminPermissionRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "permission_id "+order)
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) WhereType(p mysql.Predicate, value int32) *umsAdminPermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", p),
		value,
	})
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) WhereTypeIn(value []int32) *umsAdminPermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "IN"),
		value,
	})
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) WhereTypeNotIn(value []int32) *umsAdminPermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsAdminPermissionRelationQueryBuilder) OrderByType(asc bool) *umsAdminPermissionRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "type "+order)
	return qb
}
