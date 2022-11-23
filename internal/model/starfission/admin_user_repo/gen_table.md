#### starfission.admin_user 
后台管理员表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | auid |  | bigint(20) unsigned | PRI | NO | auto_increment |  |
| 2 | username |  | char(24) |  | NO |  | '' |
| 3 | password | 密码 | char(32) |  | NO |  | '' |
| 4 | salt | 密码混淆码 | char(24) |  | NO |  | '' |
| 5 | dateline |  | int(10) unsigned |  | NO |  | 0 |
| 6 | permission | 权限列表 | longtext |  | NO |  |  |
| 7 | token | 用户授权令牌 | char(255) |  | NO |  | '' |
| 8 | updated_at | 更新时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 9 | created_at | 创建时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 10 | deleted_at | 标记删除时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
