package ums_member_member_tag_relation

// UmsMemberMemberTagRelation 用户和标签关系表
//
//go:generate gormgen -structs UmsMemberMemberTagRelation -input .
type UmsMemberMemberTagRelation struct {
	Id       int64 //
	MemberId int64 //
	TagId    int64 //
}
