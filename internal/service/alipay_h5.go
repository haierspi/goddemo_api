package service

import (
	"encoding/json"
	"strconv"
	"time"

	"starfission_go_api/global"
	"starfission_go_api/internal/dao"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/errcode"
	"starfission_go_api/pkg/timef"

	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
)

type CreateAlipayPayH5Request struct {
	OrderId     int64  `form:"orderID" binding:"required,gte=1"` //订单ID
	RedirectUrl string `form:"redirectURL"`                      //跳转链接
}

type AlipayH5Pay struct {
	H5Url string `json:"h5Url"` //订单ID
}

// 创建订单
func (svc *Service) OrderAlipayH5PayCreate(uid int64, param *CreateAlipayPayH5Request, c *gin.Context) (*AlipayH5Pay, error) {

	var accessHost string

	if host, exist := c.Get("access_host"); exist {
		accessHost = host.(string)
	} else {
		accessHost = global.ServerSetting.DefaultAccessHost
	}

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

	client, err := alipay.NewClient(global.AlipayH5PaySetting.Appid, global.AlipayH5PaySetting.PrivateKey, global.AlipayH5PaySetting.IsProd)
	if err != nil {
		return nil, err
	}
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		//SetReturnUrl("https://www.fmm.ink").
		SetNotifyUrl(accessHost + "/api/v1/pay/alipay_h5_notify")
	if param.RedirectUrl != "" {
		client.SetReturnUrl(param.RedirectUrl)
	}

	//请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", orderDao.GoodsName)
	bm.Set("out_trade_no", orderDao.OrderSn)
	//bm.Set("quit_url", "https://www.fmm.ink")
	bm.Set("total_amount", orderDao.GoodsTotalCost)
	bm.Set("product_code", "QUICK_WAP_WAY")
	bm.Set("time_expire", time.Now().Add(+time.Minute*15).Format("2006-01-02 15:04:05"))
	bm.Set("passback_params", orderDao.OrderId)

	//bm.SetBodyMap("extend_params", func(bm gopay.BodyMap) {
	//	bm.Set("sys_service_provider_id", orderDao.OrderId)
	//})

	//手机网站支付请求
	h5Url, err := client.TradeWapPay(svc.ctx, bm)
	if err != nil {
		//xlog.Error("err:", err)
		return nil, err
	}

	return &AlipayH5Pay{H5Url: h5Url}, nil

}

// 支付宝H5支付确认
func (svc *Service) OrderAlipayH5PayNotify(c *gin.Context) error {

	notifyReq, err := alipay.ParseNotifyToBodyMap(c.Request) // c.Request 是 gin 框架的写法

	if err != nil {
		return err
	}

	ok, err := alipay.VerifySign(global.AlipayH5PaySetting.AliPublicKey, notifyReq)

	if !ok {
		return err
	}

	orderId, _ := convert.StrTo(notifyReq.Get("passback_params")).Int64()

	successTime, _ := time.Parse(timef.TimeFormat, notifyReq.Get("gmt_payment"))
	totalAmount, _ := strconv.ParseFloat(notifyReq.Get("total_amount"), 64)
	notifyReqContent, _ := json.Marshal(notifyReq)

	err = svc.OrderPayNotify(orderId, &OrderPayParams{
		PaymentType:             "alipay",
		PaymentCost:             totalAmount,
		PaymentDatetime:         timef.Time(successTime),
		PaymentPayTransactionId: notifyReq.Get("trade_no"),
		PaymentData:             string(notifyReqContent),
	})

	if err != nil {
		return err
	}

	return nil

}
