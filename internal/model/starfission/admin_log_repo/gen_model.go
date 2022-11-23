package admin_log_repo

import "starfission_go_api/pkg/timef"

// 后台访问日志
//
//go:generate gormgen -structs AdminLog -input . -pre pre_
type AdminLog struct {
	Id        int64      `gorm:"column:id;primary_key;default:0" json:"id" form:"id"`                                    // ID
	Auid      int64      `gorm:"column:auid;default:0" json:"auid" form:"auid"`                                          // 用户AUid
	Username  string     `gorm:"column:username;default:''" json:"username" form:"username"`                             //
	Dateline  int32      `gorm:"column:dateline;default:0" json:"dateline" form:"dateline"`                              //
	Action    string     `gorm:"column:action;default:''" json:"action" form:"action"`                                   //
	Url       string     `gorm:"column:url" json:"url" form:"url"`                                                       //
	Method    string     `gorm:"column:method;default:''" json:"method" form:"method"`                                   //
	Request   string     `gorm:"column:request" json:"request" form:"request"`                                           //
	UpdatedAt timef.Time `gorm:"column:updated_at;time;default:'0000-00-00 00:00:00'" json:"updatedAt" form:"updatedAt"` // 更新时间
	CreatedAt timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"` // 创建时间
	DeletedAt timef.Time `gorm:"column:deleted_at;time;default:'0000-00-00 00:00:00'" json:"deletedAt" form:"deletedAt"` // 标记删除时间
}
