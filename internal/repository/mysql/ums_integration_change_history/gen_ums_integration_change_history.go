///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package ums_integration_change_history

import (
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *UmsIntegrationChangeHistory {
	return new(UmsIntegrationChangeHistory)
}

func NewQueryBuilder() *umsIntegrationChangeHistoryQueryBuilder {
	return new(umsIntegrationChangeHistoryQueryBuilder)
}

func (t *UmsIntegrationChangeHistory) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type umsIntegrationChangeHistoryQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *umsIntegrationChangeHistoryQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&UmsIntegrationChangeHistory{})

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

func (qb *umsIntegrationChangeHistoryQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&UmsIntegrationChangeHistory{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&UmsIntegrationChangeHistory{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) First(db *gorm.DB) (*UmsIntegrationChangeHistory, error) {
	ret := &UmsIntegrationChangeHistory{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) QueryOne(db *gorm.DB) (*UmsIntegrationChangeHistory, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) QueryAll(db *gorm.DB) ([]*UmsIntegrationChangeHistory, error) {
	var ret []*UmsIntegrationChangeHistory
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) Limit(limit int) *umsIntegrationChangeHistoryQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) Offset(offset int) *umsIntegrationChangeHistoryQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereId(p mysql.Predicate, value int64) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereIdIn(value []int64) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereIdNotIn(value []int64) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) OrderById(asc bool) *umsIntegrationChangeHistoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereMemberId(p mysql.Predicate, value int64) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_id", p),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereMemberIdIn(value []int64) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_id", "IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereMemberIdNotIn(value []int64) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) OrderByMemberId(asc bool) *umsIntegrationChangeHistoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "member_id "+order)
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereCreateTime(p mysql.Predicate, value time.Time) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereCreateTimeIn(value []time.Time) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereCreateTimeNotIn(value []time.Time) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) OrderByCreateTime(asc bool) *umsIntegrationChangeHistoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereChangeType(p mysql.Predicate, value int32) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "change_type", p),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereChangeTypeIn(value []int32) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "change_type", "IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereChangeTypeNotIn(value []int32) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "change_type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) OrderByChangeType(asc bool) *umsIntegrationChangeHistoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "change_type "+order)
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereChangeCount(p mysql.Predicate, value int32) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "change_count", p),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereChangeCountIn(value []int32) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "change_count", "IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereChangeCountNotIn(value []int32) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "change_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) OrderByChangeCount(asc bool) *umsIntegrationChangeHistoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "change_count "+order)
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereOperateMan(p mysql.Predicate, value string) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate_man", p),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereOperateManIn(value []string) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate_man", "IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereOperateManNotIn(value []string) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate_man", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) OrderByOperateMan(asc bool) *umsIntegrationChangeHistoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "operate_man "+order)
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereOperateNote(p mysql.Predicate, value string) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate_note", p),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereOperateNoteIn(value []string) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate_note", "IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereOperateNoteNotIn(value []string) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate_note", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) OrderByOperateNote(asc bool) *umsIntegrationChangeHistoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "operate_note "+order)
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereSourceType(p mysql.Predicate, value int32) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "source_type", p),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereSourceTypeIn(value []int32) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "source_type", "IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) WhereSourceTypeNotIn(value []int32) *umsIntegrationChangeHistoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "source_type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsIntegrationChangeHistoryQueryBuilder) OrderBySourceType(asc bool) *umsIntegrationChangeHistoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "source_type "+order)
	return qb
}
