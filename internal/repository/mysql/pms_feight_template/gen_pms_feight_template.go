///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package pms_feight_template

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *PmsFeightTemplate {
	return new(PmsFeightTemplate)
}

func NewQueryBuilder() *pmsFeightTemplateQueryBuilder {
	return new(pmsFeightTemplateQueryBuilder)
}

func (t *PmsFeightTemplate) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type pmsFeightTemplateQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *pmsFeightTemplateQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *pmsFeightTemplateQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&PmsFeightTemplate{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *pmsFeightTemplateQueryBuilder) Update(db *gorm.DB, data *PmsFeightTemplate) (cnt int64, err error) {
	db = db.Model(&PmsFeightTemplate{})

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

func (qb *pmsFeightTemplateQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&PmsFeightTemplate{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *pmsFeightTemplateQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&PmsFeightTemplate{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *pmsFeightTemplateQueryBuilder) First(db *gorm.DB) (*PmsFeightTemplate, error) {
	ret := &PmsFeightTemplate{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *pmsFeightTemplateQueryBuilder) QueryOne(db *gorm.DB) (*PmsFeightTemplate, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *pmsFeightTemplateQueryBuilder) QueryAll(db *gorm.DB) ([]*PmsFeightTemplate, error) {
	var ret []*PmsFeightTemplate
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *pmsFeightTemplateQueryBuilder) Limit(limit int) *pmsFeightTemplateQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) Offset(offset int) *pmsFeightTemplateQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereId(p mysql.Predicate, value int64) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereIdIn(value []int64) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereIdNotIn(value []int64) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) OrderById(asc bool) *pmsFeightTemplateQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereName(p mysql.Predicate, value string) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereNameIn(value []string) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereNameNotIn(value []string) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) OrderByName(asc bool) *pmsFeightTemplateQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereChargeType(p mysql.Predicate, value int32) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "charge_type", p),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereChargeTypeIn(value []int32) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "charge_type", "IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereChargeTypeNotIn(value []int32) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "charge_type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) OrderByChargeType(asc bool) *pmsFeightTemplateQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "charge_type "+order)
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereFirstWeight(p mysql.Predicate, value float64) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "first_weight", p),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereFirstWeightIn(value []float64) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "first_weight", "IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereFirstWeightNotIn(value []float64) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "first_weight", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) OrderByFirstWeight(asc bool) *pmsFeightTemplateQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "first_weight "+order)
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereFirstFee(p mysql.Predicate, value float64) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "first_fee", p),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereFirstFeeIn(value []float64) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "first_fee", "IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereFirstFeeNotIn(value []float64) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "first_fee", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) OrderByFirstFee(asc bool) *pmsFeightTemplateQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "first_fee "+order)
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereContinueWeight(p mysql.Predicate, value float64) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "continue_weight", p),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereContinueWeightIn(value []float64) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "continue_weight", "IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereContinueWeightNotIn(value []float64) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "continue_weight", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) OrderByContinueWeight(asc bool) *pmsFeightTemplateQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "continue_weight "+order)
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereContinmeFee(p mysql.Predicate, value float64) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "continme_fee", p),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereContinmeFeeIn(value []float64) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "continme_fee", "IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereContinmeFeeNotIn(value []float64) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "continme_fee", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) OrderByContinmeFee(asc bool) *pmsFeightTemplateQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "continme_fee "+order)
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereDest(p mysql.Predicate, value string) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "dest", p),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereDestIn(value []string) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "dest", "IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) WhereDestNotIn(value []string) *pmsFeightTemplateQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "dest", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsFeightTemplateQueryBuilder) OrderByDest(asc bool) *pmsFeightTemplateQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "dest "+order)
	return qb
}