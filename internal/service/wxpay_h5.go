package service

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"starfission_go_api/global"
	"starfission_go_api/internal/dao"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/errcode"
	"starfission_go_api/pkg/timef"

	"github.com/gin-gonic/gin"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type WechatH5Pay struct {
	H5Url string `json:"h5Url"` //订单ID
}

type CreateWechatPayH5Request struct {
	OrderId     int64  `form:"orderID" binding:"required,gte=1"` //订单ID
	RedirectUrl string `form:"redirectURL"`                      //跳转链接
}

// 创建订单
func (svc *Service) OrderWechatH5PayCreate(uid int64, param *CreateWechatPayH5Request, requestIp string) (*WechatH5Pay, error) {

	//查询订单
	orderDao, err := svc.dao.GetOrderById(param.OrderId)

	if err != nil {
		return nil, err
	}

	//不需要支付
	if orderDao.OrderType != dao.OrderTypeUserPayOrder {
		return nil, errcode.ErrorPayOrderNotNeedFail
	}

	//订单金额为零 不需要支付
	if orderDao.GoodsTotalCost == 0 {
		return nil, errcode.ErrorPayOrderTotalCostZeroFail
	}

	//非本人订单不能支付
	if orderDao.Uid != uid {
		return nil, errcode.ErrorPayOrderOwnFail
	}

	//已经支付
	if orderDao.Status == 2 {
		return nil, errcode.ErrorPayOrderPaidFail
	}

	//其他状态
	if orderDao.Status != 1 {
		return nil, errcode.ErrorPayOrderNotUnpaidFail
	}

	var (
		appid                      string = global.WechatH5PaySetting.Appid                // 商户号
		mchID                      string = global.WechatH5PaySetting.Mchid                // 商户号
		mchCertificateSerialNumber string = global.WechatH5PaySetting.MchCertificateSerial // 商户证书序列号
		mchAPIv3Key                string = global.WechatH5PaySetting.Apiv3Key             // 商户APIv3密钥
		mchPrivateKeyFilePath      string = global.WechatH5PaySetting.MchPrivateKeyFilePath
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(mchPrivateKeyFilePath)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	svcx := h5.H5ApiService{Client: client}
	resp, _, err := svcx.Prepay(ctx,
		h5.PrepayRequest{
			Appid:       core.String(appid),
			Mchid:       core.String(mchID),
			Description: core.String(orderDao.GoodsName),
			OutTradeNo:  core.String(orderDao.OrderSn),
			TimeExpire:  core.Time(time.Time(orderDao.CreatedAt).Add(+time.Minute * 15)),
			Attach:      core.String(strconv.Itoa(int(orderDao.OrderId))),
			//NotifyUrl:     core.String("https://www.starfission.top/pay/WxPayNotify"),
			NotifyUrl:     core.String("https://goapi.starfission.com/api/v1/pay/wechat_h5_notify"),
			GoodsTag:      core.String(orderDao.GoodsSn),
			SupportFapiao: core.Bool(false),
			Amount: &h5.Amount{
				Currency: core.String("CNY"),
				Total:    core.Int64(int64(orderDao.GoodsTotalCost * 100)),
			},
			SceneInfo: &h5.SceneInfo{
				PayerClientIp: core.String(requestIp),
				H5Info: &h5.H5Info{
					Type: core.String("wap"),
				},
			},
		},
	)

	if err != nil {
		// 处理错误
		return nil, err
	}
	var H5Url string = *resp.H5Url
	if param.RedirectUrl != "" {
		H5Url += "&redirect_url=" + url.QueryEscape(param.RedirectUrl)
	}
	return &WechatH5Pay{H5Url: H5Url}, nil
}

// 创建订单
func (svc *Service) OrderWechatH5PayNotify(c *gin.Context) (interface{}, error) {

	var (
		//appid                      string = global.WechatH5PaySetting.Appid                // 商户号
		mchID                      string = global.WechatH5PaySetting.Mchid                // 商户号
		mchCertificateSerialNumber string = global.WechatH5PaySetting.MchCertificateSerial // 商户证书序列号
		mchAPIv3Key                string = global.WechatH5PaySetting.Apiv3Key             // 商户APIv3密钥
		mchPrivateKeyFilePath      string = global.WechatH5PaySetting.MchPrivateKeyFilePath
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(mchPrivateKeyFilePath)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
	err = downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, mchPrivateKey, mchCertificateSerialNumber, mchID, mchAPIv3Key)
	if err != nil {
		return nil, err
	}

	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(mchID)
	// 3. 使用证书访问器初始化 `notify.Handler`
	handler := notify.NewNotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))

	transaction := new(payments.Transaction)
	_, err = handler.ParseNotifyRequest(context.Background(), c.Request, transaction)
	// 如果验签未通过，或者解密失败
	if err != nil {
		return nil, err
	}

	orderId, _ := convert.StrTo(*transaction.Attach).Int64()

	transactionJson, _ := json.Marshal(transaction)

	successTime, _ := time.Parse(time.RFC3339, *transaction.SuccessTime)

	svc.OrderPayNotify(orderId, &OrderPayParams{
		PaymentType:             "wxpay",
		PaymentCost:             float64(*transaction.Amount.Total),
		PaymentDatetime:         timef.Time(successTime),
		PaymentPayTransactionId: *transaction.TransactionId,
		PaymentData:             string(transactionJson),
	})

	return transaction, nil
}
