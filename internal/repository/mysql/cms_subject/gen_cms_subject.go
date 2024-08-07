///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package cms_subject

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
)

func NewModel() *CmsSubject {
	return new(CmsSubject)
}

func NewQueryBuilder() *cmsSubjectQueryBuilder {
	return new(cmsSubjectQueryBuilder)
}

func (t *CmsSubject) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type cmsSubjectQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *cmsSubjectQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *cmsSubjectQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&CmsSubject{})

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

func (qb *cmsSubjectQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&CmsSubject{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *cmsSubjectQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&CmsSubject{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *cmsSubjectQueryBuilder) First(db *gorm.DB) (*CmsSubject, error) {
	ret := &CmsSubject{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ret, res.Error
}

func (qb *cmsSubjectQueryBuilder) QueryOne(db *gorm.DB) (*CmsSubject, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *cmsSubjectQueryBuilder) QueryAll(db *gorm.DB) ([]*CmsSubject, error) {
	var ret []*CmsSubject
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *cmsSubjectQueryBuilder) Limit(limit int) *cmsSubjectQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *cmsSubjectQueryBuilder) Offset(offset int) *cmsSubjectQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereId(p mysql.Predicate, value int64) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereIdIn(value []int64) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereIdNotIn(value []int64) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderById(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereCategoryId(p mysql.Predicate, value int64) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereCategoryIdIn(value []int64) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereCategoryIdNotIn(value []int64) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderByCategoryId(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "category_id "+order)
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereTitle(p mysql.Predicate, value string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereTitleIn(value []string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereTitleNotIn(value []string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderByTitle(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "title "+order)
	return qb
}

func (qb *cmsSubjectQueryBuilder) WherePic(p mysql.Predicate, value string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pic", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WherePicIn(value []string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pic", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WherePicNotIn(value []string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pic", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderByPic(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "pic "+order)
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereProductCount(p mysql.Predicate, value int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_count", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereProductCountIn(value []int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_count", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereProductCountNotIn(value []int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderByProductCount(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_count "+order)
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereRecommendStatus(p mysql.Predicate, value int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "recommend_status", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereRecommendStatusIn(value []int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "recommend_status", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereRecommendStatusNotIn(value []int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "recommend_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderByRecommendStatus(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "recommend_status "+order)
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereCreateTime(p mysql.Predicate, value time.Time) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereCreateTimeIn(value []time.Time) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereCreateTimeNotIn(value []time.Time) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderByCreateTime(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereCollectCount(p mysql.Predicate, value int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "collect_count", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereCollectCountIn(value []int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "collect_count", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereCollectCountNotIn(value []int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "collect_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderByCollectCount(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "collect_count "+order)
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereReadCount(p mysql.Predicate, value int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "read_count", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereReadCountIn(value []int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "read_count", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereReadCountNotIn(value []int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "read_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderByReadCount(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "read_count "+order)
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereCommentCount(p mysql.Predicate, value int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "comment_count", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereCommentCountIn(value []int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "comment_count", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereCommentCountNotIn(value []int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "comment_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderByCommentCount(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "comment_count "+order)
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereAlbumPics(p mysql.Predicate, value string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "album_pics", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereAlbumPicsIn(value []string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "album_pics", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereAlbumPicsNotIn(value []string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "album_pics", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderByAlbumPics(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "album_pics "+order)
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereDescription(p mysql.Predicate, value string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereDescriptionIn(value []string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereDescriptionNotIn(value []string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderByDescription(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "description "+order)
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereShowStatus(p mysql.Predicate, value int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "show_status", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereShowStatusIn(value []int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "show_status", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereShowStatusNotIn(value []int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "show_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderByShowStatus(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "show_status "+order)
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereContent(p mysql.Predicate, value string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereContentIn(value []string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereContentNotIn(value []string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderByContent(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "content "+order)
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereForwardCount(p mysql.Predicate, value int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "forward_count", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereForwardCountIn(value []int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "forward_count", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereForwardCountNotIn(value []int32) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "forward_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderByForwardCount(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "forward_count "+order)
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereCategoryName(p mysql.Predicate, value string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_name", p),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereCategoryNameIn(value []string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_name", "IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) WhereCategoryNameNotIn(value []string) *cmsSubjectQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cmsSubjectQueryBuilder) OrderByCategoryName(asc bool) *cmsSubjectQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "category_name "+order)
	return qb
}
