package router

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	clientController "orange_message_service/app/controllers/client"
	serverController "orange_message_service/app/controllers/server"
)

func RegisterRoutes(app *iris.Application) {
	//健康检查
	app.Get("/health/check", func(ctx context.Context) {
		_, _ = ctx.WriteString("200")
	})
	//
	app.Get("/", func(ctx context.Context) {
		_, _ = ctx.WriteString("welcome app ")
	})

	//客户端接收消息 发送mq 适合消息量少的
	app.Post("/client/send", clientController.SendHandler)

	//客户端接收消息 发送mq 适合消息量大的 比如百万级 只是做个模拟。具体细节需要自己实现
	app.Post("/client/send-batch", clientController.SendBatchHandler)

	//服务端接mq实时消费
	app.Post("/server/send", serverController.SendHandle)
}
