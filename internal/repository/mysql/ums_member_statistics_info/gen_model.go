package ums_member_statistics_info

import "time"

// UmsMemberStatisticsInfo 会员统计信息
//
//go:generate gormgen -structs UmsMemberStatisticsInfo -input .
type UmsMemberStatisticsInfo struct {
	Id                  int64     //
	MemberId            int64     //
	ConsumeAmount       float64   // 累计消费金额
	OrderCount          int32     // 订单数量
	CouponCount         int32     // 优惠券数量
	CommentCount        int32     // 评价数
	ReturnOrderCount    int32     // 退货数量
	LoginCount          int32     // 登录次数
	AttendCount         int32     // 关注数量
	FansCount           int32     // 粉丝数量
	CollectProductCount int32     //
	CollectSubjectCount int32     //
	CollectTopicCount   int32     //
	CollectCommentCount int32     //
	InviteFriendCount   int32     //
	RecentOrderTime     time.Time `gorm:"time"` // 最后一次下订单时间
}
