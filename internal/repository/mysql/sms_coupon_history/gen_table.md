#### mall.sms_coupon_history 
优惠券使用、领取历史表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | coupon_id |  | bigint | MUL | YES |  |  |
| 3 | member_id |  | bigint | MUL | YES |  |  |
| 4 | coupon_code |  | varchar(64) |  | YES |  |  |
| 5 | member_nickname | 领取人昵称 | varchar(64) |  | YES |  |  |
| 6 | get_type | 获取类型：0->后台赠送；1->主动获取 | int |  | YES |  |  |
| 7 | create_time |  | datetime |  | YES |  |  |
| 8 | use_status | 使用状态：0->未使用；1->已使用；2->已过期 | int |  | YES |  |  |
| 9 | use_time | 使用时间 | datetime |  | YES |  |  |
| 10 | order_id | 订单编号 | bigint |  | YES |  |  |
| 11 | order_sn | 订单号码 | varchar(100) |  | YES |  |  |
