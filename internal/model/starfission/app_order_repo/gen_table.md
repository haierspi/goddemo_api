#### starfission.app_order 
应用订单信息

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | app_order_id | APP 订单ID | bigint(20) unsigned | PRI | NO | auto_increment |  |
| 2 | order_id | 订单ID | bigint(20) unsigned | MUL | NO |  | 0 |
| 3 | app_uid | 应用UID | bigint(20) unsigned |  | NO |  | 0 |
| 4 | open_id | 用户openID | varchar(40) |  | NO |  |  |
| 5 | goods_id | 商品ID | bigint(20) unsigned |  | NO |  | 0 |
| 6 | uid | 对应的用户id | bigint(20) unsigned |  | NO |  | 0 |
| 7 | app_id | appID | bigint(20) |  | NO |  | 0 |
| 8 | out_trade_no | 商户订单 | varchar(255) |  | NO |  |  |
| 9 | metadata_image | 图像元数据 | varchar(255) |  | NO |  |  |
| 10 | notify_url | 通知URL | varchar(255) |  | NO |  |  |
| 11 | is_success_notifyed | 是否已经成功通知 | tinyint(1) unsigned |  | NO |  | 0 |
| 12 | is_deleted | 是否删除 | tinyint(1) unsigned |  | NO |  | 0 |
| 13 | created_at | 创建时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 14 | updated_at | 更新时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 15 | deleted_at | 标记删除时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
