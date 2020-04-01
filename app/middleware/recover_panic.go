package middleware

import (
	"bytes"
	"fmt"
	"github.com/kataras/iris/context"
	"orange_message_service/app/common"
	"orange_message_service/app/components/mlog"
	"orange_message_service/app/controllers"
	"runtime"
)

func NewRecoverPanic() context.Handler {
	return func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() {
					return
				}

				var stacktrace string
				for i := 1; ; i++ {
					_, f, l, got := runtime.Caller(i)
					if !got {
						break
					}

					stacktrace += fmt.Sprintf("%s:%d\n", f, l)
				}

				var logMessage bytes.Buffer
				// when stack finishes
				logMessage.WriteString(fmt.Sprintf("Recovered from a route's Handler('%s')\n", ctx.HandlerName()))
				logMessage.WriteString(fmt.Sprintf("At Request: %s\n", ctx.RemoteAddr()))
				logMessage.WriteString(fmt.Sprintf("Trace: %s\n", err))
				logMessage.WriteString(fmt.Sprintf("\n%s", stacktrace))

				if businessErr, ok := err.(*common.ComError); ok {
					businessErr.SetTrace(logMessage.String())
					new(controllers.BaseController).RenderError(ctx, businessErr)
				} else {
					mlog.Warn(ctx, logMessage.String())
					ctx.StatusCode(500)
				}

				ctx.StopExecution()
			}
		}()

		ctx.Next()
	}
}
