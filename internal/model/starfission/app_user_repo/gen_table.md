#### starfission.app_user 
应用用户信息

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | app_uid | 应用ID | bigint(20) unsigned | PRI | NO | auto_increment |  |
| 2 | open_id | 用户openID | varchar(40) | UNI | NO |  |  |
| 3 | uid | 对应的用户id | bigint(20) unsigned |  | NO |  | 0 |
| 4 | app_id | appID | bigint(20) |  | NO |  | 0 |
| 5 | is_deleted | 是否删除 | tinyint(1) unsigned |  | NO |  | 0 |
| 6 | created_at | 创建时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 7 | updated_at | 更新时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 8 | deleted_at | 标记删除时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
