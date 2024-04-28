#### mall.cms_topic_comment 
专题评论表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | member_nick_name |  | varchar(255) |  | YES |  |  |
| 3 | topic_id |  | bigint |  | YES |  |  |
| 4 | member_icon |  | varchar(255) |  | YES |  |  |
| 5 | content |  | varchar(1000) |  | YES |  |  |
| 6 | create_time |  | datetime |  | YES |  |  |
| 7 | show_status |  | int |  | YES |  |  |
