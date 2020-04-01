package middleware

import (
	"github.com/kataras/iris/context"
	"orange_message_service/app/common"
	"time"
)

func RequestBootstrap() context.Handler {
	return func(ctx context.Context) {
		ctx.Values().Set(common.COMMON_START_TIME, time.Now())
		ctx.Next()
	}
}
