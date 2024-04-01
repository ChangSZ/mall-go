#### mall.ums_member 
会员表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | member_level_id |  | bigint |  | YES |  |  |
| 3 | username | 用户名 | varchar(64) | UNI | YES |  |  |
| 4 | password | 密码 | varchar(64) |  | YES |  |  |
| 5 | nickname | 昵称 | varchar(64) |  | YES |  |  |
| 6 | phone | 手机号码 | varchar(64) | UNI | YES |  |  |
| 7 | status | 帐号启用状态:0->禁用；1->启用 | int |  | YES |  |  |
| 8 | create_time | 注册时间 | datetime |  | YES |  |  |
| 9 | icon | 头像 | varchar(500) |  | YES |  |  |
| 10 | gender | 性别：0->未知；1->男；2->女 | int |  | YES |  |  |
| 11 | birthday | 生日 | date |  | YES |  |  |
| 12 | city | 所做城市 | varchar(64) |  | YES |  |  |
| 13 | job | 职业 | varchar(100) |  | YES |  |  |
| 14 | personalized_signature | 个性签名 | varchar(200) |  | YES |  |  |
| 15 | source_type | 用户来源 | int |  | YES |  |  |
| 16 | integration | 积分 | int |  | YES |  |  |
| 17 | growth | 成长值 | int |  | YES |  |  |
| 18 | luckey_count | 剩余抽奖次数 | int |  | YES |  |  |
| 19 | history_integration | 历史积分数量 | int |  | YES |  |  |
