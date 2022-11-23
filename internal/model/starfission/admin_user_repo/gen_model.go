package admin_user_repo

import "starfission_go_api/pkg/timef"

// 后台管理员表
//
//go:generate gormgen -structs AdminUser -input . -pre pre_
type AdminUser struct {
	Auid       int64      `gorm:"column:auid;primary_key;AUTO_INCREMENT" json:"auid" form:"auid"`                         //
	Username   string     `gorm:"column:username;default:''" json:"username" form:"username"`                             //
	Password   string     `gorm:"column:password;default:''" json:"password" form:"password"`                             // 密码
	Salt       string     `gorm:"column:salt;default:''" json:"salt" form:"salt"`                                         // 密码混淆码
	Dateline   int32      `gorm:"column:dateline;default:0" json:"dateline" form:"dateline"`                              //
	Permission string     `gorm:"column:permission" json:"permission" form:"permission"`                                  // 权限列表
	Token      string     `gorm:"column:token;default:''" json:"token" form:"token"`                                      // 用户授权令牌
	UpdatedAt  timef.Time `gorm:"column:updated_at;time;default:'0000-00-00 00:00:00'" json:"updatedAt" form:"updatedAt"` // 更新时间
	CreatedAt  timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"` // 创建时间
	DeletedAt  timef.Time `gorm:"column:deleted_at;time;default:'0000-00-00 00:00:00'" json:"deletedAt" form:"deletedAt"` // 标记删除时间
}
