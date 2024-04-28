#### mall.cms_subject 
专题表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | category_id |  | bigint |  | YES |  |  |
| 3 | title |  | varchar(100) |  | YES |  |  |
| 4 | pic | 专题主图 | varchar(500) |  | YES |  |  |
| 5 | product_count | 关联产品数量 | int |  | YES |  |  |
| 6 | recommend_status |  | int |  | YES |  |  |
| 7 | create_time |  | datetime |  | YES |  |  |
| 8 | collect_count |  | int |  | YES |  |  |
| 9 | read_count |  | int |  | YES |  |  |
| 10 | comment_count |  | int |  | YES |  |  |
| 11 | album_pics | 画册图片用逗号分割 | varchar(1000) |  | YES |  |  |
| 12 | description |  | varchar(1000) |  | YES |  |  |
| 13 | show_status | 显示状态：0->不显示；1->显示 | int |  | YES |  |  |
| 14 | content |  | text |  | YES |  |  |
| 15 | forward_count | 转发数 | int |  | YES |  |  |
| 16 | category_name | 专题分类名称 | varchar(200) |  | YES |  |  |
