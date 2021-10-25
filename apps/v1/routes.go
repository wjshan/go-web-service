package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoute(route * gin.RouterGroup)   {
	route.GET("/user_info", func(context *gin.Context) {
		context.JSONP(http.StatusOK, gin.H{"user_id":1234})
	})
}