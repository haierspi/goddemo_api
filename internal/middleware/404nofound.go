package middleware

import (
	"net/http"

	"starfission_go_api/pkg/app"

	"github.com/gin-gonic/gin"
)

func NoFound() gin.HandlerFunc {
	return func(c *gin.Context) {

		response := app.NewResponse(c)
		response.SendResponse(http.StatusOK, app.ErrResult{
			Code:   http.StatusNotFound,
			Status: http.StatusNotFound,
			Msg:    "API No Found",
			Data:   c.Request.RequestURI,
		})
	}
}
