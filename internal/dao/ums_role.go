package dao

import (
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_menu"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"

	"gorm.io/gorm"
)

type UmsRoleDao struct{}

/**
 * 根据后台用户ID获取菜单
 */
func (t *UmsRoleDao) GetMenuList(tx *gorm.DB, adminId int64) ([]ums_menu.UmsMenu, error) {
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
func (t *UmsRoleDao) GetMenuListByRoleId(tx *gorm.DB, roleId int64) ([]ums_menu.UmsMenu, error) {
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
	 ums_role_menu_relation rmr
		LEFT JOIN ums_menu m ON rmr.menu_id = m.id
WHERE
	rmr.role_id = ?
  AND m.id IS NOT NULL
GROUP BY
	m.id
`
	err := tx.Raw(sql, roleId).Scan(&res).Error
	return res, err
}

/**
 * 根据角色ID获取资源
 */
func (t *UmsRoleDao) GetResourceListByRoleId(tx *gorm.DB, roleId int64) ([]ums_resource.UmsResource, error) {
	res := make([]ums_resource.UmsResource, 0)
	sql := `
SELECT
	r.id id,
	r.create_time create_time,
	r.name name,
	r.url url,
	r.description description,
	r.category_id category_id
FROM
	ums_role_resource_relation rrr
		LEFT JOIN ums_resource r ON rrr.resource_id = r.id
WHERE
	rrr.role_id = ?
  AND r.id IS NOT NULL
GROUP BY
	r.id
`
	err := tx.Raw(sql, roleId).Scan(&res).Error
	return res, err
}
