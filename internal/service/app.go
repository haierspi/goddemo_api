package service

type AppRequest struct {
	AppID int64 `form:"app_id"  binding:"required,gte=1"` //APP_ID
}
