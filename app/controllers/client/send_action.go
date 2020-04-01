package controllers

import (
	"github.com/kataras/iris/context"
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
