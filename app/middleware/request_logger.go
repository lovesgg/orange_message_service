package middleware

import (
	"bytes"
	"github.com/kataras/iris/context"
	"io/ioutil"
	"orange_message_service/app/common"
	"orange_message_service/app/components/mlog"
	"orange_message_service/app/utils"
)

func NewRequestLogHandler() context.Handler {
	return func(ctx context.Context) {
		req := ctx.Request()
		comFields := common.CommonLogFields{
			IP:      ctx.RemoteAddr(),
			Method:  ctx.Method(),
			Path:    req.URL.RequestURI(),
			TraceID: ctx.GetHeader("trace_id"),
			Header:  ctx.Request().Header,
		}

		// 设置日志公共字段
		ctx.Values().Set(common.COMMON_LOG_FIELD_KEY, comFields)

		// 黑名单接口不写日志
		if !utils.InStringArray(comFields.Path, common.REQUEST_IN_LOG_BLACKLIST) {
			oldBody := req.Body
			buf, err := ioutil.ReadAll(oldBody)
			defer oldBody.Close()
			if err != nil {
				mlog.Warnf(ctx, "read req body error||err=%#v", err)
				ctx.Next()
				return
			}

			body := ioutil.NopCloser(bytes.NewBuffer(buf))
			mlog.Debugf(ctx, "type=request_in||request_body=%s", string(buf))
			req.Body = body
		}

		ctx.Next()
	}
}
