#### mall.pms_product_operate_log 

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | product_id |  | bigint |  | YES |  |  |
| 3 | price_old |  | decimal(10,2) |  | YES |  |  |
| 4 | price_new |  | decimal(10,2) |  | YES |  |  |
| 5 | sale_price_old |  | decimal(10,2) |  | YES |  |  |
| 6 | sale_price_new |  | decimal(10,2) |  | YES |  |  |
| 7 | gift_point_old | 赠送的积分 | int |  | YES |  |  |
| 8 | gift_point_new |  | int |  | YES |  |  |
| 9 | use_point_limit_old |  | int |  | YES |  |  |
| 10 | use_point_limit_new |  | int |  | YES |  |  |
| 11 | operate_man | 操作人 | varchar(64) |  | YES |  |  |
| 12 | create_time |  | datetime |  | YES |  |  |
