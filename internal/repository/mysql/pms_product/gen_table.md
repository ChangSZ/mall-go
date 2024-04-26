#### mall.pms_product 
商品信息

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint | PRI | NO | auto_increment |  |
| 2 | brand_id |  | bigint |  | YES |  |  |
| 3 | product_category_id |  | bigint |  | YES |  |  |
| 4 | feight_template_id |  | bigint |  | YES |  |  |
| 5 | product_attribute_category_id |  | bigint |  | YES |  |  |
| 6 | name |  | varchar(200) |  | NO |  |  |
| 7 | pic |  | varchar(255) |  | YES |  |  |
| 8 | product_sn | 货号 | varchar(64) |  | NO |  |  |
| 9 | delete_status | 删除状态：0->未删除；1->已删除 | int |  | YES |  |  |
| 10 | publish_status | 上架状态：0->下架；1->上架 | int |  | YES |  |  |
| 11 | new_status | 新品状态:0->不是新品；1->新品 | int |  | YES |  |  |
| 12 | recommand_status | 推荐状态；0->不推荐；1->推荐 | int |  | YES |  |  |
| 13 | verify_status | 审核状态：0->未审核；1->审核通过 | int |  | YES |  |  |
| 14 | sort | 排序 | int |  | YES |  |  |
| 15 | sale | 销量 | int |  | YES |  |  |
| 16 | price |  | decimal(10,2) |  | YES |  |  |
| 17 | promotion_price | 促销价格 | decimal(10,2) |  | YES |  |  |
| 18 | gift_growth | 赠送的成长值 | int |  | YES |  | 0 |
| 19 | gift_point | 赠送的积分 | int |  | YES |  | 0 |
| 20 | use_point_limit | 限制使用的积分数 | int |  | YES |  |  |
| 21 | sub_title | 副标题 | varchar(255) |  | YES |  |  |
| 22 | description | 商品描述 | text |  | YES |  |  |
| 23 | original_price | 市场价 | decimal(10,2) |  | YES |  |  |
| 24 | stock | 库存 | int |  | YES |  |  |
| 25 | low_stock | 库存预警值 | int |  | YES |  |  |
| 26 | unit | 单位 | varchar(16) |  | YES |  |  |
| 27 | weight | 商品重量，默认为克 | decimal(10,2) |  | YES |  |  |
| 28 | preview_status | 是否为预告商品：0->不是；1->是 | int |  | YES |  |  |
| 29 | service_ids | 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮 | varchar(64) |  | YES |  |  |
| 30 | keywords |  | varchar(255) |  | YES |  |  |
| 31 | note |  | varchar(255) |  | YES |  |  |
| 32 | album_pics | 画册图片，连产品图片限制为5张，以逗号分割 | varchar(255) |  | YES |  |  |
| 33 | detail_title |  | varchar(255) |  | YES |  |  |
| 34 | detail_desc |  | text |  | YES |  |  |
| 35 | detail_html | 产品详情网页内容 | text |  | YES |  |  |
| 36 | detail_mobile_html | 移动端网页详情 | text |  | YES |  |  |
| 37 | promotion_start_time | 促销开始时间 | datetime |  | YES |  |  |
| 38 | promotion_end_time | 促销结束时间 | datetime |  | YES |  |  |
| 39 | promotion_per_limit | 活动限购数量 | int |  | YES |  |  |
| 40 | promotion_type | 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购 | int |  | YES |  |  |
| 41 | brand_name | 品牌名称 | varchar(255) |  | YES |  |  |
| 42 | product_category_name | 商品分类名称 | varchar(255) |  | YES |  |  |
