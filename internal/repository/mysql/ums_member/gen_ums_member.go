///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package ums_member

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
)

func NewModel() *UmsMember {
	return new(UmsMember)
}

func NewQueryBuilder() *umsMemberQueryBuilder {
	return new(umsMemberQueryBuilder)
}

func (t *UmsMember) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type umsMemberQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *umsMemberQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *umsMemberQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&UmsMember{})

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

func (qb *umsMemberQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&UmsMember{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *umsMemberQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&UmsMember{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *umsMemberQueryBuilder) First(db *gorm.DB) (*UmsMember, error) {
	ret := &UmsMember{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *umsMemberQueryBuilder) QueryOne(db *gorm.DB) (*UmsMember, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *umsMemberQueryBuilder) QueryAll(db *gorm.DB) ([]*UmsMember, error) {
	var ret []*UmsMember
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *umsMemberQueryBuilder) Limit(limit int) *umsMemberQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *umsMemberQueryBuilder) Offset(offset int) *umsMemberQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *umsMemberQueryBuilder) WhereId(p mysql.Predicate, value int64) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereIdIn(value []int64) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereIdNotIn(value []int64) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderById(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WhereMemberLevelId(p mysql.Predicate, value int64) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_level_id", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereMemberLevelIdIn(value []int64) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_level_id", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereMemberLevelIdNotIn(value []int64) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_level_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByMemberLevelId(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "member_level_id "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WhereUsername(p mysql.Predicate, value string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereUsernameIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereUsernameNotIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByUsername(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "username "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WherePassword(p mysql.Predicate, value string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "password", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WherePasswordIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "password", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WherePasswordNotIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "password", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByPassword(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "password "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WhereNickname(p mysql.Predicate, value string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereNicknameIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereNicknameNotIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByNickname(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "nickname "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WherePhone(p mysql.Predicate, value string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phone", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WherePhoneIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phone", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WherePhoneNotIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phone", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByPhone(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "phone "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WhereStatus(p mysql.Predicate, value int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereStatusIn(value []int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereStatusNotIn(value []int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByStatus(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "status "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WhereCreateTime(p mysql.Predicate, value time.Time) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereCreateTimeIn(value []time.Time) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereCreateTimeNotIn(value []time.Time) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByCreateTime(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WhereIcon(p mysql.Predicate, value string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "icon", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereIconIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "icon", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereIconNotIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "icon", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByIcon(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "icon "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WhereGender(p mysql.Predicate, value int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "gender", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereGenderIn(value []int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "gender", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereGenderNotIn(value []int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "gender", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByGender(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "gender "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WhereBirthday(p mysql.Predicate, value time.Time) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "birthday", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereBirthdayIn(value []time.Time) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "birthday", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereBirthdayNotIn(value []time.Time) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "birthday", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByBirthday(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "birthday "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WhereCity(p mysql.Predicate, value string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "city", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereCityIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "city", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereCityNotIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "city", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByCity(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "city "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WhereJob(p mysql.Predicate, value string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "job", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereJobIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "job", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereJobNotIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "job", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByJob(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "job "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WherePersonalizedSignature(p mysql.Predicate, value string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "personalized_signature", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WherePersonalizedSignatureIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "personalized_signature", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WherePersonalizedSignatureNotIn(value []string) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "personalized_signature", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByPersonalizedSignature(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "personalized_signature "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WhereSourceType(p mysql.Predicate, value int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "source_type", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereSourceTypeIn(value []int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "source_type", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereSourceTypeNotIn(value []int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "source_type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderBySourceType(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "source_type "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WhereIntegration(p mysql.Predicate, value int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "integration", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereIntegrationIn(value []int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "integration", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereIntegrationNotIn(value []int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "integration", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByIntegration(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "integration "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WhereGrowth(p mysql.Predicate, value int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "growth", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereGrowthIn(value []int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "growth", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereGrowthNotIn(value []int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "growth", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByGrowth(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "growth "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WhereLuckeyCount(p mysql.Predicate, value int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "luckey_count", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereLuckeyCountIn(value []int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "luckey_count", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereLuckeyCountNotIn(value []int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "luckey_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByLuckeyCount(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "luckey_count "+order)
	return qb
}

func (qb *umsMemberQueryBuilder) WhereHistoryIntegration(p mysql.Predicate, value int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "history_integration", p),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereHistoryIntegrationIn(value []int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "history_integration", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) WhereHistoryIntegrationNotIn(value []int32) *umsMemberQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "history_integration", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberQueryBuilder) OrderByHistoryIntegration(asc bool) *umsMemberQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "history_integration "+order)
	return qb
}
