#### mall.ums_member_task 
会员任务表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | name |  | varchar(100) |  | YES |  |  |
| 3 | growth | 赠送成长值 | int |  | YES |  |  |
| 4 | intergration | 赠送积分 | int |  | YES |  |  |
| 5 | type | 任务类型：0->新手任务；1->日常任务 | int |  | YES |  |  |
