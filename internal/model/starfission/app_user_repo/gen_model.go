package app_user_repo

import "starfission_go_api/pkg/timef"

// 应用用户信息
//
//go:generate gormgen -structs AppUser -input . -pre pre_
type AppUser struct {
	AppUid    int64      `gorm:"column:app_uid;primary_key;AUTO_INCREMENT" json:"appUid" form:"appUid"`                  // 应用ID
	OpenId    string     `gorm:"column:open_id" json:"openId" form:"openId"`                                             // 用户openID
	Uid       int64      `gorm:"column:uid;default:0" json:"uid" form:"uid"`                                             // 对应的用户id
	AppId     int64      `gorm:"column:app_id;default:0" json:"appId" form:"appId"`                                      // appID
	IsDeleted int32      `gorm:"column:is_deleted;default:0" json:"isDeleted" form:"isDeleted"`                          // 是否删除
	CreatedAt timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt timef.Time `gorm:"column:updated_at;time;default:'0000-00-00 00:00:00'" json:"updatedAt" form:"updatedAt"` // 更新时间
	DeletedAt timef.Time `gorm:"column:deleted_at;time;default:'0000-00-00 00:00:00'" json:"deletedAt" form:"deletedAt"` // 标记删除时间
}
