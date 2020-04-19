package controllers

import (
	"github.com/kataras/iris/context"
	"orange_message_service/app/controllers"
)

var customer CustomerController

type CustomerController struct {
	controllers.BaseController
}

func CustomerSayHandler(ctx context.Context) {
	customer.Say(ctx)
}
