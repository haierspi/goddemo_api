package v1

import (
	"starfission_go_api/global"
	"starfission_go_api/internal/service"
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Wallet struct{}

func NewWallet() Wallet {
	return Wallet{}
}

func (t *Wallet) List(c *gin.Context) {

	response := app.NewResponse(c)

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	transferSSvc, transferCount, err := svc.WalletList(uid)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorPayCreateFail)
		}
		return
	}

	response.ToResponseList(transferSSvc, transferCount)
	return
}

func (t *Wallet) Details(c *gin.Context) {

	param := &service.WalletDetailsRequest{Id: convert.StrTo(c.Param("id")).MustInt64()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	orderSvc, err := svc.WalletDetails(uid, param)

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

func (t *Wallet) Given(c *gin.Context) {

	param := &service.WalletGivenRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	id, err := svc.WalletGiven(uid, param)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorOrderViewFail)
		}
		return
	}

	response.ToResponse(id)
	return
}
