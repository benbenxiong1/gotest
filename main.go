package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gotest/bootstrap"
	btsConfig "gotest/config"
	"gotest/pkg/config"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// new一个gin engine 实例
	r := gin.New()

	bootstrap.SetupRoute(r)

	// 运行服务 默认为8080
	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
