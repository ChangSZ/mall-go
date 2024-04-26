#### mall.pms_brand 
品牌表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | name |  | varchar(64) |  | YES |  |  |
| 3 | first_letter | 首字母 | varchar(8) |  | YES |  |  |
| 4 | sort |  | int |  | YES |  |  |
| 5 | factory_status | 是否为品牌制造商：0->不是；1->是 | int |  | YES |  |  |
| 6 | show_status |  | int |  | YES |  |  |
| 7 | product_count | 产品数量 | int |  | YES |  |  |
| 8 | product_comment_count | 产品评论数量 | int |  | YES |  |  |
| 9 | logo | 品牌logo | varchar(255) |  | YES |  |  |
| 10 | big_pic | 专区大图 | varchar(255) |  | YES |  |  |
| 11 | brand_story | 品牌故事 | text |  | YES |  |  |
