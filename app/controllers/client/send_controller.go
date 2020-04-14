package controllers

import (
	"github.com/kataras/iris/context"
	"orange_message_service/app/controllers"
)

var client ClientController

type ClientController struct {
	controllers.BaseController
}

func SendHandler(ctx context.Context) {
	client.Send(ctx)
}

func SendBatchHandler(ctx context.Context) {
	client.SendBatch(ctx)
}

func SendBySync(ctx context.Context)  {
	client.SendBySync(ctx)
}

func SendByUsers(ctx context.Context)  {
	client.SendByUsers(ctx)
}