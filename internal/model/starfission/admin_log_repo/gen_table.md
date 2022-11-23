#### starfission.admin_log 
后台访问日志

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | ID | bigint(20) unsigned | PRI | NO |  | 0 |
| 2 | auid | 用户AUid | bigint(20) unsigned | MUL | NO |  | 0 |
| 3 | username |  | char(24) |  | NO |  | '' |
| 4 | dateline |  | int(10) unsigned | MUL | NO |  | 0 |
| 5 | action |  | char(15) |  | NO |  | '' |
| 6 | url |  | text |  | NO |  |  |
| 7 | method |  | char(15) |  | NO |  | '' |
| 8 | request |  | text |  | NO |  |  |
| 9 | updated_at | 更新时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 10 | created_at | 创建时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 11 | deleted_at | 标记删除时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
