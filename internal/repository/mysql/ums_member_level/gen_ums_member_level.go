///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package ums_member_level

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *UmsMemberLevel {
	return new(UmsMemberLevel)
}

func NewQueryBuilder() *umsMemberLevelQueryBuilder {
	return new(umsMemberLevelQueryBuilder)
}

func (t *UmsMemberLevel) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type umsMemberLevelQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *umsMemberLevelQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *umsMemberLevelQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&UmsMemberLevel{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *umsMemberLevelQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&UmsMemberLevel{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *umsMemberLevelQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&UmsMemberLevel{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *umsMemberLevelQueryBuilder) First(db *gorm.DB) (*UmsMemberLevel, error) {
	ret := &UmsMemberLevel{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *umsMemberLevelQueryBuilder) QueryOne(db *gorm.DB) (*UmsMemberLevel, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *umsMemberLevelQueryBuilder) QueryAll(db *gorm.DB) ([]*UmsMemberLevel, error) {
	var ret []*UmsMemberLevel
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *umsMemberLevelQueryBuilder) Limit(limit int) *umsMemberLevelQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *umsMemberLevelQueryBuilder) Offset(offset int) *umsMemberLevelQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereId(p mysql.Predicate, value int64) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereIdIn(value []int64) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereIdNotIn(value []int64) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) OrderById(asc bool) *umsMemberLevelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereName(p mysql.Predicate, value string) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereNameIn(value []string) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereNameNotIn(value []string) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) OrderByName(asc bool) *umsMemberLevelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereGrowthPoint(p mysql.Predicate, value int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "growth_point", p),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereGrowthPointIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "growth_point", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereGrowthPointNotIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "growth_point", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) OrderByGrowthPoint(asc bool) *umsMemberLevelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "growth_point "+order)
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereDefaultStatus(p mysql.Predicate, value int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "default_status", p),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereDefaultStatusIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "default_status", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereDefaultStatusNotIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "default_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) OrderByDefaultStatus(asc bool) *umsMemberLevelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "default_status "+order)
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereFreeFreightPoint(p mysql.Predicate, value float64) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "free_freight_point", p),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereFreeFreightPointIn(value []float64) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "free_freight_point", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereFreeFreightPointNotIn(value []float64) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "free_freight_point", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) OrderByFreeFreightPoint(asc bool) *umsMemberLevelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "free_freight_point "+order)
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereCommentGrowthPoint(p mysql.Predicate, value int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "comment_growth_point", p),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereCommentGrowthPointIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "comment_growth_point", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereCommentGrowthPointNotIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "comment_growth_point", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) OrderByCommentGrowthPoint(asc bool) *umsMemberLevelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "comment_growth_point "+order)
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgeFreeFreight(p mysql.Predicate, value int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_free_freight", p),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgeFreeFreightIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_free_freight", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgeFreeFreightNotIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_free_freight", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) OrderByPriviledgeFreeFreight(asc bool) *umsMemberLevelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "priviledge_free_freight "+order)
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgeSignIn(p mysql.Predicate, value int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_sign_in", p),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgeSignInIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_sign_in", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgeSignInNotIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_sign_in", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) OrderByPriviledgeSignIn(asc bool) *umsMemberLevelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "priviledge_sign_in "+order)
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgeComment(p mysql.Predicate, value int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_comment", p),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgeCommentIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_comment", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgeCommentNotIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_comment", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) OrderByPriviledgeComment(asc bool) *umsMemberLevelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "priviledge_comment "+order)
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgePromotion(p mysql.Predicate, value int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_promotion", p),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgePromotionIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_promotion", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgePromotionNotIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_promotion", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) OrderByPriviledgePromotion(asc bool) *umsMemberLevelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "priviledge_promotion "+order)
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgeMemberPrice(p mysql.Predicate, value int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_member_price", p),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgeMemberPriceIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_member_price", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgeMemberPriceNotIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_member_price", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) OrderByPriviledgeMemberPrice(asc bool) *umsMemberLevelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "priviledge_member_price "+order)
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgeBirthday(p mysql.Predicate, value int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_birthday", p),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgeBirthdayIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_birthday", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WherePriviledgeBirthdayNotIn(value []int32) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "priviledge_birthday", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) OrderByPriviledgeBirthday(asc bool) *umsMemberLevelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "priviledge_birthday "+order)
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereNote(p mysql.Predicate, value string) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "note", p),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereNoteIn(value []string) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "note", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) WhereNoteNotIn(value []string) *umsMemberLevelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "note", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLevelQueryBuilder) OrderByNote(asc bool) *umsMemberLevelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "note "+order)
	return qb
}
