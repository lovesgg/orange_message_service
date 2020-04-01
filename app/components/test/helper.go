package test

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"orange_message_service/app/components"
	"orange_message_service/app/components/config"
	"orange_message_service/app/components/mlog"
	"orange_message_service/app/components/redis"
	"net/http"
)

var app *iris.Application

func TestInit() {
	config.Init()
	mlog.Init()
	redis.Init()
	components.EventInit()
}

func NewCtx() context.Context {
	if app == nil {
		app = iris.Default()
	}
	ctx := context.NewContext(app)
	ctx.BeginRequest(nil, &http.Request{
		Header: map[string][]string{
			"mj-trace-id": {"trace_id"},
		},
	})
	return ctx
}

func Dump(args ...interface{}) {
	spew.Dump(args...)
}
