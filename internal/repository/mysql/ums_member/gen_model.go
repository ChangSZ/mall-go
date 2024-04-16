package ums_member

import "time"

// UmsMember 会员表
//
//go:generate gormgen -structs UmsMember -input .
type UmsMember struct {
	Id                    int64     //
	MemberLevelId         int64     //
	Username              string    // 用户名
	Password              string    // 密码
	Nickname              string    // 昵称
	Phone                 string    // 手机号码
	Status                int32     // 帐号启用状态:0->禁用；1->启用
	CreateTime            time.Time `gorm:"autoCreateTime"` // 注册时间
	Icon                  string    // 头像
	Gender                int32     // 性别：0->未知；1->男；2->女
	Birthday              string    // 生日
	City                  string    // 所做城市
	Job                   string    // 职业
	PersonalizedSignature string    // 个性签名
	SourceType            int32     // 用户来源
	Integration           int32     // 积分
	Growth                int32     // 成长值
	LuckeyCount           int32     // 剩余抽奖次数
	HistoryIntegration    int32     // 历史积分数量
}
