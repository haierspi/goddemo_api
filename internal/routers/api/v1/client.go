package v1

import (
	"errors"

	"starfission_go_api/global"
	"starfission_go_api/internal/service"
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/errcode"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Client struct{}

func NewClient() Client {
	return Client{}
}

// @Summary 获取客户端版本信息
// @Produce  json
// @Param platform query string false "标签名称" maxlength(100)
// @Param versionCode query int false "状态" Enums(0, 1) default(1)
// @Success 200 {object} app.ResListResult "成功"
// @Failure 200 {object} app.ErrResult  "请求错误"
// @response 200 {object} service.ClientVersion  "请求错误"
// @Router /api/v1/client/version [get]
func (t *Client) Version(c *gin.Context) {

	param := &service.ClientVersionRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	clientVersion, err := svc.ClientVersion(param)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.ToResponse(nil)
	} else {
		response.ToResponse(clientVersion)
	}

	return
}
