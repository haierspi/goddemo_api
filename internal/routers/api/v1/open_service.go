package v1

import (
	"starfission_go_api/global"
	"starfission_go_api/internal/service"
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type OpenService struct{}

func NewOpenService() OpenService {
	return OpenService{}
}

func (t *OpenService) SyncAppMember(c *gin.Context) {

	param := &service.AppMemberSyncRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	svcRes, err := svc.AppMemberSync(param)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorOpenServiceMemberFail)
		}
		return
	}

	response.ToResponse(svcRes)
	return
}

func (t *OpenService) CreateAppOrder(c *gin.Context) {

	param := &service.AppOrderCreateRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	svcRes, err := svc.AppOrderCreate(param)

	if err != nil {
		global.Logger.Errorf(c, "svc.AddressList err: %v", err)
		if curErr, ok := err.(*errcode.Error); ok {
			response.ToErrorResponse(curErr)
		} else {
			response.ToErrorResponse(errcode.ErrorOpenServiceMemberFail)
		}
		return
	}

	response.ToResponse(svcRes)
	return
}
