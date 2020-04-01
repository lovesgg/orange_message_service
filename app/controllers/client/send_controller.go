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
