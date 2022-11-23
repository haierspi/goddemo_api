package v1

import (
	"starfission_go_api/global"
	"starfission_go_api/internal/service"
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func (t *Order) OrderWalletGoodsMetadataCreate(c *gin.Context) {
	param := &service.OrderWalletGoodsMetadataCreateRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	orderSvc, err := svc.OrderWalletGoodsMetadataCreate(uid, param)

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

func (t *Order) OrderWalletGoodsMetadataList(c *gin.Context) {

	param := &service.OrderWalletGoodsMetadataRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	orderSvc, err := svc.OrderWalletGoodsMetadataLatestDetails(uid, param)

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
