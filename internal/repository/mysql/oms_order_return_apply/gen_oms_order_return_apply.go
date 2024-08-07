///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package oms_order_return_apply

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
)

func NewModel() *OmsOrderReturnApply {
	return new(OmsOrderReturnApply)
}

func NewQueryBuilder() *omsOrderReturnApplyQueryBuilder {
	return new(omsOrderReturnApplyQueryBuilder)
}

func (t *OmsOrderReturnApply) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type omsOrderReturnApplyQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *omsOrderReturnApplyQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *omsOrderReturnApplyQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&OmsOrderReturnApply{})

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

func (qb *omsOrderReturnApplyQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&OmsOrderReturnApply{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *omsOrderReturnApplyQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&OmsOrderReturnApply{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *omsOrderReturnApplyQueryBuilder) First(db *gorm.DB) (*OmsOrderReturnApply, error) {
	ret := &OmsOrderReturnApply{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *omsOrderReturnApplyQueryBuilder) QueryOne(db *gorm.DB) (*OmsOrderReturnApply, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *omsOrderReturnApplyQueryBuilder) QueryAll(db *gorm.DB) ([]*OmsOrderReturnApply, error) {
	var ret []*OmsOrderReturnApply
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *omsOrderReturnApplyQueryBuilder) Limit(limit int) *omsOrderReturnApplyQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) Offset(offset int) *omsOrderReturnApplyQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereId(p mysql.Predicate, value int64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereIdIn(value []int64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereIdNotIn(value []int64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderById(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereOrderId(p mysql.Predicate, value int64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_id", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereOrderIdIn(value []int64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_id", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereOrderIdNotIn(value []int64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByOrderId(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "order_id "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereCompanyAddressId(p mysql.Predicate, value int64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company_address_id", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereCompanyAddressIdIn(value []int64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company_address_id", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereCompanyAddressIdNotIn(value []int64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company_address_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByCompanyAddressId(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "company_address_id "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductId(p mysql.Predicate, value int64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductIdIn(value []int64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductIdNotIn(value []int64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByProductId(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_id "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereOrderSn(p mysql.Predicate, value string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_sn", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereOrderSnIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_sn", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereOrderSnNotIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "order_sn", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByOrderSn(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "order_sn "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereCreateTime(p mysql.Predicate, value time.Time) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereCreateTimeIn(value []time.Time) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereCreateTimeNotIn(value []time.Time) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByCreateTime(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereMemberUsername(p mysql.Predicate, value string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_username", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereMemberUsernameIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_username", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereMemberUsernameNotIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_username", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByMemberUsername(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "member_username "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReturnAmount(p mysql.Predicate, value float64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "return_amount", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReturnAmountIn(value []float64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "return_amount", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReturnAmountNotIn(value []float64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "return_amount", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByReturnAmount(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "return_amount "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReturnName(p mysql.Predicate, value string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "return_name", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReturnNameIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "return_name", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReturnNameNotIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "return_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByReturnName(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "return_name "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReturnPhone(p mysql.Predicate, value string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "return_phone", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReturnPhoneIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "return_phone", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReturnPhoneNotIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "return_phone", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByReturnPhone(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "return_phone "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereStatus(p mysql.Predicate, value int32) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereStatusIn(value []int32) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereStatusNotIn(value []int32) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByStatus(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "status "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereHandleTime(p mysql.Predicate, value time.Time) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "handle_time", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereHandleTimeIn(value []time.Time) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "handle_time", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereHandleTimeNotIn(value []time.Time) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "handle_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByHandleTime(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "handle_time "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductPic(p mysql.Predicate, value string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_pic", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductPicIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_pic", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductPicNotIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_pic", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByProductPic(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_pic "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductName(p mysql.Predicate, value string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_name", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductNameIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_name", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductNameNotIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByProductName(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_name "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductBrand(p mysql.Predicate, value string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_brand", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductBrandIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_brand", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductBrandNotIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_brand", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByProductBrand(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_brand "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductAttr(p mysql.Predicate, value string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attr", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductAttrIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attr", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductAttrNotIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attr", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByProductAttr(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_attr "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductCount(p mysql.Predicate, value int32) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_count", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductCountIn(value []int32) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_count", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductCountNotIn(value []int32) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByProductCount(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_count "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductPrice(p mysql.Predicate, value float64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_price", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductPriceIn(value []float64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_price", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductPriceNotIn(value []float64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_price", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByProductPrice(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_price "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductRealPrice(p mysql.Predicate, value float64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_real_price", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductRealPriceIn(value []float64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_real_price", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProductRealPriceNotIn(value []float64) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_real_price", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByProductRealPrice(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_real_price "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReason(p mysql.Predicate, value string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "reason", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReasonIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "reason", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReasonNotIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "reason", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByReason(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "reason "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereDescription(p mysql.Predicate, value string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereDescriptionIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereDescriptionNotIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByDescription(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "description "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProofPics(p mysql.Predicate, value string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "proof_pics", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProofPicsIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "proof_pics", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereProofPicsNotIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "proof_pics", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByProofPics(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "proof_pics "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereHandleNote(p mysql.Predicate, value string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "handle_note", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereHandleNoteIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "handle_note", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereHandleNoteNotIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "handle_note", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByHandleNote(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "handle_note "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereHandleMan(p mysql.Predicate, value string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "handle_man", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereHandleManIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "handle_man", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereHandleManNotIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "handle_man", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByHandleMan(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "handle_man "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReceiveMan(p mysql.Predicate, value string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "receive_man", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReceiveManIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "receive_man", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReceiveManNotIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "receive_man", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByReceiveMan(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "receive_man "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReceiveTime(p mysql.Predicate, value time.Time) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "receive_time", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReceiveTimeIn(value []time.Time) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "receive_time", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReceiveTimeNotIn(value []time.Time) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "receive_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByReceiveTime(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "receive_time "+order)
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReceiveNote(p mysql.Predicate, value string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "receive_note", p),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReceiveNoteIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "receive_note", "IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) WhereReceiveNoteNotIn(value []string) *omsOrderReturnApplyQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "receive_note", "NOT IN"),
		value,
	})
	return qb
}

func (qb *omsOrderReturnApplyQueryBuilder) OrderByReceiveNote(asc bool) *omsOrderReturnApplyQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "receive_note "+order)
	return qb
}
