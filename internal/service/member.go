package service

import (
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/timef"
)

type Member struct {
	Uid              int64      `gorm:"column:uid;primary_key;AUTO_INCREMENT" json:"uid" form:"uid"`                            // 用户UID
	Nickname         string     `gorm:"column:nickname;default:''" json:"nickname" form:"nickname"`                             // 用户昵称
	Avatar           string     `gorm:"column:avatar;default:''" json:"avatar" form:"avatar"`                                   // 头像
	Mobile           string     `gorm:"column:mobile;default:''" json:"mobile" form:"mobile"`                                   // 电话号码
	Name             string     `gorm:"column:name;default:''" json:"name" form:"name"`                                         // 真实姓名
	Idcard           string     `gorm:"column:idcard;default:''" json:"idcard" form:"idcard"`                                   // 身份证号
	IsValidate       int32      `gorm:"column:is_validate;default:0" json:"isValidate" form:"isValidate"`                       // 是否身份验证
	IsWalletPassword int32      `json:"IsWalletPassword"`                                                                       //是否设置过二级密码
	Openid           string     `gorm:"column:openid;default:''" json:"openid" form:"openid"`                                   // openid
	Unionid          string     `gorm:"column:unionid;default:''" json:"unionid" form:"unionid"`                                // unionid
	Gender           string     `gorm:"column:gender;default:''" json:"gender" form:"gender"`                                   // 性别
	Language         string     `gorm:"column:language;default:''" json:"language" form:"language"`                             // 语言
	City             string     `gorm:"column:city;default:''" json:"city" form:"city"`                                         // 城市
	Province         string     `gorm:"column:province;default:''" json:"province" form:"province"`                             // 省份
	Country          string     `gorm:"column:country;default:''" json:"country" form:"country"`                                // 国家
	AvatarUrl        string     `gorm:"column:avatar_url;default:''" json:"avatarUrl" form:"avatarUrl"`                         // 微信的头像
	ChangeNameNum    int32      `gorm:"column:change_name_num;default:0" json:"changeNameNum" form:"changeNameNum"`             // 用户昵称剩余修改次数
	CreatedAt        timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"` // 创建时间
}

func (svc *Service) MemberProfile(uid int64) (*Member, error) {

	memberDao, err := svc.dao.GetMemberByUID(uid)

	if err != nil {
		return nil, err
	}

	walletSecurity, _ := svc.dao.GetWalletSecurityByUid(uid)

	memberSVC := convert.StructAssign(memberDao, &Member{}).(*Member)

	if walletSecurity != nil {
		memberSVC.IsWalletPassword = 1
	}

	return memberSVC, nil
}
