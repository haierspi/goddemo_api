#### starfission.admin_perm 
后台权限表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | auid | 后台用户UID | bigint(20) unsigned | PRI | NO |  | 0 |
| 2 | modkeyword | 模块权限关键字 | char(255) | PRI | NO |  | '' |
| 3 | updated_at | 更新时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 4 | created_at | 创建时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 5 | deleted_at | 标记删除时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
