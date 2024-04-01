///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package ums_member_task

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *UmsMemberTask {
	return new(UmsMemberTask)
}

func NewQueryBuilder() *umsMemberTaskQueryBuilder {
	return new(umsMemberTaskQueryBuilder)
}

func (t *UmsMemberTask) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type umsMemberTaskQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *umsMemberTaskQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *umsMemberTaskQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&UmsMemberTask{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *umsMemberTaskQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&UmsMemberTask{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *umsMemberTaskQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&UmsMemberTask{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *umsMemberTaskQueryBuilder) First(db *gorm.DB) (*UmsMemberTask, error) {
	ret := &UmsMemberTask{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *umsMemberTaskQueryBuilder) QueryOne(db *gorm.DB) (*UmsMemberTask, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *umsMemberTaskQueryBuilder) QueryAll(db *gorm.DB) ([]*UmsMemberTask, error) {
	var ret []*UmsMemberTask
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *umsMemberTaskQueryBuilder) Limit(limit int) *umsMemberTaskQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *umsMemberTaskQueryBuilder) Offset(offset int) *umsMemberTaskQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *umsMemberTaskQueryBuilder) WhereId(p mysql.Predicate, value int64) *umsMemberTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *umsMemberTaskQueryBuilder) WhereIdIn(value []int64) *umsMemberTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberTaskQueryBuilder) WhereIdNotIn(value []int64) *umsMemberTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberTaskQueryBuilder) OrderById(asc bool) *umsMemberTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *umsMemberTaskQueryBuilder) WhereName(p mysql.Predicate, value string) *umsMemberTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *umsMemberTaskQueryBuilder) WhereNameIn(value []string) *umsMemberTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberTaskQueryBuilder) WhereNameNotIn(value []string) *umsMemberTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberTaskQueryBuilder) OrderByName(asc bool) *umsMemberTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *umsMemberTaskQueryBuilder) WhereGrowth(p mysql.Predicate, value int32) *umsMemberTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "growth", p),
		value,
	})
	return qb
}

func (qb *umsMemberTaskQueryBuilder) WhereGrowthIn(value []int32) *umsMemberTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "growth", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberTaskQueryBuilder) WhereGrowthNotIn(value []int32) *umsMemberTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "growth", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberTaskQueryBuilder) OrderByGrowth(asc bool) *umsMemberTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "growth "+order)
	return qb
}

func (qb *umsMemberTaskQueryBuilder) WhereIntergration(p mysql.Predicate, value int32) *umsMemberTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "intergration", p),
		value,
	})
	return qb
}

func (qb *umsMemberTaskQueryBuilder) WhereIntergrationIn(value []int32) *umsMemberTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "intergration", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberTaskQueryBuilder) WhereIntergrationNotIn(value []int32) *umsMemberTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "intergration", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberTaskQueryBuilder) OrderByIntergration(asc bool) *umsMemberTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "intergration "+order)
	return qb
}

func (qb *umsMemberTaskQueryBuilder) WhereType(p mysql.Predicate, value int32) *umsMemberTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", p),
		value,
	})
	return qb
}

func (qb *umsMemberTaskQueryBuilder) WhereTypeIn(value []int32) *umsMemberTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberTaskQueryBuilder) WhereTypeNotIn(value []int32) *umsMemberTaskQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberTaskQueryBuilder) OrderByType(asc bool) *umsMemberTaskQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "type "+order)
	return qb
}
