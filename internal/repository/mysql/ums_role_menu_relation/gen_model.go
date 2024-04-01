package ums_role_menu_relation

// UmsRoleMenuRelation 后台角色菜单关系表
//
//go:generate gormgen -structs UmsRoleMenuRelation -input .
type UmsRoleMenuRelation struct {
	Id     int64 //
	RoleId int64 // 角色ID
	MenuId int64 // 菜单ID
}
