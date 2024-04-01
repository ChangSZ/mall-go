package ums_admin_permission_relation

// UmsAdminPermissionRelation 后台用户和权限关系表(除角色中定义的权限以外的加减权限)
//
//go:generate gormgen -structs UmsAdminPermissionRelation -input .
type UmsAdminPermissionRelation struct {
	Id           int64 //
	AdminId      int64 //
	PermissionId int64 //
	Type         int32 //
}
