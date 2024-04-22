package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"

	"gorm.io/gorm"
)

type UmsAdminDao struct{}

// 分页获取admin列表
func (t *UmsAdminDao) AdminPageList(ctx context.Context, tx *gorm.DB, keyword string, pageSize, pageNum int) (
	[]ums_admin.UmsAdmin, int64, error) {
	res := make([]ums_admin.UmsAdmin, 0)
	var total int64
	// 构建查询条件
	query := tx.WithContext(ctx).Model(&ums_admin.UmsAdmin{})
	if keyword != "" {
		query = query.Where("username LIKE ? OR nick_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	query.Count(&total)
	// 分页查询
	query = query.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	// 执行查询
	err := query.Find(&res).Error
	return res, total, err
}
