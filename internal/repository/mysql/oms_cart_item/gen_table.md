#### mall.oms_cart_item 
购物车表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | product_id |  | bigint |  | YES |  |  |
| 3 | product_sku_id |  | bigint |  | YES |  |  |
| 4 | member_id |  | bigint |  | YES |  |  |
| 5 | quantity | 购买数量 | int |  | YES |  |  |
| 6 | price | 添加到购物车的价格 | decimal(10,2) |  | YES |  |  |
| 7 | product_pic | 商品主图 | varchar(1000) |  | YES |  |  |
| 8 | product_name | 商品名称 | varchar(500) |  | YES |  |  |
| 9 | product_sub_title | 商品副标题（卖点） | varchar(500) |  | YES |  |  |
| 10 | product_sku_code | 商品sku条码 | varchar(200) |  | YES |  |  |
| 11 | member_nickname | 会员昵称 | varchar(500) |  | YES |  |  |
| 12 | create_date | 创建时间 | datetime |  | YES |  |  |
| 13 | modify_date | 修改时间 | datetime |  | YES |  |  |
| 14 | delete_status | 是否删除 | int |  | YES |  | 0 |
| 15 | product_category_id | 商品分类 | bigint |  | YES |  |  |
| 16 | product_brand |  | varchar(200) |  | YES |  |  |
| 17 | product_sn |  | varchar(200) |  | YES |  |  |
| 18 | product_attr | 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}] | varchar(500) |  | YES |  |  |
