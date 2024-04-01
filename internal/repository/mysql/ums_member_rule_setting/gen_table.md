#### mall.ums_member_rule_setting 
会员积分成长规则表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | continue_sign_day | 连续签到天数 | int |  | YES |  |  |
| 3 | continue_sign_point | 连续签到赠送数量 | int |  | YES |  |  |
| 4 | consume_per_point | 每消费多少元获取1个点 | decimal(10,2) |  | YES |  |  |
| 5 | low_order_amount | 最低获取点数的订单金额 | decimal(10,2) |  | YES |  |  |
| 6 | max_point_per_order | 每笔订单最高获取点数 | int |  | YES |  |  |
| 7 | type | 类型：0->积分规则；1->成长值规则 | int |  | YES |  |  |
