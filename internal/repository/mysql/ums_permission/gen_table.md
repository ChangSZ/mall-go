#### mall.ums_permission 
后台用户权限表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | pid | 父级权限id | bigint |  | YES |  |  |
| 3 | name | 名称 | varchar(100) |  | YES |  |  |
| 4 | value | 权限值 | varchar(200) |  | YES |  |  |
| 5 | icon | 图标 | varchar(500) |  | YES |  |  |
| 6 | type | 权限类型：0->目录；1->菜单；2->按钮（接口绑定权限） | int |  | YES |  |  |
| 7 | uri | 前端资源路径 | varchar(200) |  | YES |  |  |
| 8 | status | 启用状态；0->禁用；1->启用 | int |  | YES |  |  |
| 9 | create_time | 创建时间 | datetime |  | YES |  |  |
| 10 | sort | 排序 | int |  | YES |  |  |
