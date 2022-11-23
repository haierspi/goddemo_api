package v1

import (
	"net/http"

	"starfission_go_api/global"
	"starfission_go_api/internal/service"
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Pay struct{}

func NewPay() Pay {
	return Pay{}
}

// @Summary 微信H5支付
func (t *Pay) OrderWechatH5PayCreate(c *gin.Context) {
	param := service.CreateWechatPayH5Request{}
	requestIp := app.GetRequestIP(c)
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	orderSvc, err := svc.OrderWechatH5PayCreate(uid, &param, requestIp)
	if err != nil {
		global.Logger.Errorf(c, "svc.WechatPayCreate err: %v", err)

		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorPayCreateFail)
		}

		return
	}

	response.ToResponse(orderSvc)
	return
}

// @Summary 微信JSSDK支付
func (t *Pay) OrderWechatJSAPIPayCreate(c *gin.Context) {
	param := service.CreateWechatPayJSAPIRequest{}
	requestIp := app.GetRequestIP(c)
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	orderSvc, err := svc.OrderWechatJSAPIPayCreate(uid, &param, requestIp)
	if err != nil {
		global.Logger.Errorf(c, "svc.WechatPayCreate err: %v", err)

		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {

			response.ToErrorResponseData(errcode.ErrorPayCreateFail, err)

		}

		return
	}

	response.ToResponse(orderSvc)
	return
}

// @Summary 微信H5支付通知
func (t *Pay) OrderWechatH5PayNotify(c *gin.Context) {
	response := app.NewResponse(c)

	svc := service.New(c.Request.Context())
	orderSvc, err := svc.OrderWechatH5PayNotify(c)
	if err != nil {
		global.Logger.Errorf(c, "svc.WechatPayCreate err: %v", err)

		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {

			response.ToErrorResponseData(errcode.ErrorPayCreateFail, err)

		}

		return
	}

	response.ToResponse(orderSvc)
	return

}

// @Summary 微信JSSDK支付通知
func (t *Pay) OrderWechatJSSDKPayNotify(c *gin.Context) {
	response := app.NewResponse(c)

	svc := service.New(c.Request.Context())
	orderSvc, err := svc.OrderWechatJSAPIPayNotify(c)
	if err != nil {
		global.Logger.Errorf(c, "svc.WechatPayCreate err: %v", err)

		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {

			response.ToErrorResponseData(errcode.ErrorPayCreateFail, err)

		}

		return
	}

	response.ToResponse(orderSvc)
	return

}

// @Summary 支付宝H5支付
func (t *Pay) OrderAlipayH5PayCreate(c *gin.Context) {

	param := service.CreateAlipayPayH5Request{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	orderSvc, err := svc.OrderAlipayH5PayCreate(uid, &param, c)
	if err != nil {
		global.Logger.Errorf(c, "svc.WechatPayCreate err: %v", err)

		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorPayCreateFail)
		}

		return
	}

	response.ToResponse(orderSvc)
	return
}

// @Summary 微信H5支付通知
func (t *Pay) OrderAlipayH5PayNotify(c *gin.Context) {

	svc := service.New(c.Request.Context())
	err := svc.OrderAlipayH5PayNotify(c)
	if err != nil {
		global.Logger.Errorf(c, "svc.OrderAlipayH5PayNotify err: %v", err)
		c.String(http.StatusOK, "%s", "fail")
		return
	}
	c.String(http.StatusOK, "%s", "success")
	return

}
