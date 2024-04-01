#### mall.ums_member_receive_address 
会员收货地址表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | member_id |  | bigint |  | YES |  |  |
| 3 | name | 收货人名称 | varchar(100) |  | YES |  |  |
| 4 | phone_number |  | varchar(64) |  | YES |  |  |
| 5 | default_status | 是否为默认 | int |  | YES |  |  |
| 6 | post_code | 邮政编码 | varchar(100) |  | YES |  |  |
| 7 | province | 省份/直辖市 | varchar(100) |  | YES |  |  |
| 8 | city | 城市 | varchar(100) |  | YES |  |  |
| 9 | region | 区 | varchar(100) |  | YES |  |  |
| 10 | detail_address | 详细地址(街道) | varchar(128) |  | YES |  |  |
