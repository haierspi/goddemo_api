package global

import (
	"starfission_go_api/pkg/logger"
	"starfission_go_api/pkg/setting"
)

var (
	ServerSetting         *setting.ServerSettingS
	AppSetting            *setting.AppSettingS
	EmailSetting          *setting.EmailSettingS
	JWTSetting            *setting.JWTSettingS
	SecuritySetting       *setting.SecuritySettingS
	DatabaseSetting       *setting.DatabaseSettingS
	RedisSetting          *setting.RedisS
	RabbitMQSetting       *setting.RabbitMQS
	WechatJSAPIPaySetting *setting.WechatPaySettingS
	WechatH5PaySetting    *setting.WechatPaySettingS
	AlipayH5PaySetting    *setting.AlipaySettingS
	Logger                *logger.Logger
	OSSSetting            *setting.OSSSettingS
)
