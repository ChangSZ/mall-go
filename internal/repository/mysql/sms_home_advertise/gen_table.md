#### mall.sms_home_advertise 
首页轮播广告表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | name |  | varchar(100) |  | YES |  |  |
| 3 | type | 轮播位置：0->PC首页轮播；1->app首页轮播 | int |  | YES |  |  |
| 4 | pic |  | varchar(500) |  | YES |  |  |
| 5 | start_time |  | datetime |  | YES |  |  |
| 6 | end_time |  | datetime |  | YES |  |  |
| 7 | status | 上下线状态：0->下线；1->上线 | int |  | YES |  |  |
| 8 | click_count | 点击数 | int |  | YES |  |  |
| 9 | order_count | 下单数 | int |  | YES |  |  |
| 10 | url | 链接地址 | varchar(500) |  | YES |  |  |
| 11 | note | 备注 | varchar(500) |  | YES |  |  |
| 12 | sort | 排序 | int |  | YES |  | 0 |
