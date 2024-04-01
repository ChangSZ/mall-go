#### mall.ums_member_level 
会员等级表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | name |  | varchar(100) |  | YES |  |  |
| 3 | growth_point |  | int |  | YES |  |  |
| 4 | default_status | 是否为默认等级：0->不是；1->是 | int |  | YES |  |  |
| 5 | free_freight_point | 免运费标准 | decimal(10,2) |  | YES |  |  |
| 6 | comment_growth_point | 每次评价获取的成长值 | int |  | YES |  |  |
| 7 | priviledge_free_freight | 是否有免邮特权 | int |  | YES |  |  |
| 8 | priviledge_sign_in | 是否有签到特权 | int |  | YES |  |  |
| 9 | priviledge_comment | 是否有评论获奖励特权 | int |  | YES |  |  |
| 10 | priviledge_promotion | 是否有专享活动特权 | int |  | YES |  |  |
| 11 | priviledge_member_price | 是否有会员价格特权 | int |  | YES |  |  |
| 12 | priviledge_birthday | 是否有生日特权 | int |  | YES |  |  |
| 13 | note |  | varchar(200) |  | YES |  |  |
