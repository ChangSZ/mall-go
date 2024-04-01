#### mall.ums_member_statistics_info 
会员统计信息

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | member_id |  | bigint |  | YES |  |  |
| 3 | consume_amount | 累计消费金额 | decimal(10,2) |  | YES |  |  |
| 4 | order_count | 订单数量 | int |  | YES |  |  |
| 5 | coupon_count | 优惠券数量 | int |  | YES |  |  |
| 6 | comment_count | 评价数 | int |  | YES |  |  |
| 7 | return_order_count | 退货数量 | int |  | YES |  |  |
| 8 | login_count | 登录次数 | int |  | YES |  |  |
| 9 | attend_count | 关注数量 | int |  | YES |  |  |
| 10 | fans_count | 粉丝数量 | int |  | YES |  |  |
| 11 | collect_product_count |  | int |  | YES |  |  |
| 12 | collect_subject_count |  | int |  | YES |  |  |
| 13 | collect_topic_count |  | int |  | YES |  |  |
| 14 | collect_comment_count |  | int |  | YES |  |  |
| 15 | invite_friend_count |  | int |  | YES |  |  |
| 16 | recent_order_time | 最后一次下订单时间 | datetime |  | YES |  |  |
