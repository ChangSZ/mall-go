///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package pms_comment_replay

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
)

func NewModel() *PmsCommentReplay {
	return new(PmsCommentReplay)
}

func NewQueryBuilder() *pmsCommentReplayQueryBuilder {
	return new(pmsCommentReplayQueryBuilder)
}

func (t *PmsCommentReplay) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type pmsCommentReplayQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *pmsCommentReplayQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *pmsCommentReplayQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&PmsCommentReplay{})

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

func (qb *pmsCommentReplayQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&PmsCommentReplay{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *pmsCommentReplayQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&PmsCommentReplay{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *pmsCommentReplayQueryBuilder) First(db *gorm.DB) (*PmsCommentReplay, error) {
	ret := &PmsCommentReplay{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *pmsCommentReplayQueryBuilder) QueryOne(db *gorm.DB) (*PmsCommentReplay, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *pmsCommentReplayQueryBuilder) QueryAll(db *gorm.DB) ([]*PmsCommentReplay, error) {
	var ret []*PmsCommentReplay
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *pmsCommentReplayQueryBuilder) Limit(limit int) *pmsCommentReplayQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) Offset(offset int) *pmsCommentReplayQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereId(p mysql.Predicate, value int64) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereIdIn(value []int64) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereIdNotIn(value []int64) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) OrderById(asc bool) *pmsCommentReplayQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereCommentId(p mysql.Predicate, value int64) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "comment_id", p),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereCommentIdIn(value []int64) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "comment_id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereCommentIdNotIn(value []int64) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "comment_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) OrderByCommentId(asc bool) *pmsCommentReplayQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "comment_id "+order)
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereMemberNickName(p mysql.Predicate, value string) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_nick_name", p),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereMemberNickNameIn(value []string) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_nick_name", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereMemberNickNameNotIn(value []string) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_nick_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) OrderByMemberNickName(asc bool) *pmsCommentReplayQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "member_nick_name "+order)
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereMemberIcon(p mysql.Predicate, value string) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_icon", p),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereMemberIconIn(value []string) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_icon", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereMemberIconNotIn(value []string) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_icon", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) OrderByMemberIcon(asc bool) *pmsCommentReplayQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "member_icon "+order)
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereContent(p mysql.Predicate, value string) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", p),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereContentIn(value []string) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereContentNotIn(value []string) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) OrderByContent(asc bool) *pmsCommentReplayQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "content "+order)
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereCreateTime(p mysql.Predicate, value time.Time) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereCreateTimeIn(value []time.Time) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereCreateTimeNotIn(value []time.Time) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) OrderByCreateTime(asc bool) *pmsCommentReplayQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereType(p mysql.Predicate, value int32) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", p),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereTypeIn(value []int32) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) WhereTypeNotIn(value []int32) *pmsCommentReplayQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentReplayQueryBuilder) OrderByType(asc bool) *pmsCommentReplayQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "type "+order)
	return qb
}
