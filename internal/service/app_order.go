package service

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strconv"

	"starfission_go_api/internal/dao"
	"starfission_go_api/pkg/errcode"
	"starfission_go_api/pkg/oss"
	"starfission_go_api/pkg/util"

	mapset "github.com/deckarep/golang-set"
	"github.com/segmentio/ksuid"
)

type AppOrderCreateRequest struct {
	AppRequest
	OpenId        string `form:"openid" binding:"required"`         //openid
	OutTradeNo    string `form:"out_trade_no" binding:"required"`   //商户订单号
	MetadataImage string `form:"metadata_image" binding:"required"` //元数据图片
	NotifyUrl     string `form:"notify_url" binding:"required"`     //通知地址
	GoodsId       int64  `form:"goods_id"`                          //商品ID
}

type AppOrderGoodsDetailsRequest struct {
	AppOrderId int64 `form:"appOrderID" binding:"required,gte=1"`
}

type AppOrderCreateResponse struct {
	RedirectUrl string `json:"redirect_url"`
}

var appGoodsPay = `https://m.starfission.com%s?appOrderID=%d`
var metadataHost = `https://metadata.starfission.com/`

// AppOrderCreate 商户订单创建
func (svc *Service) AppOrderCreate(param *AppOrderCreateRequest) (*AppOrderCreateResponse, error) {

	app, err := svc.dao.GetAppByID(param.AppID)
	if err != nil {
		return nil, err
	}

	nUrl, err := url.Parse(param.NotifyUrl)
	if err != nil {
		return nil, errcode.ErrorOpenServiceNotifyUrlParseFail
	}

	if nUrl.Scheme != "http" && nUrl.Scheme != "https" {
		return nil, errcode.ErrorOpenServiceNotifySchemeVerifyFail
	}

	if nUrl.Host != app.NotifyDomain {
		return nil, errcode.ErrorOpenServiceNotifyDomainVerifyFail
	}

	mUrl, err := url.Parse(param.MetadataImage)
	if err != nil {
		return nil, errcode.ErrorOpenServiceImageUrlParseFail
	}

	if mUrl.Scheme != "http" && mUrl.Scheme != "https" {
		return nil, errcode.ErrorOpenServiceNotifySchemeVerifyFail
	}

	var g *dao.Goods
	if app.GoodsId == 0 {
		g, err = svc.dao.GetGoodsById(param.GoodsId)
		if err != nil {
			return nil, err
		}
	} else {
		g, err = svc.dao.GetGoodsById(app.GoodsId)
		if err != nil {
			return nil, err
		}
	}

	u, err := svc.dao.GetAppUserByOpenId(param.OpenId)

	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(mUrl.Path)

	extMap := mapset.NewSetFromSlice([]interface{}{".gif", ".jpg", ".jpeg", ".png", ".webp"})

	if !extMap.Contains(ext) {
		return nil, errcode.ErrorOpenServiceOrderImageTypeFail
	}

	id := ksuid.New()

	name := util.EncodeMD5(param.NotifyUrl) + "_" + id.String()

	path := "app/appid-" + strconv.Itoa(int(app.AppId)) + "/" + name + ext

	err = oss.UploadByURL(path, param.MetadataImage)

	if err != nil {
		return nil, errcode.ErrorOpenServiceOrderImageTypeFail

	}
	metadataImage := metadataHost + path

	appOrderId, err := svc.dao.CreateAppOrder(&dao.AppOrder{
		AppUid:        u.AppUid,
		OpenId:        u.OpenId,
		Uid:           u.Uid,
		OutTradeNo:    param.OutTradeNo,
		MetadataImage: metadataImage,
		NotifyUrl:     param.NotifyUrl,
		AppId:         app.AppId,
		GoodsId:       g.GoodsId,
	})

	if err != nil {
		return nil, err
	}

	redirectUrl := fmt.Sprintf(appGoodsPay, app.AppUrl, appOrderId)

	return &AppOrderCreateResponse{RedirectUrl: redirectUrl}, nil
}

// AppOrderGoodsDetails 根据商户订单ID 获取商品详情
func (svc *Service) AppOrderGoodsDetails(param *AppOrderGoodsDetailsRequest) (*Goods, error) {
	a, err := svc.dao.GetAppOrderById(param.AppOrderId)
	if err != nil {
		return nil, err
	}
	if a.OrderId != 0 {
		o, err := svc.dao.GetOrderById(a.OrderId)
		if err != nil {
			return nil, err
		}
		if o.Status != dao.OrderStatusClosed {
			return nil, errcode.ErrorOpenServiceOrderRepeatCreateFail
		}
	}

	return svc.GoodsDetails(&GoodsDetailsRequest{GoodsID: a.GoodsId})
}
