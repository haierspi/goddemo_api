package app_repo

import "starfission_go_api/pkg/timef"

// 开放应用
//
//go:generate gormgen -structs App -input . -pre pre_
type App struct {
	AppId        int64      `gorm:"column:app_id;primary_key;AUTO_INCREMENT" json:"appId" form:"appId"`                     // 应用ID
	AppName      string     `gorm:"column:app_name" json:"appName" form:"appName"`                                          // 应用名称
	AppUrl       string     `gorm:"column:app_url" json:"appUrl" form:"appUrl"`                                             // 专题商品页访问地址
	AppSecret    string     `gorm:"column:app_secret" json:"appSecret" form:"appSecret"`                                    // 用于加密的密钥
	GoodsId      int64      `gorm:"column:goods_id;default:0" json:"goodsId" form:"goodsId"`                                // 上链商品ID
	NotifyDomain string     `gorm:"column:notify_domain" json:"notifyDomain" form:"notifyDomain"`                           // 异步通知安全域名
	IsPublished  int32      `gorm:"column:is_published;default:0" json:"isPublished" form:"isPublished"`                    // 是否已上线(已发布)
	IsDeleted    int32      `gorm:"column:is_deleted;default:0" json:"isDeleted" form:"isDeleted"`                          // 是否删除
	CreatedAt    timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt    timef.Time `gorm:"column:updated_at;time;default:'0000-00-00 00:00:00'" json:"updatedAt" form:"updatedAt"` // 更新时间
	DeletedAt    timef.Time `gorm:"column:deleted_at;time;default:'0000-00-00 00:00:00'" json:"deletedAt" form:"deletedAt"` // 标记删除时间
}
