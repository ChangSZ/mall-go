#### mall.sms_coupon_product_category_relation 
优惠券和产品分类关系表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | coupon_id |  | bigint |  | YES |  |  |
| 3 | product_category_id |  | bigint |  | YES |  |  |
| 4 | product_category_name | 产品分类名称 | varchar(200) |  | YES |  |  |
| 5 | parent_category_name | 父分类名称 | varchar(200) |  | YES |  |  |
