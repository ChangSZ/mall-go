#### mall.pms_comment_replay 
产品评价回复表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | comment_id |  | bigint |  | YES |  |  |
| 3 | member_nick_name |  | varchar(255) |  | YES |  |  |
| 4 | member_icon |  | varchar(255) |  | YES |  |  |
| 5 | content |  | varchar(1000) |  | YES |  |  |
| 6 | create_time |  | datetime |  | YES |  |  |
| 7 | type | 评论人员类型；0->会员；1->管理员 | int |  | YES |  |  |
