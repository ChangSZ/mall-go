#### mall.pms_feight_template 
运费模版

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | name |  | varchar(64) |  | YES |  |  |
| 3 | charge_type | 计费类型:0->按重量；1->按件数 | int |  | YES |  |  |
| 4 | first_weight | 首重kg | decimal(10,2) |  | YES |  |  |
| 5 | first_fee | 首费（元） | decimal(10,2) |  | YES |  |  |
| 6 | continue_weight |  | decimal(10,2) |  | YES |  |  |
| 7 | continme_fee |  | decimal(10,2) |  | YES |  |  |
| 8 | dest | 目的地（省、市） | varchar(255) |  | YES |  |  |
