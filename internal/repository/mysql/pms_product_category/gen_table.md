#### mall.pms_product_category 
产品分类

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | parent_id | 上机分类的编号：0表示一级分类 | bigint |  | YES |  |  |
| 3 | name |  | varchar(64) |  | YES |  |  |
| 4 | level | 分类级别：0->1级；1->2级 | int |  | YES |  |  |
| 5 | product_count |  | int |  | YES |  |  |
| 6 | product_unit |  | varchar(64) |  | YES |  |  |
| 7 | nav_status | 是否显示在导航栏：0->不显示；1->显示 | int |  | YES |  |  |
| 8 | show_status | 显示状态：0->不显示；1->显示 | int |  | YES |  |  |
| 9 | sort |  | int |  | YES |  |  |
| 10 | icon | 图标 | varchar(255) |  | YES |  |  |
| 11 | keywords |  | varchar(255) |  | YES |  |  |
| 12 | description | 描述 | text |  | YES |  |  |
