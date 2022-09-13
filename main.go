package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gotest/bootstrap"
)

func main() {
	// new一个gin engine 实例
	r := gin.New()

	bootstrap.SetupRoute(r)

	// 运行服务 默认为8080
	err := r.Run(":8000")
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
