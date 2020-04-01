package main

import (
	"fmt"
	"github.com/kataras/iris"
	"orange_message_service/app/components"
	"orange_message_service/app/components/config"
	"orange_message_service/app/middleware"
	"orange_message_service/app/router"
	"os"
)

func main() {
	//加载配置
	config.Init()
	//加载组件
	components.Init()
	//服务启动
	run()

}

func run() {
	config := config.GetConfig()

	app := iris.New()
	app.Use(middleware.RequestBootstrap())
	app.Use(middleware.NewRecoverPanic())
	app.Use(middleware.NewRequestLogHandler())

	router.RegisterRoutes(app)

	port := config.GetString("port")
	err := app.Run(iris.Addr(":"+port), iris.WithConfiguration(iris.Configuration{
		EnableOptimizations: true, //使用 jsoniter readJson
	}))
	if err != nil {
		fmt.Printf("run server failed. err=%v\n", err)
		os.Exit(1)
	}
}
