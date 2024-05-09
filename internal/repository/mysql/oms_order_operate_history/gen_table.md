#### mall.oms_order_operate_history 
订单操作历史记录

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | order_id | 订单id | bigint |  | YES |  |  |
| 3 | operate_man | 操作人：用户；系统；后台管理员 | varchar(100) |  | YES |  |  |
| 4 | create_time | 操作时间 | datetime |  | YES |  |  |
| 5 | order_status | 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单 | int |  | YES |  |  |
| 6 | note | 备注 | varchar(500) |  | YES |  |  |
