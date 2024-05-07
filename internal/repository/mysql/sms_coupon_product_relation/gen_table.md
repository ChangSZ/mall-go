#### mall.sms_coupon_product_relation 
优惠券和产品的关系表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | coupon_id |  | bigint |  | YES |  |  |
| 3 | product_id |  | bigint |  | YES |  |  |
| 4 | product_name | 商品名称 | varchar(500) |  | YES |  |  |
| 5 | product_sn | 商品编码 | varchar(200) |  | YES |  |  |
