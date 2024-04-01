///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package ums_member_product_category_relation

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *UmsMemberProductCategoryRelation {
	return new(UmsMemberProductCategoryRelation)
}

func NewQueryBuilder() *umsMemberProductCategoryRelationQueryBuilder {
	return new(umsMemberProductCategoryRelationQueryBuilder)
}

func (t *UmsMemberProductCategoryRelation) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type umsMemberProductCategoryRelationQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *umsMemberProductCategoryRelationQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&UmsMemberProductCategoryRelation{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&UmsMemberProductCategoryRelation{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&UmsMemberProductCategoryRelation{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) First(db *gorm.DB) (*UmsMemberProductCategoryRelation, error) {
	ret := &UmsMemberProductCategoryRelation{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) QueryOne(db *gorm.DB) (*UmsMemberProductCategoryRelation, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) QueryAll(db *gorm.DB) ([]*UmsMemberProductCategoryRelation, error) {
	var ret []*UmsMemberProductCategoryRelation
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) Limit(limit int) *umsMemberProductCategoryRelationQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) Offset(offset int) *umsMemberProductCategoryRelationQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) WhereId(p mysql.Predicate, value int64) *umsMemberProductCategoryRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) WhereIdIn(value []int64) *umsMemberProductCategoryRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) WhereIdNotIn(value []int64) *umsMemberProductCategoryRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) OrderById(asc bool) *umsMemberProductCategoryRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) WhereMemberId(p mysql.Predicate, value int64) *umsMemberProductCategoryRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_id", p),
		value,
	})
	return qb
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) WhereMemberIdIn(value []int64) *umsMemberProductCategoryRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_id", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) WhereMemberIdNotIn(value []int64) *umsMemberProductCategoryRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) OrderByMemberId(asc bool) *umsMemberProductCategoryRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "member_id "+order)
	return qb
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) WhereProductCategoryId(p mysql.Predicate, value int64) *umsMemberProductCategoryRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_category_id", p),
		value,
	})
	return qb
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) WhereProductCategoryIdIn(value []int64) *umsMemberProductCategoryRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_category_id", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) WhereProductCategoryIdNotIn(value []int64) *umsMemberProductCategoryRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_category_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberProductCategoryRelationQueryBuilder) OrderByProductCategoryId(asc bool) *umsMemberProductCategoryRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_category_id "+order)
	return qb
}
