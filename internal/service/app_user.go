package service

import (
	"errors"

	"starfission_go_api/internal/dao"
	"starfission_go_api/pkg/rand"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type AppMemberSyncRequest struct {
	AppRequest
	Name   string `form:"name" binding:"required"`   //姓名
	Mobile string `form:"mobile" binding:"required"` //手机号
	Idcard string `form:"idcard" binding:"required"` //身份证
}

type AppMemberSyncResponse struct {
	OpenId string `json:"openid"`
}

// AppMemberSync 会员数据同步
func (svc *Service) AppMemberSync(param *AppMemberSyncRequest) (*AppMemberSyncResponse, error) {

	var uid int64

	app, err := svc.dao.GetAppByID(param.AppID)
	if err != nil {
		return nil, err
	}

	d, err := svc.dao.GetMemberByMobile(param.Mobile)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uid, err = svc.dao.CreateMember(&dao.Member{
				Nickname:   "USER_" + rand.GetRandString(10),
				Avatar:     "https://assets.starfission.com/images/avatars/1.png",
				Name:       param.Name,
				Mobile:     param.Mobile,
				Idcard:     param.Idcard,
				IsValidate: 1,
				AppId:      app.AppId,
			})

			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		uid = d.Uid
	}

	var openId string

	au, err := svc.dao.GetAppUserByAppIdAndUid(app.AppId, uid)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			id := ksuid.New()
			openId = id.String()
			_, err = svc.dao.CreateAppUser(&dao.AppUser{
				OpenId: openId,
				Uid:    uid,
				AppId:  app.AppId,
			})

			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		openId = au.OpenId
	}

	_, err = svc.Wallet(d.Uid)

	if err != nil {
		return nil, err
	}

	return &AppMemberSyncResponse{OpenId: openId}, nil
}
