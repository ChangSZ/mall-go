#### mall.pms_product_ladder 
产品阶梯价格表(只针对同商品)

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | product_id |  | bigint |  | YES |  |  |
| 3 | count | 满足的商品数量 | int |  | YES |  |  |
| 4 | discount | 折扣 | decimal(10,2) |  | YES |  |  |
| 5 | price | 折后价格 | decimal(10,2) |  | YES |  |  |
