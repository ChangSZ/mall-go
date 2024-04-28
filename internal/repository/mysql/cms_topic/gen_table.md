#### mall.cms_topic 
话题表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | category_id |  | bigint |  | YES |  |  |
| 3 | name |  | varchar(255) |  | YES |  |  |
| 4 | create_time |  | datetime |  | YES |  |  |
| 5 | start_time |  | datetime |  | YES |  |  |
| 6 | end_time |  | datetime |  | YES |  |  |
| 7 | attend_count | 参与人数 | int |  | YES |  |  |
| 8 | attention_count | 关注人数 | int |  | YES |  |  |
| 9 | read_count |  | int |  | YES |  |  |
| 10 | award_name | 奖品名称 | varchar(100) |  | YES |  |  |
| 11 | attend_type | 参与方式 | varchar(100) |  | YES |  |  |
| 12 | content | 话题内容 | text |  | YES |  |  |
