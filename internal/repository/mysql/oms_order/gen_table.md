#### mall.oms_order 
订单表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 订单id | bigint | PRI | NO | auto_increment |  |
| 2 | member_id |  | bigint |  | NO |  |  |
| 3 | coupon_id |  | bigint |  | YES |  |  |
| 4 | order_sn | 订单编号 | varchar(64) |  | YES |  |  |
| 5 | create_time | 提交时间 | datetime |  | YES |  |  |
| 6 | member_username | 用户帐号 | varchar(64) |  | YES |  |  |
| 7 | total_amount | 订单总金额 | decimal(10,2) |  | YES |  |  |
| 8 | pay_amount | 应付金额（实际支付金额） | decimal(10,2) |  | YES |  |  |
| 9 | freight_amount | 运费金额 | decimal(10,2) |  | YES |  |  |
| 10 | promotion_amount | 促销优化金额（促销价、满减、阶梯价） | decimal(10,2) |  | YES |  |  |
| 11 | integration_amount | 积分抵扣金额 | decimal(10,2) |  | YES |  |  |
| 12 | coupon_amount | 优惠券抵扣金额 | decimal(10,2) |  | YES |  |  |
| 13 | discount_amount | 管理员后台调整订单使用的折扣金额 | decimal(10,2) |  | YES |  |  |
| 14 | pay_type | 支付方式：0->未支付；1->支付宝；2->微信 | int |  | YES |  |  |
| 15 | source_type | 订单来源：0->PC订单；1->app订单 | int |  | YES |  |  |
| 16 | status | 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单 | int |  | YES |  |  |
| 17 | order_type | 订单类型：0->正常订单；1->秒杀订单 | int |  | YES |  |  |
| 18 | delivery_company | 物流公司(配送方式) | varchar(64) |  | YES |  |  |
| 19 | delivery_sn | 物流单号 | varchar(64) |  | YES |  |  |
| 20 | auto_confirm_day | 自动确认时间（天） | int |  | YES |  |  |
| 21 | integration | 可以获得的积分 | int |  | YES |  |  |
| 22 | growth | 可以活动的成长值 | int |  | YES |  |  |
| 23 | promotion_info | 活动信息 | varchar(100) |  | YES |  |  |
| 24 | bill_type | 发票类型：0->不开发票；1->电子发票；2->纸质发票 | int |  | YES |  |  |
| 25 | bill_header | 发票抬头 | varchar(200) |  | YES |  |  |
| 26 | bill_content | 发票内容 | varchar(200) |  | YES |  |  |
| 27 | bill_receiver_phone | 收票人电话 | varchar(32) |  | YES |  |  |
| 28 | bill_receiver_email | 收票人邮箱 | varchar(64) |  | YES |  |  |
| 29 | receiver_name | 收货人姓名 | varchar(100) |  | NO |  |  |
| 30 | receiver_phone | 收货人电话 | varchar(32) |  | NO |  |  |
| 31 | receiver_post_code | 收货人邮编 | varchar(32) |  | YES |  |  |
| 32 | receiver_province | 省份/直辖市 | varchar(32) |  | YES |  |  |
| 33 | receiver_city | 城市 | varchar(32) |  | YES |  |  |
| 34 | receiver_region | 区 | varchar(32) |  | YES |  |  |
| 35 | receiver_detail_address | 详细地址 | varchar(200) |  | YES |  |  |
| 36 | note | 订单备注 | varchar(500) |  | YES |  |  |
| 37 | confirm_status | 确认收货状态：0->未确认；1->已确认 | int |  | YES |  |  |
| 38 | delete_status | 删除状态：0->未删除；1->已删除 | int |  | NO |  | 0 |
| 39 | use_integration | 下单时使用的积分 | int |  | YES |  |  |
| 40 | payment_time | 支付时间 | datetime |  | YES |  |  |
| 41 | delivery_time | 发货时间 | datetime |  | YES |  |  |
| 42 | receive_time | 确认收货时间 | datetime |  | YES |  |  |
| 43 | comment_time | 评价时间 | datetime |  | YES |  |  |
| 44 | modify_time | 修改时间 | datetime |  | YES |  |  |
