package middleware

import (
	"github.com/gin-gonic/gin"
)

func DisEncrypt() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dis_encrypt", 1)
		c.Next()
	}
}
