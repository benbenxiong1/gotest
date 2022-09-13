package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	// new一个gin engine 实例
	r := gin.New()

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 注册路由
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"欢迎你": "hello word",
		})
	})

	// 处理404请求
	r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回404")
		} else {
			//默认返回json
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

	// 运行服务 默认为8080
	r.Run(":8000")
}
