#### starfission.app 
开放应用

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | app_id | 应用ID | bigint(20) unsigned | PRI | NO | auto_increment |  |
| 2 | app_name | 应用名称 | varchar(255) |  | NO |  |  |
| 3 | app_url | 专题商品页访问地址 | varchar(255) |  | NO |  |  |
| 4 | app_secret | 用于加密的密钥 | varchar(255) |  | NO |  |  |
| 5 | goods_id | 上链商品ID | bigint(20) |  | NO |  | 0 |
| 6 | notify_domain | 异步通知安全域名 | varchar(255) |  | NO |  |  |
| 7 | is_published | 是否已上线(已发布) | tinyint(1) unsigned |  | NO |  | 0 |
| 8 | is_deleted | 是否删除 | tinyint(1) unsigned |  | NO |  | 0 |
| 9 | created_at | 创建时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 10 | updated_at | 更新时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 11 | deleted_at | 标记删除时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
