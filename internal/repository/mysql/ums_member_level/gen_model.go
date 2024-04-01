package ums_member_level

// UmsMemberLevel 会员等级表
//
//go:generate gormgen -structs UmsMemberLevel -input .
type UmsMemberLevel struct {
	Id                    int64   //
	Name                  string  //
	GrowthPoint           int32   //
	DefaultStatus         int32   // 是否为默认等级：0->不是；1->是
	FreeFreightPoint      float64 // 免运费标准
	CommentGrowthPoint    int32   // 每次评价获取的成长值
	PriviledgeFreeFreight int32   // 是否有免邮特权
	PriviledgeSignIn      int32   // 是否有签到特权
	PriviledgeComment     int32   // 是否有评论获奖励特权
	PriviledgePromotion   int32   // 是否有专享活动特权
	PriviledgeMemberPrice int32   // 是否有会员价格特权
	PriviledgeBirthday    int32   // 是否有生日特权
	Note                  string  //
}
