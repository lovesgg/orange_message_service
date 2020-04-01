package commands

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"net/http"
)

var Ctx iris.Context

func init() {
	app := iris.Default()
	Ctx = context.NewContext(app)
	Ctx.BeginRequest(nil, &http.Request{})
}
