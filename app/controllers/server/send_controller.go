package controllers

import (
	"github.com/kataras/iris/context"
	"orange_message_service/app/controllers"
)

var server ServerController

type ServerController struct {
	controllers.BaseController
}

func SendHandle(ctx context.Context) {
	server.Send(ctx)
}
