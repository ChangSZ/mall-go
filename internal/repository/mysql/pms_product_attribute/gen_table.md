#### mall.pms_product_attribute 
商品属性参数表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | product_attribute_category_id |  | bigint |  | YES |  |  |
| 3 | name |  | varchar(64) |  | YES |  |  |
| 4 | select_type | 属性选择类型：0->唯一；1->单选；2->多选 | int |  | YES |  |  |
| 5 | input_type | 属性录入方式：0->手工录入；1->从列表中选取 | int |  | YES |  |  |
| 6 | input_list | 可选值列表，以逗号隔开 | varchar(255) |  | YES |  |  |
| 7 | sort | 排序字段：最高的可以单独上传图片 | int |  | YES |  |  |
| 8 | filter_type | 分类筛选样式：1->普通；1->颜色 | int |  | YES |  |  |
| 9 | search_type | 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索 | int |  | YES |  |  |
| 10 | related_status | 相同属性产品是否关联；0->不关联；1->关联 | int |  | YES |  |  |
| 11 | hand_add_status | 是否支持手动新增；0->不支持；1->支持 | int |  | YES |  |  |
| 12 | type | 属性的类型；0->规格；1->参数 | int |  | YES |  |  |
