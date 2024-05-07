#### mall.sms_coupon 
优惠券表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | type | 优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券 | int |  | YES |  |  |
| 3 | name |  | varchar(100) |  | YES |  |  |
| 4 | platform | 使用平台：0->全部；1->移动；2->PC | int |  | YES |  |  |
| 5 | count | 数量 | int |  | YES |  |  |
| 6 | amount | 金额 | decimal(10,2) |  | YES |  |  |
| 7 | per_limit | 每人限领张数 | int |  | YES |  |  |
| 8 | min_point | 使用门槛；0表示无门槛 | decimal(10,2) |  | YES |  |  |
| 9 | start_time |  | datetime |  | YES |  |  |
| 10 | end_time |  | datetime |  | YES |  |  |
| 11 | use_type | 使用类型：0->全场通用；1->指定分类；2->指定商品 | int |  | YES |  |  |
| 12 | note | 备注 | varchar(200) |  | YES |  |  |
| 13 | publish_count | 发行数量 | int |  | YES |  |  |
| 14 | use_count | 已使用数量 | int |  | YES |  |  |
| 15 | receive_count | 领取数量 | int |  | YES |  |  |
| 16 | enable_time | 可以领取的日期 | datetime |  | YES |  |  |
| 17 | code | 优惠码 | varchar(64) |  | YES |  |  |
| 18 | member_level | 可领取的会员类型：0->无限时 | int |  | YES |  |  |
