#### mall.oms_company_address 
公司收发货地址表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | address_name | 地址名称 | varchar(200) |  | YES |  |  |
| 3 | send_status | 默认发货地址：0->否；1->是 | int |  | YES |  |  |
| 4 | receive_status | 是否默认收货地址：0->否；1->是 | int |  | YES |  |  |
| 5 | name | 收发货人姓名 | varchar(64) |  | YES |  |  |
| 6 | phone | 收货人电话 | varchar(64) |  | YES |  |  |
| 7 | province | 省/直辖市 | varchar(64) |  | YES |  |  |
| 8 | city | 市 | varchar(64) |  | YES |  |  |
| 9 | region | 区 | varchar(64) |  | YES |  |  |
| 10 | detail_address | 详细地址 | varchar(200) |  | YES |  |  |
