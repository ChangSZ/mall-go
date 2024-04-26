#### mall.pms_product_category_attribute_relation 
产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | product_category_id |  | bigint |  | YES |  |  |
| 3 | product_attribute_id |  | bigint |  | YES |  |  |
