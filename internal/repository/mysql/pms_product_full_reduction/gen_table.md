#### mall.pms_product_full_reduction 
产品满减表(只针对同商品)

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | product_id |  | bigint |  | YES |  |  |
| 3 | full_price |  | decimal(10,2) |  | YES |  |  |
| 4 | reduce_price |  | decimal(10,2) |  | YES |  |  |
