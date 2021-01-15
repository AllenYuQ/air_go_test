package routers

import (
	"github.com/gin-gonic/gin"
	"task5/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	//省略一些中间件的设置和分组设计

	r.LoadHTMLGlob("./templates/*")
	r.GET("/user/list", listUserHandler)
	r.GET("/user/login", loginGet)
	r.POST("/user/login", loginPost)
	r.GET("/predictions/check", serviceGet)
	r.POST("/predictions/check", servicePost)
	r.POST("/go_python", goPythonPost)

	return r
}
