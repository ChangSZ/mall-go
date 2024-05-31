package dto

import "time"

type UmsMember struct {
	Id                    int64     `json:"id"`                    //
	MemberLevelId         int64     `json:"memberLevelId"`         //
	Username              string    `json:"username"`              // 用户名
	Password              string    `json:"password"`              // 密码
	Nickname              string    `json:"nickname"`              // 昵称
	Phone                 string    `json:"phone"`                 // 手机号码
	Status                int32     `json:"status"`                // 帐号启用状态:0->禁用；1->启用
	CreateTime            time.Time `json:"createTime"`            // 注册时间
	Icon                  string    `json:"icon"`                  // 头像
	Gender                int32     `json:"gender"`                // 性别：0->未知；1->男；2->女
	Birthday              string    `json:"birthday"`              // 生日
	City                  string    `json:"city"`                  // 所做城市
	Job                   string    `json:"job"`                   // 职业
	PersonalizedSignature string    `json:"personalizedSignature"` // 个性签名
	Integration           int32     `json:"integration"`           // 积分
	Growth                int32     `json:"growth"`                // 成长值
}
