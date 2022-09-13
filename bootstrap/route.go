// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gotest/routes"
	"net/http"
	"strings"
)

func SetupRoute(e *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(e)

	//注册api路由
	routes.RegisterAPIRoutes(e)

	//注册404
	setup404Handler(e)

}

func registerGlobalMiddleWare(e *gin.Engine) {
	e.Use(gin.Logger(), gin.Recovery())
}

func setup404Handler(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面不存在")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
