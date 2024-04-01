#### mall.ums_member_login_log 
会员登录记录

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | member_id |  | bigint |  | YES |  |  |
| 3 | create_time |  | datetime |  | YES |  |  |
| 4 | ip |  | varchar(64) |  | YES |  |  |
| 5 | city |  | varchar(64) |  | YES |  |  |
| 6 | login_type | 登录类型：0->PC；1->android;2->ios;3->小程序 | int |  | YES |  |  |
| 7 | province |  | varchar(64) |  | YES |  |  |
