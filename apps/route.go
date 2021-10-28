package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"web-service-gin/apps/v1"
	"web-service-gin/models"
	"web-service-gin/units"
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
	route.POST("/db-migrate", autoMigrate)

}

func autoMigrate(ctx *gin.Context) {
	db := units.DBInstance.GetPGSQL()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		message := fmt.Sprintf("迁移失败%s", err.Error())
		ctx.JSONP(http.StatusOK, gin.H{"code": -1, "message": message, "data": nil})
	}
	ctx.JSONP(http.StatusOK, gin.H{"code": 0, "message": nil, "data": nil})
}
