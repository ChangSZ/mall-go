#### mall.oms_order_item 
订单中所包含的商品

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | order_id | 订单id | bigint |  | YES |  |  |
| 3 | order_sn | 订单编号 | varchar(64) |  | YES |  |  |
| 4 | product_id |  | bigint |  | YES |  |  |
| 5 | product_pic |  | varchar(500) |  | YES |  |  |
| 6 | product_name |  | varchar(200) |  | YES |  |  |
| 7 | product_brand |  | varchar(200) |  | YES |  |  |
| 8 | product_sn |  | varchar(64) |  | YES |  |  |
| 9 | product_price | 销售价格 | decimal(10,2) |  | YES |  |  |
| 10 | product_quantity | 购买数量 | int |  | YES |  |  |
| 11 | product_sku_id | 商品sku编号 | bigint |  | YES |  |  |
| 12 | product_sku_code | 商品sku条码 | varchar(50) |  | YES |  |  |
| 13 | product_category_id | 商品分类id | bigint |  | YES |  |  |
| 14 | promotion_name | 商品促销名称 | varchar(200) |  | YES |  |  |
| 15 | promotion_amount | 商品促销分解金额 | decimal(10,2) |  | YES |  |  |
| 16 | coupon_amount | 优惠券优惠分解金额 | decimal(10,2) |  | YES |  |  |
| 17 | integration_amount | 积分优惠分解金额 | decimal(10,2) |  | YES |  |  |
| 18 | real_amount | 该商品经过优惠后的分解金额 | decimal(10,2) |  | YES |  |  |
| 19 | gift_integration |  | int |  | YES |  | 0 |
| 20 | gift_growth |  | int |  | YES |  | 0 |
| 21 | product_attr | 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}] | varchar(500) |  | YES |  |  |
