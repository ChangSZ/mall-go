#### mall.sms_flash_promotion 
限时购表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | title | 秒杀时间段名称 | varchar(200) |  | YES |  |  |
| 3 | start_date | 开始日期 | date |  | YES |  |  |
| 4 | end_date | 结束日期 | date |  | YES |  |  |
| 5 | status | 上下线状态 | int |  | YES |  |  |
| 6 | create_time | 创建时间 | datetime |  | YES |  |  |
