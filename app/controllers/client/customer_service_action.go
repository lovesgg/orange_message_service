package controllers

import (
	"github.com/kataras/iris/context"
	models "orange_message_service/app/models/request"
	clientService "orange_message_service/app/services/client"
)

func (c *CustomerController) Say(ctx context.Context) {
	var req models.CustomerSayReq
	c.GetRequest(ctx, &req)

	ret := clientService.CustomerSay(ctx, req)
	c.RenderJson(ctx, ret)
	return
}

