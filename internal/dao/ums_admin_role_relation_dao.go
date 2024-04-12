package dao

import (
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin_role_relation"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_role"

	"gorm.io/gorm"
)

// 批量插入用户角色关系
func InsertList(tx *gorm.DB, adminRoleRelationList []ums_admin_role_relation.UmsAdminRoleRelation) int {
	return 0
}

// 获取用户所有角色
func GetRoleList(tx *gorm.DB, adminId int64) ([]ums_role.UmsRole, error) {
	res := make([]ums_role.UmsRole, 0)
	err := tx.Raw(`select r.* from ums_admin_role_relation ar left join ums_role r 
		on ar.role_id = r.id where ar.admin_id = ?`, adminId).Scan(&res).Error
	return res, err
}

// 获取用户所有可访问资源
func GetResourceList(tx *gorm.DB, adminId int64) []ums_resource.UmsResource {
	return nil
}

// 获取资源相关用户ID列表
func GetAdminIdList(tx *gorm.DB, resourceId int64) []int64 {
	return nil
}
