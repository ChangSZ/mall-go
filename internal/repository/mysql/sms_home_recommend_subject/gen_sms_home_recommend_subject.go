///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package sms_home_recommend_subject

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
)

func NewModel() *SmsHomeRecommendSubject {
	return new(SmsHomeRecommendSubject)
}

func NewQueryBuilder() *smsHomeRecommendSubjectQueryBuilder {
	return new(smsHomeRecommendSubjectQueryBuilder)
}

func (t *SmsHomeRecommendSubject) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type smsHomeRecommendSubjectQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *smsHomeRecommendSubjectQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *smsHomeRecommendSubjectQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&SmsHomeRecommendSubject{})

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

func (qb *smsHomeRecommendSubjectQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&SmsHomeRecommendSubject{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *smsHomeRecommendSubjectQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&SmsHomeRecommendSubject{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *smsHomeRecommendSubjectQueryBuilder) First(db *gorm.DB) (*SmsHomeRecommendSubject, error) {
	ret := &SmsHomeRecommendSubject{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *smsHomeRecommendSubjectQueryBuilder) QueryOne(db *gorm.DB) (*SmsHomeRecommendSubject, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *smsHomeRecommendSubjectQueryBuilder) QueryAll(db *gorm.DB) ([]*SmsHomeRecommendSubject, error) {
	var ret []*SmsHomeRecommendSubject
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *smsHomeRecommendSubjectQueryBuilder) Limit(limit int) *smsHomeRecommendSubjectQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) Offset(offset int) *smsHomeRecommendSubjectQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) WhereId(p mysql.Predicate, value int64) *smsHomeRecommendSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) WhereIdIn(value []int64) *smsHomeRecommendSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) WhereIdNotIn(value []int64) *smsHomeRecommendSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) OrderById(asc bool) *smsHomeRecommendSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) WhereSubjectId(p mysql.Predicate, value int64) *smsHomeRecommendSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "subject_id", p),
		value,
	})
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) WhereSubjectIdIn(value []int64) *smsHomeRecommendSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "subject_id", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) WhereSubjectIdNotIn(value []int64) *smsHomeRecommendSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "subject_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) OrderBySubjectId(asc bool) *smsHomeRecommendSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "subject_id "+order)
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) WhereSubjectName(p mysql.Predicate, value string) *smsHomeRecommendSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "subject_name", p),
		value,
	})
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) WhereSubjectNameIn(value []string) *smsHomeRecommendSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "subject_name", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) WhereSubjectNameNotIn(value []string) *smsHomeRecommendSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "subject_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) OrderBySubjectName(asc bool) *smsHomeRecommendSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "subject_name "+order)
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) WhereRecommendStatus(p mysql.Predicate, value int32) *smsHomeRecommendSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "recommend_status", p),
		value,
	})
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) WhereRecommendStatusIn(value []int32) *smsHomeRecommendSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "recommend_status", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) WhereRecommendStatusNotIn(value []int32) *smsHomeRecommendSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "recommend_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) OrderByRecommendStatus(asc bool) *smsHomeRecommendSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "recommend_status "+order)
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) WhereSort(p mysql.Predicate, value int32) *smsHomeRecommendSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", p),
		value,
	})
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) WhereSortIn(value []int32) *smsHomeRecommendSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "IN"),
		value,
	})
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) WhereSortNotIn(value []int32) *smsHomeRecommendSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsHomeRecommendSubjectQueryBuilder) OrderBySort(asc bool) *smsHomeRecommendSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sort "+order)
	return qb
}
