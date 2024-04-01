#### mall.ums_admin_login_log 
后台用户登录日志表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | admin_id |  | bigint |  | YES |  |  |
| 3 | create_time |  | datetime |  | YES |  |  |
| 4 | ip |  | varchar(64) |  | YES |  |  |
| 5 | address |  | varchar(100) |  | YES |  |  |
| 6 | user_agent | 浏览器登录类型 | varchar(100) |  | YES |  |  |
