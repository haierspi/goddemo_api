/**
  @author: HaierSpi
  @since: 2022/9/14
  @desc:
**/

package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"time"

	"starfission_go_api/global"
	"starfission_go_api/internal/dao"
	"starfission_go_api/pkg/app"
	"starfission_go_api/pkg/convert"
	"starfission_go_api/pkg/errcode"
	"starfission_go_api/pkg/gin_tools"

	"github.com/gin-gonic/gin"
)

//type SignParam struct {
//	AppID int64  `form:"app_id"  binding:"required,gte=1"` //APP_ID
//	Ts    int64  `form:"ts"  binding:"required"`           //APP_ID
//	Sign  string `form:"sign"  binding:"required"`         //APP_ID
//	Debug int64  `form:"debug"`                            //debug 不参数签名
//}

// Sign 签名
func Sign() gin.HandlerFunc {
	return func(c *gin.Context) {

		daov := dao.New(global.DBEngine)
		response := app.NewResponse(c)
		p, err := gin_tools.RequestParams(c)

		if err != nil {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
			c.Abort()
		}

		params := convert.MapAnyToMapStr(p)

		appId, _ := strconv.ParseInt(params["app_id"], 10, 64)

		if appId == 0 {
			response.ToErrorResponse(errcode.InvalidParamsAppId)
			c.Abort()

		}

		app, err := daov.GetAppByID(appId)

		if err != nil {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
			c.Abort()
		}

		if app.IsPublished == 1 {
			delete(params, "debug")
		}

		sign, verr := verifySign(params, app.AppSecret)
		if sign != nil {
			response.ToResponse(sign)
			c.Abort()
		}
		if verr != nil {
			response.ToErrorResponse(verr)
			c.Abort()
		}
		c.Next()
	}
}

// 验证签名
func verifySign(params map[string]string, appSecret string) (map[string]string, *errcode.Error) {

	debug := params["debug"]
	sign := params["sign"]
	ts := params["ts"]

	if sign == "" {
		return nil, errcode.SignParamMissing
	}

	if debug == "1" {

		var rets string

		if ts == "" {
			currentUnix := time.Now().Unix()
			rets = strconv.FormatInt(currentUnix, 10)
			params["ts"] = rets
		} else {
			rets = params["ts"]
		}

		res := map[string]string{
			"ts":       rets,
			"sign":     createSign(params, appSecret),
			"sign_str": createEncryptStr(params),
		}
		return res, nil
	}

	// 验证签名
	if sign == "" {
		return nil, errcode.SignParamMissing
	}

	// 验证过期时间
	timestamp := time.Now().Unix()
	tsInt, _ := strconv.ParseInt(ts, 10, 64)
	if tsInt > timestamp || timestamp-tsInt >= global.ServerSetting.SignExpiry {
		return nil, errcode.SignExpired
	}
	// 验证签名
	if sign != createSign(params, appSecret) {
		return nil, errcode.SignError
	}
	return nil, nil
}

// 创建签名
func createSign(params map[string]string, appSecret string) string {
	hash := md5.Sum([]byte(appSecret + createEncryptStr(params) + appSecret))
	return hex.EncodeToString(hash[:])
}

func createEncryptStr(params map[string]string) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sign" && k != "debug" {
			key = append(key, k)
		}
	}
	sort.Strings(key)

	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params[key[i]])
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params[key[i]])
		}
	}
	return str
}
