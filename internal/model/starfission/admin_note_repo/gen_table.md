#### starfission.admin_note 
后台用户便签

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | ID | bigint(20) unsigned | PRI | NO |  | 0 |
| 2 | auid | 用户AUid | bigint(20) unsigned | MUL | NO |  | 0 |
| 3 | notetext | 便签本文 | text |  | YES |  | NULL |
| 4 | updated_at | 更新时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 5 | created_at | 创建时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 6 | deleted_at | 标记删除时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
