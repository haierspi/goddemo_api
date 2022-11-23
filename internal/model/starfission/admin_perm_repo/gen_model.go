package admin_perm_repo

import "starfission_go_api/pkg/timef"

// 后台权限表
//
//go:generate gormgen -structs AdminPerm -input . -pre pre_
type AdminPerm struct {
	Auid       int64      `gorm:"column:auid;primary_key;default:0" json:"auid" form:"auid"`                              // 后台用户UID
	Modkeyword string     `gorm:"column:modkeyword;primary_key;default:''" json:"modkeyword" form:"modkeyword"`           // 模块权限关键字
	UpdatedAt  timef.Time `gorm:"column:updated_at;time;default:'0000-00-00 00:00:00'" json:"updatedAt" form:"updatedAt"` // 更新时间
	CreatedAt  timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"` // 创建时间
	DeletedAt  timef.Time `gorm:"column:deleted_at;time;default:'0000-00-00 00:00:00'" json:"deletedAt" form:"deletedAt"` // 标记删除时间
}
