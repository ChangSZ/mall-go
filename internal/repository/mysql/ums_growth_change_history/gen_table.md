#### mall.ums_growth_change_history 
成长值变化历史记录表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | member_id |  | bigint |  | YES |  |  |
| 3 | create_time |  | datetime |  | YES |  |  |
| 4 | change_type | 改变类型：0->增加；1->减少 | int |  | YES |  |  |
| 5 | change_count | 积分改变数量 | int |  | YES |  |  |
| 6 | operate_man | 操作人员 | varchar(100) |  | YES |  |  |
| 7 | operate_note | 操作备注 | varchar(200) |  | YES |  |  |
| 8 | source_type | 积分来源：0->购物；1->管理员修改 | int |  | YES |  |  |
