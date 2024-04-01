package ums_admin_role_relation

// UmsAdminRoleRelation 后台用户和角色关系表
//
//go:generate gormgen -structs UmsAdminRoleRelation -input .
type UmsAdminRoleRelation struct {
	Id      int64 //
	AdminId int64 //
	RoleId  int64 //
}
