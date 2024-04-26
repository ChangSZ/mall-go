#### mall.pms_sku_stock 
sku的库存

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | product_id |  | bigint |  | YES |  |  |
| 3 | sku_code | sku编码 | varchar(64) |  | NO |  |  |
| 4 | price |  | decimal(10,2) |  | YES |  |  |
| 5 | stock | 库存 | int |  | YES |  | 0 |
| 6 | low_stock | 预警库存 | int |  | YES |  |  |
| 7 | pic | 展示图片 | varchar(255) |  | YES |  |  |
| 8 | sale | 销量 | int |  | YES |  |  |
| 9 | promotion_price | 单品促销价格 | decimal(10,2) |  | YES |  |  |
| 10 | lock_stock | 锁定库存 | int |  | YES |  | 0 |
| 11 | sp_data | 商品销售属性，json格式 | varchar(500) |  | YES |  |  |
