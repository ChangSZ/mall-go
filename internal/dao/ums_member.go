package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_member"

	"gorm.io/gorm"
)

type UmsMemberDao struct{}

// GetMembers 根据用户名手机号查询用户
func (t *UmsMemberDao) GetMembers(ctx context.Context, tx *gorm.DB, username, telephone string) (
	[]ums_member.UmsMember, error) {
	res := make([]ums_member.UmsMember, 0)
	err := tx.WithContext(ctx).Model(&ums_member.UmsMember{}).
		Where("username = ? OR telephone = ?", username, telephone).Find(&res).Error
	return res, err
}
