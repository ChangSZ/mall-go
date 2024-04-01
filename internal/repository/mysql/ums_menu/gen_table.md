#### mall.ums_menu 
后台菜单表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | parent_id | 父级ID | bigint |  | YES |  |  |
| 3 | create_time | 创建时间 | datetime |  | YES |  |  |
| 4 | title | 菜单名称 | varchar(100) |  | YES |  |  |
| 5 | level | 菜单级数 | int |  | YES |  |  |
| 6 | sort | 菜单排序 | int |  | YES |  |  |
| 7 | name | 前端名称 | varchar(100) |  | YES |  |  |
| 8 | icon | 前端图标 | varchar(200) |  | YES |  |  |
| 9 | hidden | 前端隐藏 | int |  | YES |  |  |
