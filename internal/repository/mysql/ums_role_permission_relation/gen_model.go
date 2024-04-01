package ums_role_permission_relation

// UmsRolePermissionRelation 后台用户角色和权限关系表
//
//go:generate gormgen -structs UmsRolePermissionRelation -input .
type UmsRolePermissionRelation struct {
	Id           int64 //
	RoleId       int64 //
	PermissionId int64 //
}
