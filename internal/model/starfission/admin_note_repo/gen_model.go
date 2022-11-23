package admin_note_repo

import "starfission_go_api/pkg/timef"

// 后台用户便签
//
//go:generate gormgen -structs AdminNote -input . -pre pre_
type AdminNote struct {
	Id        int64      `gorm:"column:id;primary_key;default:0" json:"id" form:"id"`                                    // ID
	Auid      int64      `gorm:"column:auid;default:0" json:"auid" form:"auid"`                                          // 用户AUid
	Notetext  string     `gorm:"column:notetext;default:NULL" json:"notetext" form:"notetext"`                           // 便签本文
	UpdatedAt timef.Time `gorm:"column:updated_at;time;default:'0000-00-00 00:00:00'" json:"updatedAt" form:"updatedAt"` // 更新时间
	CreatedAt timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"` // 创建时间
	DeletedAt timef.Time `gorm:"column:deleted_at;time;default:'0000-00-00 00:00:00'" json:"deletedAt" form:"deletedAt"` // 标记删除时间
}
