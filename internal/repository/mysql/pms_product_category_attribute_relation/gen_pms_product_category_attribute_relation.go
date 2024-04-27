///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package pms_product_category_attribute_relation

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *PmsProductCategoryAttributeRelation {
	return new(PmsProductCategoryAttributeRelation)
}

func NewQueryBuilder() *pmsProductCategoryAttributeRelationQueryBuilder {
	return new(pmsProductCategoryAttributeRelationQueryBuilder)
}

func (t *PmsProductCategoryAttributeRelation) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type pmsProductCategoryAttributeRelationQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&PmsProductCategoryAttributeRelation{})

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

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&PmsProductCategoryAttributeRelation{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&PmsProductCategoryAttributeRelation{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) First(db *gorm.DB) (*PmsProductCategoryAttributeRelation, error) {
	ret := &PmsProductCategoryAttributeRelation{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) QueryOne(db *gorm.DB) (*PmsProductCategoryAttributeRelation, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) QueryAll(db *gorm.DB) ([]*PmsProductCategoryAttributeRelation, error) {
	var ret []*PmsProductCategoryAttributeRelation
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) Limit(limit int) *pmsProductCategoryAttributeRelationQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) Offset(offset int) *pmsProductCategoryAttributeRelationQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) WhereId(p mysql.Predicate, value int64) *pmsProductCategoryAttributeRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) WhereIdIn(value []int64) *pmsProductCategoryAttributeRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) WhereIdNotIn(value []int64) *pmsProductCategoryAttributeRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) OrderById(asc bool) *pmsProductCategoryAttributeRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) WhereProductCategoryId(p mysql.Predicate, value int64) *pmsProductCategoryAttributeRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_category_id", p),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) WhereProductCategoryIdIn(value []int64) *pmsProductCategoryAttributeRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_category_id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) WhereProductCategoryIdNotIn(value []int64) *pmsProductCategoryAttributeRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_category_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) OrderByProductCategoryId(asc bool) *pmsProductCategoryAttributeRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_category_id "+order)
	return qb
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) WhereProductAttributeId(p mysql.Predicate, value int64) *pmsProductCategoryAttributeRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attribute_id", p),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) WhereProductAttributeIdIn(value []int64) *pmsProductCategoryAttributeRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attribute_id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) WhereProductAttributeIdNotIn(value []int64) *pmsProductCategoryAttributeRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attribute_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryAttributeRelationQueryBuilder) OrderByProductAttributeId(asc bool) *pmsProductCategoryAttributeRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_attribute_id "+order)
	return qb
}
