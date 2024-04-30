#### mall.cms_member_report 
用户举报表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint |  | YES |  |  |
| 2 | report_type | 举报类型：0->商品评价；1->话题内容；2->用户评论 | int |  | YES |  |  |
| 3 | report_member_name | 举报人 | varchar(100) |  | YES |  |  |
| 4 | create_time |  | datetime |  | YES |  |  |
| 5 | report_object |  | varchar(100) |  | YES |  |  |
| 6 | report_status | 举报状态：0->未处理；1->已处理 | int |  | YES |  |  |
| 7 | handle_status | 处理结果：0->无效；1->有效；2->恶意 | int |  | YES |  |  |
| 8 | note |  | varchar(200) |  | YES |  |  |
