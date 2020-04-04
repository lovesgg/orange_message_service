package controllers

import (
	"github.com/kataras/iris/context"
	models2 "orange_message_service/app/models"
	models "orange_message_service/app/models/request"
	clientService "orange_message_service/app/services/client"
)

func (c *ClientController) Send(ctx context.Context) {
	var req models.SendReq
	c.GetRequest(ctx, &req)

	if len(req.Body) > 100 {
		c.RenderJson(ctx, "发送失败")
	}
	ret := clientService.SendMessage(ctx, req)
	if ret {
		c.RenderJson(ctx, "send ok")
		return
	} else {

		c.RenderJson(ctx, "发送失败")
		return
	}
}

/**
原理仍然是批量获取用户id，然后组装参数。往mq队列里塞信息。
其实最理想的方法是把手机号等都先放队列里。由server端直接来消费。
根据实际业务需要自己优化吧。每个场景都不一样。这里就不做展开了。
*/
func (c *ClientController) SendBatch(ctx context.Context) {
	var req models.SendReq
	var message []models2.Message

	message[0].Note = "hello"
	req.SourceId = 1111
	req.MsgKey = 1000

	//批量遍历用户手机或者邮箱
	users := []string{"1", "2", "3", "4", "5"}

	for _, toUser := range users {
		message[0].Phone = toUser
		req.Body = message
		ret := clientService.SendMessage(ctx, req)
		if ret != true {
			c.RenderJson(ctx, "send error")
			return
		}
	}

	c.RenderJson(ctx, "send all ok")
}
