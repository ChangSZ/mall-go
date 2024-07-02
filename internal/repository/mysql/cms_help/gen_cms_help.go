///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package cms_help

import (
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *CmsHelp {
	return new(CmsHelp)
}

func NewQueryBuilder() *cmsHelpQueryBuilder {
	return new(cmsHelpQueryBuilder)
}

func (t *CmsHelp) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type cmsHelpQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *cmsHelpQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *cmsHelpQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&CmsHelp{})

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

func (qb *cmsHelpQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&CmsHelp{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *cmsHelpQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&CmsHelp{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *cmsHelpQueryBuilder) First(db *gorm.DB) (*CmsHelp, error) {
	ret := &CmsHelp{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *cmsHelpQueryBuilder) QueryOne(db *gorm.DB) (*CmsHelp, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *cmsHelpQueryBuilder) QueryAll(db *gorm.DB) ([]*CmsHelp, error) {
	var ret []*CmsHelp
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *cmsHelpQueryBuilder) Limit(limit int) *cmsHelpQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *cmsHelpQueryBuilder) Offset(offset int) *cmsHelpQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereId(p mysql.Predicate, value int64) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereIdIn(value []int64) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereIdNotIn(value []int64) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) OrderById(asc bool) *cmsHelpQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereCategoryId(p mysql.Predicate, value int64) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", p),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereCategoryIdIn(value []int64) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", "IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereCategoryIdNotIn(value []int64) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) OrderByCategoryId(asc bool) *cmsHelpQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "category_id "+order)
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereIcon(p mysql.Predicate, value string) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "icon", p),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereIconIn(value []string) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "icon", "IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereIconNotIn(value []string) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "icon", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) OrderByIcon(asc bool) *cmsHelpQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "icon "+order)
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereTitle(p mysql.Predicate, value string) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", p),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereTitleIn(value []string) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", "IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereTitleNotIn(value []string) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) OrderByTitle(asc bool) *cmsHelpQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "title "+order)
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereShowStatus(p mysql.Predicate, value int32) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "show_status", p),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereShowStatusIn(value []int32) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "show_status", "IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereShowStatusNotIn(value []int32) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "show_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) OrderByShowStatus(asc bool) *cmsHelpQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "show_status "+order)
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereCreateTime(p mysql.Predicate, value time.Time) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereCreateTimeIn(value []time.Time) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereCreateTimeNotIn(value []time.Time) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) OrderByCreateTime(asc bool) *cmsHelpQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereReadCount(p mysql.Predicate, value int32) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "read_count", p),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereReadCountIn(value []int32) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "read_count", "IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereReadCountNotIn(value []int32) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "read_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) OrderByReadCount(asc bool) *cmsHelpQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "read_count "+order)
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereContent(p mysql.Predicate, value string) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", p),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereContentIn(value []string) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) WhereContentNotIn(value []string) *cmsHelpQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsHelpQueryBuilder) OrderByContent(asc bool) *cmsHelpQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "content "+order)
	return qb
}
