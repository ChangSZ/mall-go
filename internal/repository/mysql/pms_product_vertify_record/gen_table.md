#### mall.pms_product_vertify_record 
商品审核记录

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | product_id |  | bigint |  | YES |  |  |
| 3 | create_time |  | datetime |  | YES |  |  |
| 4 | vertify_man | 审核人 | varchar(64) |  | YES |  |  |
| 5 | status |  | int |  | YES |  |  |
| 6 | detail | 反馈详情 | varchar(255) |  | YES |  |  |
