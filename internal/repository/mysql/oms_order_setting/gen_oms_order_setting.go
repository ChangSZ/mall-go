///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package oms_order_setting

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *OmsOrderSetting {
	return new(OmsOrderSetting)
}

func NewQueryBuilder() *omsOrderSettingQueryBuilder {
	return new(omsOrderSettingQueryBuilder)
}

func (t *OmsOrderSetting) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type omsOrderSettingQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *omsOrderSettingQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *omsOrderSettingQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&OmsOrderSetting{})

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

func (qb *omsOrderSettingQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&OmsOrderSetting{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *omsOrderSettingQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&OmsOrderSetting{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *omsOrderSettingQueryBuilder) First(db *gorm.DB) (*OmsOrderSetting, error) {
	ret := &OmsOrderSetting{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *omsOrderSettingQueryBuilder) QueryOne(db *gorm.DB) (*OmsOrderSetting, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *omsOrderSettingQueryBuilder) QueryAll(db *gorm.DB) ([]*OmsOrderSetting, error) {
	var ret []*OmsOrderSetting
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *omsOrderSettingQueryBuilder) Limit(limit int) *omsOrderSettingQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *omsOrderSettingQueryBuilder) Offset(offset int) *omsOrderSettingQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereId(p mysql.Predicate, value int64) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereIdIn(value []int64) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereIdNotIn(value []int64) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) OrderById(asc bool) *omsOrderSettingQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereFlashOrderOvertime(p mysql.Predicate, value int32) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_order_overtime", p),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereFlashOrderOvertimeIn(value []int32) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_order_overtime", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereFlashOrderOvertimeNotIn(value []int32) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "flash_order_overtime", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) OrderByFlashOrderOvertime(asc bool) *omsOrderSettingQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "flash_order_overtime "+order)
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereNormalOrderOvertime(p mysql.Predicate, value int32) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "normal_order_overtime", p),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereNormalOrderOvertimeIn(value []int32) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "normal_order_overtime", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereNormalOrderOvertimeNotIn(value []int32) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "normal_order_overtime", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) OrderByNormalOrderOvertime(asc bool) *omsOrderSettingQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "normal_order_overtime "+order)
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereConfirmOvertime(p mysql.Predicate, value int32) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "confirm_overtime", p),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereConfirmOvertimeIn(value []int32) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "confirm_overtime", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereConfirmOvertimeNotIn(value []int32) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "confirm_overtime", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) OrderByConfirmOvertime(asc bool) *omsOrderSettingQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "confirm_overtime "+order)
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereFinishOvertime(p mysql.Predicate, value int32) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "finish_overtime", p),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereFinishOvertimeIn(value []int32) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "finish_overtime", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereFinishOvertimeNotIn(value []int32) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "finish_overtime", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) OrderByFinishOvertime(asc bool) *omsOrderSettingQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "finish_overtime "+order)
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereCommentOvertime(p mysql.Predicate, value int32) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "comment_overtime", p),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereCommentOvertimeIn(value []int32) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "comment_overtime", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) WhereCommentOvertimeNotIn(value []int32) *omsOrderSettingQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "comment_overtime", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderSettingQueryBuilder) OrderByCommentOvertime(asc bool) *omsOrderSettingQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "comment_overtime "+order)
	return qb
}
