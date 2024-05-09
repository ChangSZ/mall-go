#### mall.oms_order_return_reason 
退货原因表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | name | 退货类型 | varchar(100) |  | YES |  |  |
| 3 | sort |  | int |  | YES |  |  |
| 4 | status | 状态：0->不启用；1->启用 | int |  | YES |  |  |
| 5 | create_time | 添加时间 | datetime |  | YES |  |  |
