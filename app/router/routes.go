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

	//客户端批量发送 使用协程 提高并发量 比以上接口速度更快。强烈推荐使用这方法来发送消息
	app.Post("/client/send-by-sync", clientController.SendBySync)

	/**
	{
		"msg_key":1003,
		"source_id":1,
		"users":[
			"993285153@qq.com",
			"773668121@qq.com"
		]
	}
	 */
	//客户端只根据用户账号(手机号|邮箱|微信id)和模板发送 不传冗余数据 这接口更像是给n个用户发送一个一模一样的信息数据，比如给几百上千个用户同时发送优惠券信息
	app.Post("/client/send-by-users", clientController.SendByUsers)

	//服务端接mq实时消费
	app.Post("/server/send", serverController.SendHandle)
}
