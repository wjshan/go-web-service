package apps

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-service-gin/apps/v1"
)

func ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func RegisterRoute(route *gin.Engine) {
	v1Group := route.Group("/v1")
	{
		v1Group.GET("/ping", ping)
		v1.RegisterRoute(v1Group)
	}

}
