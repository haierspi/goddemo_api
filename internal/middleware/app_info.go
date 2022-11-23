package middleware

import (
	"starfission_go_api/pkg/app"

	"github.com/gin-gonic/gin"
)

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "GoApi")
		c.Set("app_version", "1.0.0")
		c.Set("access_host", app.GetAccessHost(c))

		c.Next()
	}
}
