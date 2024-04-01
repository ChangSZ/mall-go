package ums_member_rule_setting

// UmsMemberRuleSetting 会员积分成长规则表
//
//go:generate gormgen -structs UmsMemberRuleSetting -input .
type UmsMemberRuleSetting struct {
	Id                int64   //
	ContinueSignDay   int32   // 连续签到天数
	ContinueSignPoint int32   // 连续签到赠送数量
	ConsumePerPoint   float64 // 每消费多少元获取1个点
	LowOrderAmount    float64 // 最低获取点数的订单金额
	MaxPointPerOrder  int32   // 每笔订单最高获取点数
	Type              int32   // 类型：0->积分规则；1->成长值规则
}
