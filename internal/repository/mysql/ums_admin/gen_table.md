#### mall.ums_admin 
后台用户表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | username |  | varchar(64) |  | YES |  |  |
| 3 | password |  | varchar(64) |  | YES |  |  |
| 4 | icon | 头像 | varchar(500) |  | YES |  |  |
| 5 | email | 邮箱 | varchar(100) |  | YES |  |  |
| 6 | nick_name | 昵称 | varchar(200) |  | YES |  |  |
| 7 | note | 备注信息 | varchar(500) |  | YES |  |  |
| 8 | create_time | 创建时间 | datetime |  | YES |  |  |
| 9 | login_time | 最后登录时间 | datetime |  | YES |  |  |
| 10 | status | 帐号启用状态：0->禁用；1->启用 | int |  | YES |  | 1 |
