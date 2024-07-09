///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package sms_home_advertise

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
)

func NewModel() *SmsHomeAdvertise {
	return new(SmsHomeAdvertise)
}

func NewQueryBuilder() *smsHomeAdvertiseQueryBuilder {
	return new(smsHomeAdvertiseQueryBuilder)
}

func (t *SmsHomeAdvertise) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type smsHomeAdvertiseQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *smsHomeAdvertiseQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *smsHomeAdvertiseQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&SmsHomeAdvertise{})

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

func (qb *smsHomeAdvertiseQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&SmsHomeAdvertise{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *smsHomeAdvertiseQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&SmsHomeAdvertise{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *smsHomeAdvertiseQueryBuilder) First(db *gorm.DB) (*SmsHomeAdvertise, error) {
	ret := &SmsHomeAdvertise{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *smsHomeAdvertiseQueryBuilder) QueryOne(db *gorm.DB) (*SmsHomeAdvertise, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *smsHomeAdvertiseQueryBuilder) QueryAll(db *gorm.DB) ([]*SmsHomeAdvertise, error) {
	var ret []*SmsHomeAdvertise
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *smsHomeAdvertiseQueryBuilder) Limit(limit int) *smsHomeAdvertiseQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) Offset(offset int) *smsHomeAdvertiseQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereId(p mysql.Predicate, value int64) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereIdIn(value []int64) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereIdNotIn(value []int64) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) OrderById(asc bool) *smsHomeAdvertiseQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereName(p mysql.Predicate, value string) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereNameIn(value []string) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereNameNotIn(value []string) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) OrderByName(asc bool) *smsHomeAdvertiseQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereType(p mysql.Predicate, value int32) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", p),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereTypeIn(value []int32) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereTypeNotIn(value []int32) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) OrderByType(asc bool) *smsHomeAdvertiseQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "type "+order)
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WherePic(p mysql.Predicate, value string) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pic", p),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WherePicIn(value []string) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pic", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WherePicNotIn(value []string) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pic", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) OrderByPic(asc bool) *smsHomeAdvertiseQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "pic "+order)
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereStartTime(p mysql.Predicate, value time.Time) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "start_time", p),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereStartTimeIn(value []time.Time) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "start_time", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereStartTimeNotIn(value []time.Time) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "start_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) OrderByStartTime(asc bool) *smsHomeAdvertiseQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "start_time "+order)
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereEndTime(p mysql.Predicate, value time.Time) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "end_time", p),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereEndTimeIn(value []time.Time) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "end_time", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereEndTimeNotIn(value []time.Time) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "end_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) OrderByEndTime(asc bool) *smsHomeAdvertiseQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "end_time "+order)
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereStatus(p mysql.Predicate, value int32) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", p),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereStatusIn(value []int32) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereStatusNotIn(value []int32) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) OrderByStatus(asc bool) *smsHomeAdvertiseQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "status "+order)
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereClickCount(p mysql.Predicate, value int32) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "click_count", p),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereClickCountIn(value []int32) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "click_count", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereClickCountNotIn(value []int32) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "click_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) OrderByClickCount(asc bool) *smsHomeAdvertiseQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "click_count "+order)
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereOrderCount(p mysql.Predicate, value int32) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_count", p),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereOrderCountIn(value []int32) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_count", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereOrderCountNotIn(value []int32) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) OrderByOrderCount(asc bool) *smsHomeAdvertiseQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "order_count "+order)
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereUrl(p mysql.Predicate, value string) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "url", p),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereUrlIn(value []string) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "url", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereUrlNotIn(value []string) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "url", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) OrderByUrl(asc bool) *smsHomeAdvertiseQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "url "+order)
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereNote(p mysql.Predicate, value string) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "note", p),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereNoteIn(value []string) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "note", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereNoteNotIn(value []string) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "note", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) OrderByNote(asc bool) *smsHomeAdvertiseQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "note "+order)
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereSort(p mysql.Predicate, value int32) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", p),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereSortIn(value []int32) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) WhereSortNotIn(value []int32) *smsHomeAdvertiseQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeAdvertiseQueryBuilder) OrderBySort(asc bool) *smsHomeAdvertiseQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sort "+order)
	return qb
}
