#### mall.sms_flash_promotion_product_relation 
商品限时购与商品关系表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 编号 | bigint | PRI | NO | auto_increment |  |
| 2 | flash_promotion_id |  | bigint |  | YES |  |  |
| 3 | flash_promotion_session_id | 编号 | bigint |  | YES |  |  |
| 4 | product_id |  | bigint |  | YES |  |  |
| 5 | flash_promotion_price | 限时购价格 | decimal(10,2) |  | YES |  |  |
| 6 | flash_promotion_count | 限时购数量 | int |  | YES |  |  |
| 7 | flash_promotion_limit | 每人限购数量 | int |  | YES |  |  |
| 8 | sort | 排序 | int |  | YES |  |  |
