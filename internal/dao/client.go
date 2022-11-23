package dao

import (
	"starfission_go_api/internal/model"
	"starfission_go_api/internal/model/starfission/client_version_repo"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/timef"
)

type ClientVersion struct {
	Id          int64      `json:"ID"`          // 用户id
	Platform    int32      `json:"platform"`    // 平台 1:安卓 2:ios
	VersionCode int64      `json:"versionCode"` // 版本号
	VersionName string     `json:"versionName"` // 版本名
	Details     string     `json:"details"`     // 版本更新详情
	ResourceUrl string     `json:"resourceUrl"` // 版本更新资源地址
	CreatedAt   timef.Time `json:"createdAt"`   //创建时间
	UpdatedAt   timef.Time `json:"updatedAt"`   //更新时间
	DeletedAt   timef.Time `json:"deletedAt"`   // 标记删除时间
}

// 查询更新
func (d *Dao) GetClientVersionByCodeAndPlatform(versionCode int64, platform int32) (*ClientVersion, error) {

	m, err := client_version_repo.NewQueryBuilder().
		WherePlatform(model.Eq, platform).
		WhereVersionCode(model.Gt, versionCode).
		OrderById(false).
		First()
	if err != nil {
		return nil, err
	}

	return convert.StructAssign(m, &ClientVersion{}).(*ClientVersion), nil
}
