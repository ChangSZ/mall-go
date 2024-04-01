#### mall.ums_resource 
后台资源表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | create_time | 创建时间 | datetime |  | YES |  |  |
| 3 | name | 资源名称 | varchar(200) |  | YES |  |  |
| 4 | url | 资源URL | varchar(200) |  | YES |  |  |
| 5 | description | 描述 | varchar(500) |  | YES |  |  |
| 6 | category_id | 资源分类ID | bigint |  | YES |  |  |
