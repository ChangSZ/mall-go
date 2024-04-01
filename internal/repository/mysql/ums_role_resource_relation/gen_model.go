package ums_role_resource_relation

// UmsRoleResourceRelation 后台角色资源关系表
//
//go:generate gormgen -structs UmsRoleResourceRelation -input .
type UmsRoleResourceRelation struct {
	Id         int64 //
	RoleId     int64 // 角色ID
	ResourceId int64 // 资源ID
}
