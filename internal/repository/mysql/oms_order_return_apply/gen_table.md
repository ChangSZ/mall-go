#### mall.oms_order_return_apply 
订单退货申请

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | order_id | 订单id | bigint |  | YES |  |  |
| 3 | company_address_id | 收货地址表id | bigint |  | YES |  |  |
| 4 | product_id | 退货商品id | bigint |  | YES |  |  |
| 5 | order_sn | 订单编号 | varchar(64) |  | YES |  |  |
| 6 | create_time | 申请时间 | datetime |  | YES |  |  |
| 7 | member_username | 会员用户名 | varchar(64) |  | YES |  |  |
| 8 | return_amount | 退款金额 | decimal(10,2) |  | YES |  |  |
| 9 | return_name | 退货人姓名 | varchar(100) |  | YES |  |  |
| 10 | return_phone | 退货人电话 | varchar(100) |  | YES |  |  |
| 11 | status | 申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝 | int |  | YES |  |  |
| 12 | handle_time | 处理时间 | datetime |  | YES |  |  |
| 13 | product_pic | 商品图片 | varchar(500) |  | YES |  |  |
| 14 | product_name | 商品名称 | varchar(200) |  | YES |  |  |
| 15 | product_brand | 商品品牌 | varchar(200) |  | YES |  |  |
| 16 | product_attr | 商品销售属性：颜色：红色；尺码：xl; | varchar(500) |  | YES |  |  |
| 17 | product_count | 退货数量 | int |  | YES |  |  |
| 18 | product_price | 商品单价 | decimal(10,2) |  | YES |  |  |
| 19 | product_real_price | 商品实际支付单价 | decimal(10,2) |  | YES |  |  |
| 20 | reason | 原因 | varchar(200) |  | YES |  |  |
| 21 | description | 描述 | varchar(500) |  | YES |  |  |
| 22 | proof_pics | 凭证图片，以逗号隔开 | varchar(1000) |  | YES |  |  |
| 23 | handle_note | 处理备注 | varchar(500) |  | YES |  |  |
| 24 | handle_man | 处理人员 | varchar(100) |  | YES |  |  |
| 25 | receive_man | 收货人 | varchar(100) |  | YES |  |  |
| 26 | receive_time | 收货时间 | datetime |  | YES |  |  |
| 27 | receive_note | 收货备注 | varchar(500) |  | YES |  |  |
