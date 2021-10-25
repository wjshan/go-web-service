package main

import (
	//"database/sql"
	"github.com/gin-gonic/gin"
	apps "web-service-gin/apps"
	"web-service-gin/units"
)

func main() {
	// 创建数据库链接
	dsn := "host=localhost user=dandan password=dandan dbname=golang_test port=5432 sslmode=disable TimeZone=UTC"
	errPg := units.InitPostgresDb(dsn)
	if errPg != nil {
		return
	}
	r := gin.Default()
	// 注册路由
	apps.RegisterRoute(r)
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
