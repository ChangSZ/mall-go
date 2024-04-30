///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package cms_subject_product_relation

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *CmsSubjectProductRelation {
	return new(CmsSubjectProductRelation)
}

func NewQueryBuilder() *cmsSubjectProductRelationQueryBuilder {
	return new(cmsSubjectProductRelationQueryBuilder)
}

func (t *CmsSubjectProductRelation) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type cmsSubjectProductRelationQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *cmsSubjectProductRelationQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *cmsSubjectProductRelationQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&CmsSubjectProductRelation{})

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

func (qb *cmsSubjectProductRelationQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&CmsSubjectProductRelation{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *cmsSubjectProductRelationQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&CmsSubjectProductRelation{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *cmsSubjectProductRelationQueryBuilder) First(db *gorm.DB) (*CmsSubjectProductRelation, error) {
	ret := &CmsSubjectProductRelation{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *cmsSubjectProductRelationQueryBuilder) QueryOne(db *gorm.DB) (*CmsSubjectProductRelation, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *cmsSubjectProductRelationQueryBuilder) QueryAll(db *gorm.DB) ([]*CmsSubjectProductRelation, error) {
	var ret []*CmsSubjectProductRelation
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *cmsSubjectProductRelationQueryBuilder) Limit(limit int) *cmsSubjectProductRelationQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *cmsSubjectProductRelationQueryBuilder) Offset(offset int) *cmsSubjectProductRelationQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *cmsSubjectProductRelationQueryBuilder) WhereId(p mysql.Predicate, value int64) *cmsSubjectProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectProductRelationQueryBuilder) WhereIdIn(value []int64) *cmsSubjectProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectProductRelationQueryBuilder) WhereIdNotIn(value []int64) *cmsSubjectProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectProductRelationQueryBuilder) OrderById(asc bool) *cmsSubjectProductRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *cmsSubjectProductRelationQueryBuilder) WhereSubjectId(p mysql.Predicate, value int64) *cmsSubjectProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "subject_id", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectProductRelationQueryBuilder) WhereSubjectIdIn(value []int64) *cmsSubjectProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "subject_id", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectProductRelationQueryBuilder) WhereSubjectIdNotIn(value []int64) *cmsSubjectProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "subject_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectProductRelationQueryBuilder) OrderBySubjectId(asc bool) *cmsSubjectProductRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "subject_id "+order)
	return qb
}

func (qb *cmsSubjectProductRelationQueryBuilder) WhereProductId(p mysql.Predicate, value int64) *cmsSubjectProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectProductRelationQueryBuilder) WhereProductIdIn(value []int64) *cmsSubjectProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectProductRelationQueryBuilder) WhereProductIdNotIn(value []int64) *cmsSubjectProductRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectProductRelationQueryBuilder) OrderByProductId(asc bool) *cmsSubjectProductRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_id "+order)
	return qb
}