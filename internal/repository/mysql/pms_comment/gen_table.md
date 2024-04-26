#### mall.pms_comment 
商品评价表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | product_id |  | bigint |  | YES |  |  |
| 3 | member_nick_name |  | varchar(255) |  | YES |  |  |
| 4 | product_name |  | varchar(255) |  | YES |  |  |
| 5 | star | 评价星数：0->5 | int |  | YES |  |  |
| 6 | member_ip | 评价的ip | varchar(64) |  | YES |  |  |
| 7 | create_time |  | datetime |  | YES |  |  |
| 8 | show_status |  | int |  | YES |  |  |
| 9 | product_attribute | 购买时的商品属性 | varchar(255) |  | YES |  |  |
| 10 | collect_couont |  | int |  | YES |  |  |
| 11 | read_count |  | int |  | YES |  |  |
| 12 | content |  | text |  | YES |  |  |
| 13 | pics | 上传图片地址，以逗号隔开 | varchar(1000) |  | YES |  |  |
| 14 | member_icon | 评论用户头像 | varchar(255) |  | YES |  |  |
| 15 | replay_count |  | int |  | YES |  |  |
