#### mall.pms_product_attribute_value 
存储产品参数信息的表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | product_id |  | bigint |  | YES |  |  |
| 3 | product_attribute_id |  | bigint |  | YES |  |  |
| 4 | value | 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开 | varchar(64) |  | YES |  |  |
