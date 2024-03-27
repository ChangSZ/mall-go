#### mall.ums_role 
后台用户角色表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | name | 名称 | varchar(100) |  | YES |  |  |
| 3 | description | 描述 | varchar(500) |  | YES |  |  |
| 4 | admin_count | 后台用户数量 | int |  | YES |  |  |
| 5 | create_time | 创建时间 | datetime |  | YES |  |  |
| 6 | status | 启用状态：0->禁用；1->启用 | int |  | YES |  | 1 |
| 7 | sort |  | int |  | YES |  | 0 |
