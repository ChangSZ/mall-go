///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package oms_order_operate_history

import (
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *OmsOrderOperateHistory {
	return new(OmsOrderOperateHistory)
}

func NewQueryBuilder() *omsOrderOperateHistoryQueryBuilder {
	return new(omsOrderOperateHistoryQueryBuilder)
}

func (t *OmsOrderOperateHistory) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type omsOrderOperateHistoryQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *omsOrderOperateHistoryQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *omsOrderOperateHistoryQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&OmsOrderOperateHistory{})

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

func (qb *omsOrderOperateHistoryQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&OmsOrderOperateHistory{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *omsOrderOperateHistoryQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&OmsOrderOperateHistory{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *omsOrderOperateHistoryQueryBuilder) First(db *gorm.DB) (*OmsOrderOperateHistory, error) {
	ret := &OmsOrderOperateHistory{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *omsOrderOperateHistoryQueryBuilder) QueryOne(db *gorm.DB) (*OmsOrderOperateHistory, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *omsOrderOperateHistoryQueryBuilder) QueryAll(db *gorm.DB) ([]*OmsOrderOperateHistory, error) {
	var ret []*OmsOrderOperateHistory
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *omsOrderOperateHistoryQueryBuilder) Limit(limit int) *omsOrderOperateHistoryQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) Offset(offset int) *omsOrderOperateHistoryQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereId(p mysql.Predicate, value int64) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereIdIn(value []int64) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereIdNotIn(value []int64) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) OrderById(asc bool) *omsOrderOperateHistoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereOrderId(p mysql.Predicate, value int64) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_id", p),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereOrderIdIn(value []int64) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_id", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereOrderIdNotIn(value []int64) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) OrderByOrderId(asc bool) *omsOrderOperateHistoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "order_id "+order)
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereOperateMan(p mysql.Predicate, value string) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate_man", p),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereOperateManIn(value []string) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate_man", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereOperateManNotIn(value []string) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate_man", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) OrderByOperateMan(asc bool) *omsOrderOperateHistoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "operate_man "+order)
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereCreateTime(p mysql.Predicate, value time.Time) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereCreateTimeIn(value []time.Time) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereCreateTimeNotIn(value []time.Time) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) OrderByCreateTime(asc bool) *omsOrderOperateHistoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereOrderStatus(p mysql.Predicate, value int32) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_status", p),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereOrderStatusIn(value []int32) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_status", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereOrderStatusNotIn(value []int32) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) OrderByOrderStatus(asc bool) *omsOrderOperateHistoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "order_status "+order)
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereNote(p mysql.Predicate, value string) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "note", p),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereNoteIn(value []string) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "note", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) WhereNoteNotIn(value []string) *omsOrderOperateHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "note", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderOperateHistoryQueryBuilder) OrderByNote(asc bool) *omsOrderOperateHistoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "note "+order)
	return qb
}
