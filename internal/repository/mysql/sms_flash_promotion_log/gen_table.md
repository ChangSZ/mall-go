#### mall.sms_flash_promotion_log 
限时购通知记录

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | member_id |  | int |  | YES |  |  |
| 3 | product_id |  | bigint |  | YES |  |  |
| 4 | member_phone |  | varchar(64) |  | YES |  |  |
| 5 | product_name |  | varchar(100) |  | YES |  |  |
| 6 | subscribe_time | 会员订阅时间 | datetime |  | YES |  |  |
| 7 | send_time |  | datetime |  | YES |  |  |
