#### mall.oms_order_setting 
订单设置表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | flash_order_overtime | 秒杀订单超时关闭时间(分) | int |  | YES |  |  |
| 3 | normal_order_overtime | 正常订单超时时间(分) | int |  | YES |  |  |
| 4 | confirm_overtime | 发货后自动确认收货时间（天） | int |  | YES |  |  |
| 5 | finish_overtime | 自动完成交易时间，不能申请售后（天） | int |  | YES |  |  |
| 6 | comment_overtime | 订单完成后自动好评时间（天） | int |  | YES |  |  |
