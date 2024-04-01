#### mall.ums_integration_consume_setting 
积分消费设置

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | deduction_per_amount | 每一元需要抵扣的积分数量 | int |  | YES |  |  |
| 3 | max_percent_per_order | 每笔订单最高抵用百分比 | int |  | YES |  |  |
| 4 | use_unit | 每次使用积分最小单位100 | int |  | YES |  |  |
| 5 | coupon_status | 是否可以和优惠券同用；0->不可以；1->可以 | int |  | YES |  |  |
