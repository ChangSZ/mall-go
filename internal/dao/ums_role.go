package dao

import (
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_menu"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
	"gorm.io/gorm"
)

/**
 * 根据后台用户ID获取菜单
 */
func GetMenuList(tx *gorm.DB, adminId int64) ([]ums_menu.UmsMenu, error) {
	res := make([]ums_menu.UmsMenu, 0)
	sql := `
SELECT
m.id id,
m.parent_id parent_id,
m.create_time create_time,
m.title title,
m.level level,
m.sort sort,
m.name name,
m.icon icon,
m.hidden hidden
FROM
ums_admin_role_relation arr
	LEFT JOIN ums_role r ON arr.role_id = r.id
	LEFT JOIN ums_role_menu_relation rmr ON r.id = rmr.role_id
	LEFT JOIN ums_menu m ON rmr.menu_id = m.id
WHERE
arr.admin_id = ?
AND m.id IS NOT NULL
GROUP BY
m.id
`
	err := tx.Raw(sql, adminId).Scan(&res).Error
	return res, err
}

/**
 * 根据角色ID获取菜单
 */
func GetMenuListByRoleId(roleId int64) []ums_menu.UmsMenu {
	return nil
}

/**
 * 根据角色ID获取资源
 */
func GetResourceListByRoleId(roleId int64) []ums_resource.UmsResource {
	return nil
}
