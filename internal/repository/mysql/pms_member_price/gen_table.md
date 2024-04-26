#### mall.pms_member_price 
商品会员价格表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | product_id |  | bigint |  | YES |  |  |
| 3 | member_level_id |  | bigint |  | YES |  |  |
| 4 | member_price | 会员价格 | decimal(10,2) |  | YES |  |  |
| 5 | member_level_name |  | varchar(100) |  | YES |  |  |
