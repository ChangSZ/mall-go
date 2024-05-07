#### mall.sms_flash_promotion_session 
限时购场次表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 编号 | bigint | PRI | NO | auto_increment |  |
| 2 | name | 场次名称 | varchar(200) |  | YES |  |  |
| 3 | start_time | 每日开始时间 | time |  | YES |  |  |
| 4 | end_time | 每日结束时间 | time |  | YES |  |  |
| 5 | status | 启用状态：0->不启用；1->启用 | int |  | YES |  |  |
| 6 | create_time | 创建时间 | datetime |  | YES |  |  |
