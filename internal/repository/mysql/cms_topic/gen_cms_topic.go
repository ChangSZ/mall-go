///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package cms_topic

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
)

func NewModel() *CmsTopic {
	return new(CmsTopic)
}

func NewQueryBuilder() *cmsTopicQueryBuilder {
	return new(cmsTopicQueryBuilder)
}

func (t *CmsTopic) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type cmsTopicQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *cmsTopicQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *cmsTopicQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&CmsTopic{})

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

func (qb *cmsTopicQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&CmsTopic{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *cmsTopicQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&CmsTopic{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *cmsTopicQueryBuilder) First(db *gorm.DB) (*CmsTopic, error) {
	ret := &CmsTopic{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *cmsTopicQueryBuilder) QueryOne(db *gorm.DB) (*CmsTopic, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *cmsTopicQueryBuilder) QueryAll(db *gorm.DB) ([]*CmsTopic, error) {
	var ret []*CmsTopic
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *cmsTopicQueryBuilder) Limit(limit int) *cmsTopicQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *cmsTopicQueryBuilder) Offset(offset int) *cmsTopicQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereId(p mysql.Predicate, value int64) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereIdIn(value []int64) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereIdNotIn(value []int64) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) OrderById(asc bool) *cmsTopicQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereCategoryId(p mysql.Predicate, value int64) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", p),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereCategoryIdIn(value []int64) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", "IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereCategoryIdNotIn(value []int64) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) OrderByCategoryId(asc bool) *cmsTopicQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "category_id "+order)
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereName(p mysql.Predicate, value string) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereNameIn(value []string) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereNameNotIn(value []string) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) OrderByName(asc bool) *cmsTopicQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereCreateTime(p mysql.Predicate, value time.Time) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereCreateTimeIn(value []time.Time) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereCreateTimeNotIn(value []time.Time) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) OrderByCreateTime(asc bool) *cmsTopicQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereStartTime(p mysql.Predicate, value time.Time) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "start_time", p),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereStartTimeIn(value []time.Time) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "start_time", "IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereStartTimeNotIn(value []time.Time) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "start_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) OrderByStartTime(asc bool) *cmsTopicQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "start_time "+order)
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereEndTime(p mysql.Predicate, value time.Time) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "end_time", p),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereEndTimeIn(value []time.Time) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "end_time", "IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereEndTimeNotIn(value []time.Time) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "end_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) OrderByEndTime(asc bool) *cmsTopicQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "end_time "+order)
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereAttendCount(p mysql.Predicate, value int32) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "attend_count", p),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereAttendCountIn(value []int32) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "attend_count", "IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereAttendCountNotIn(value []int32) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "attend_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) OrderByAttendCount(asc bool) *cmsTopicQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "attend_count "+order)
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereAttentionCount(p mysql.Predicate, value int32) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "attention_count", p),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereAttentionCountIn(value []int32) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "attention_count", "IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereAttentionCountNotIn(value []int32) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "attention_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) OrderByAttentionCount(asc bool) *cmsTopicQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "attention_count "+order)
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereReadCount(p mysql.Predicate, value int32) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "read_count", p),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereReadCountIn(value []int32) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "read_count", "IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereReadCountNotIn(value []int32) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "read_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) OrderByReadCount(asc bool) *cmsTopicQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "read_count "+order)
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereAwardName(p mysql.Predicate, value string) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "award_name", p),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereAwardNameIn(value []string) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "award_name", "IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereAwardNameNotIn(value []string) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "award_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) OrderByAwardName(asc bool) *cmsTopicQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "award_name "+order)
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereAttendType(p mysql.Predicate, value string) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "attend_type", p),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereAttendTypeIn(value []string) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "attend_type", "IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereAttendTypeNotIn(value []string) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "attend_type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) OrderByAttendType(asc bool) *cmsTopicQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "attend_type "+order)
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereContent(p mysql.Predicate, value string) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", p),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereContentIn(value []string) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) WhereContentNotIn(value []string) *cmsTopicQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsTopicQueryBuilder) OrderByContent(asc bool) *cmsTopicQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "content "+order)
	return qb
}
