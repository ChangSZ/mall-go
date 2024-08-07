///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package ums_member_receive_address

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
)

func NewModel() *UmsMemberReceiveAddress {
	return new(UmsMemberReceiveAddress)
}

func NewQueryBuilder() *umsMemberReceiveAddressQueryBuilder {
	return new(umsMemberReceiveAddressQueryBuilder)
}

func (t *UmsMemberReceiveAddress) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type umsMemberReceiveAddressQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *umsMemberReceiveAddressQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *umsMemberReceiveAddressQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&UmsMemberReceiveAddress{})

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

func (qb *umsMemberReceiveAddressQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&UmsMemberReceiveAddress{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *umsMemberReceiveAddressQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&UmsMemberReceiveAddress{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *umsMemberReceiveAddressQueryBuilder) First(db *gorm.DB) (*UmsMemberReceiveAddress, error) {
	ret := &UmsMemberReceiveAddress{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *umsMemberReceiveAddressQueryBuilder) QueryOne(db *gorm.DB) (*UmsMemberReceiveAddress, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *umsMemberReceiveAddressQueryBuilder) QueryAll(db *gorm.DB) ([]*UmsMemberReceiveAddress, error) {
	var ret []*UmsMemberReceiveAddress
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *umsMemberReceiveAddressQueryBuilder) Limit(limit int) *umsMemberReceiveAddressQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) Offset(offset int) *umsMemberReceiveAddressQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereId(p mysql.Predicate, value int64) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereIdIn(value []int64) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereIdNotIn(value []int64) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) OrderById(asc bool) *umsMemberReceiveAddressQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereMemberId(p mysql.Predicate, value int64) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_id", p),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereMemberIdIn(value []int64) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_id", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereMemberIdNotIn(value []int64) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) OrderByMemberId(asc bool) *umsMemberReceiveAddressQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "member_id "+order)
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereName(p mysql.Predicate, value string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereNameIn(value []string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereNameNotIn(value []string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) OrderByName(asc bool) *umsMemberReceiveAddressQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WherePhoneNumber(p mysql.Predicate, value string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phone_number", p),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WherePhoneNumberIn(value []string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phone_number", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WherePhoneNumberNotIn(value []string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phone_number", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) OrderByPhoneNumber(asc bool) *umsMemberReceiveAddressQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "phone_number "+order)
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereDefaultStatus(p mysql.Predicate, value int32) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "default_status", p),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereDefaultStatusIn(value []int32) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "default_status", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereDefaultStatusNotIn(value []int32) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "default_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) OrderByDefaultStatus(asc bool) *umsMemberReceiveAddressQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "default_status "+order)
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WherePostCode(p mysql.Predicate, value string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "post_code", p),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WherePostCodeIn(value []string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "post_code", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WherePostCodeNotIn(value []string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "post_code", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) OrderByPostCode(asc bool) *umsMemberReceiveAddressQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "post_code "+order)
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereProvince(p mysql.Predicate, value string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "province", p),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereProvinceIn(value []string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "province", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereProvinceNotIn(value []string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "province", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) OrderByProvince(asc bool) *umsMemberReceiveAddressQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "province "+order)
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereCity(p mysql.Predicate, value string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "city", p),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereCityIn(value []string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "city", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereCityNotIn(value []string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "city", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) OrderByCity(asc bool) *umsMemberReceiveAddressQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "city "+order)
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereRegion(p mysql.Predicate, value string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "region", p),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereRegionIn(value []string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "region", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereRegionNotIn(value []string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "region", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) OrderByRegion(asc bool) *umsMemberReceiveAddressQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "region "+order)
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereDetailAddress(p mysql.Predicate, value string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "detail_address", p),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereDetailAddressIn(value []string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "detail_address", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) WhereDetailAddressNotIn(value []string) *umsMemberReceiveAddressQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "detail_address", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberReceiveAddressQueryBuilder) OrderByDetailAddress(asc bool) *umsMemberReceiveAddressQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "detail_address "+order)
	return qb
}
