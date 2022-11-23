package dao

import (
	"starfission_go_api/internal/model"
	"starfission_go_api/internal/model/starfission/hannels_repo"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/timef"
)

type Hannels struct {
	HannelsId    int64      `gorm:"column:hannels_id;primary_key;AUTO_INCREMENT" json:"hannelsId" form:"hannelsId"`         // 渠道ID
	HannelsTitle string     `gorm:"column:hannels_title;default:''" json:"hannelsTitle" form:"hannelsTitle"`                // 渠道名字
	HannelsKey   string     `gorm:"column:hannels_key;default:''" json:"hannelsKey" form:"hannelsKey"`                      // 渠道关键字
	IsDeleted    int32      `gorm:"column:is_deleted;default:0" json:"isDeleted" form:"isDeleted"`                          // 是否删除
	CreatedAt    timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt    timef.Time `gorm:"column:updated_at;time;default:'0000-00-00 00:00:00'" json:"updatedAt" form:"updatedAt"` // 更新时间
	DeletedAt    timef.Time `gorm:"column:deleted_at;time;default:'0000-00-00 00:00:00'" json:"deletedAt" form:"deletedAt"` // 标记删除时间
}

// GetHannelsById 查询渠道详情By渠道ID
func (d *Dao) GetHannelsById(id int64) (*Hannels, error) {
	m, err := hannels_repo.NewQueryBuilder().
		WhereHannelsId(model.Eq, id).
		WhereIsDeleted(model.Eq, 0).
		First()
	if err != nil {
		return nil, err
	}
	return convert.StructAssign(m, &Hannels{}).(*Hannels), nil
}

// GetHannelsByKey 查询渠道详情By渠道Key
func (d *Dao) GetHannelsByKey(key string) (*Hannels, error) {
	m, err := hannels_repo.NewQueryBuilder().
		WhereHannelsKey(model.Eq, key).
		WhereIsDeleted(model.Eq, 0).
		First()
	if err != nil {
		return nil, err
	}
	return convert.StructAssign(m, &Hannels{}).(*Hannels), nil
}
