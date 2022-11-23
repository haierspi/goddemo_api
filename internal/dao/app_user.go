package dao

import (
	"starfission_go_api/internal/model"
	"starfission_go_api/internal/model/starfission/app_user_repo"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/timef"
)

type AppUser struct {
	AppUid    int64      `gorm:"column:app_uid;primary_key;AUTO_INCREMENT" json:"appUid" form:"appUid"`                  // 应用ID
	OpenId    string     `gorm:"column:openid" json:"openid" form:"openid"`                                              // 用户openID
	Uid       int64      `gorm:"column:uid;default:0" json:"uid" form:"uid"`                                             // 对应的用户id
	AppId     int64      `gorm:"column:app_id;default:0" json:"appId" form:"appId"`                                      // appID
	IsDeleted int32      `gorm:"column:is_deleted;default:0" json:"isDeleted" form:"isDeleted"`                          // 是否删除
	CreatedAt timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt timef.Time `gorm:"column:updated_at;time;default:'0000-00-00 00:00:00'" json:"updatedAt" form:"updatedAt"` // 更新时间
	DeletedAt timef.Time `gorm:"column:deleted_at;time;default:'0000-00-00 00:00:00'" json:"deletedAt" form:"deletedAt"` // 标记删除时间
}

func (d *Dao) GetAppUserByAppIdAndUid(appId int64, uid int64) (*AppUser, error) {

	m, err := app_user_repo.NewQueryBuilder().
		WhereAppId(model.Eq, appId).
		WhereUid(model.Eq, uid).
		WhereIsDeleted(model.Eq, 0).
		First()

	if err != nil {
		return nil, err
	}
	return convert.StructAssign(m, &AppUser{}).(*AppUser), nil
}

func (d *Dao) GetAppUserByOpenId(openId string) (*AppUser, error) {

	m, err := app_user_repo.NewQueryBuilder().
		WhereOpenId(model.Eq, openId).
		WhereIsDeleted(model.Eq, 0).
		First()

	if err != nil {
		return nil, err
	}
	return convert.StructAssign(m, &AppUser{}).(*AppUser), nil
}

func (d *Dao) GetUserByOpenId(openId string) (*Member, error) {

	m, err := app_user_repo.NewQueryBuilder().
		WhereOpenId(model.Eq, openId).
		WhereIsDeleted(model.Eq, 0).
		First()

	if err != nil {
		return nil, err
	}
	return d.GetMemberByUID(m.Uid)
}

func (d *Dao) CreateAppUser(dao *AppUser) (int64, error) {

	m := convert.StructAssign(dao, &app_user_repo.AppUser{}).(*app_user_repo.AppUser)

	id, err := m.Create()

	if err != nil {
		return 0, err
	}
	return id, nil
}
