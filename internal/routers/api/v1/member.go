package v1

import (
	"starfission_go_api/global"
	"starfission_go_api/internal/service"
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/errcode"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Member struct{}

func NewMember() Member {
	return Member{}
}

func (t *Member) Profile(c *gin.Context) {

	response := app.NewResponse(c)
	uid := app.GetUid(c)

	svc := service.New(c.Request.Context())
	memberSvc, err := svc.MemberProfile(uid)

	if err != nil {
		global.Logger.Errorf(c, "svc.MemberProfile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserNotExist)
		return
	}

	response.ToResponse(memberSvc)
	return
}

func (t *Member) Wallet(c *gin.Context) {

	response := app.NewResponse(c)
	uid := app.GetUid(c)

	svc := service.New(c.Request.Context())
	memberWalletSvc, err := svc.Wallet(uid)

	if err != nil {
		global.Logger.Errorf(c, "svc.Wallet err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserNotExist)
		return
	}

	response.ToResponse(memberWalletSvc)
	return
}

func (t *Member) AddressList(c *gin.Context) {

	response := app.NewResponse(c)
	uid := app.GetUid(c)

	svc := service.New(c.Request.Context())
	addressSSvc, addressCount, err := svc.AddressList(uid)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetAddressFail)
		return
	}

	response.ToResponseList(addressSSvc, addressCount)
	return
}

func (t *Member) AddressDefault(c *gin.Context) {
	response := app.NewResponse(c)
	uid := app.GetUid(c)

	svc := service.New(c.Request.Context())
	addressSvc, err := svc.AddressDefault(uid)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ToErrorResponse(errcode.ErrorGetDefaultAddressFail)
		} else {
			response.ToErrorResponse(errcode.ErrorGetAddressFail)
		}

		return
	}

	response.ToResponse(addressSvc)
	return
}

func (t *Member) AddressCreate(c *gin.Context) {
	param := service.CreateAddressRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	addressSvc, err := svc.AddressCreate(uid, &param)
	if err != nil {
		global.Logger.Errorf(c, "svc.AddressCreate err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateAddressFail)
		return
	}

	response.ToResponse(addressSvc)
	return
}

func (t *Member) AddressUpdate(c *gin.Context) {
	param := service.UpdateAddressRequest{Id: convert.StrTo(c.Param("id")).MustInt64()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	err := svc.AddressUpdate(uid, &param)
	if err != nil {
		global.Logger.Errorf(c, "svc.AddressUpdate err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateAddressFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

func (t *Member) AddressDelete(c *gin.Context) {
	param := service.DeleteAddressRequest{Id: convert.StrTo(c.Param("id")).MustInt64()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	uid := app.GetUid(c)
	svc := service.New(c.Request.Context())
	err := svc.AddressDelete(uid, &param)
	if err != nil {
		global.Logger.Errorf(c, "svc.AddressDelete err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteAddressFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
