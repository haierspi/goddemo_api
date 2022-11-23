package v1

import (
	"starfission_go_api/global"
	"starfission_go_api/internal/service"
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Order struct{}

func NewOrder() Order {
	return Order{}
}

func (t *Order) Create(c *gin.Context) {
	param := service.CreateOrderRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	orderSvc, err := svc.CreateOrder(uid, &param)
	if err != nil {
		global.Logger.Errorf(c, "svc.OrderCreate err: %v", err)

		if orderErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(orderErr)
		} else {
			response.ToErrorResponse(errcode.ErrorOrderCreateFail)
		}

		return
	}

	response.ToResponse(orderSvc)
	return
}

func (t *Order) Give(c *gin.Context) {

	param := &service.OrderDetailsRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	orderSvc, err := svc.OrderGive(uid, param)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorOrderGiveFail)
		}
		return
	}

	response.ToResponse(orderSvc)
	return
}

func (t *Order) OrderFreeReceiveCreate(c *gin.Context) {
	param := service.CreateFreeReceiveOrderRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	orderSvc, err := svc.CreateFreeReceiveOrder(uid, &param)
	if err != nil {
		global.Logger.Errorf(c, "svc.OrderCreate err: %v", err)

		if orderErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(orderErr)
		} else {
			response.ToErrorResponse(errcode.ErrorOrderCreateFail)
		}

		return
	}

	response.ToResponse(orderSvc)
	return
}

func (t *Order) OrderRedeemCodeCreate(c *gin.Context) {
	param := service.CreateRedeemCodeOrderRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	orderSvc, err := svc.CreateRedeemCodeOrder(uid, &param)
	if err != nil {
		global.Logger.Errorf(c, "svc.OrderCreate err: %v", err)

		if orderErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(orderErr)
		} else {
			response.ToErrorResponse(errcode.ErrorOrderCreateFail)
		}

		return
	}

	response.ToResponse(orderSvc)
	return
}

func (t *Order) List(c *gin.Context) {

	param := &service.ListOrderRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	app.RequestParamStrParse(c, param)

	pager := &app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	orderSSvc, orderCount, err := svc.OrderList(uid, param, pager)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorPayCreateFail)
		}
		return
	}

	response.ToResponseList(orderSSvc, orderCount)
	return
}

func (t *Order) URLDetails(c *gin.Context) {

	param := &service.OrderDetailsRequest{OrderId: convert.StrTo(c.Param("id")).MustInt64()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	orderSvc, err := svc.OrderDetails(uid, param)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorOrderViewFail)
		}
		return
	}

	response.ToResponse(orderSvc)
	return
}

func (t *Order) Details(c *gin.Context) {

	param := &service.OrderDetailsRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	orderSvc, err := svc.OrderDetails(uid, param)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorOrderViewFail)
		}
		return
	}

	response.ToResponse(orderSvc)
	return
}

func (t *Order) OrderChargeOff(c *gin.Context) {
	param := &service.OrderChargeOffRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.OrderChargeOff(param)
	if err != nil {
		global.Logger.Errorf(c, "svc.OrderCreate err: %v", err)

		if orderErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(orderErr)
		} else {
			response.ToErrorResponse(errcode.ErrorOrderChargeOffFail)
		}

		return
	}

	response.ToErrorResponse(errcode.SuccessOrderChargeOff)
	return
}
