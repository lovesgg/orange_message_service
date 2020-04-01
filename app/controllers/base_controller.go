package controllers

import (
	"github.com/kataras/iris/context"
	"gopkg.in/go-playground/validator.v9"
	"orange_message_service/app/common"
	"orange_message_service/app/components/mlog"
	"time"
)

type BaseController struct {
}

//返回错误，并且打warn日志
func (BaseController) RenderError(ctx context.Context, error *common.ComError) {
	startTime := ctx.Values().Get(common.COMMON_START_TIME).(time.Time)
	now := time.Now()
	logTime := now.Sub(startTime).String()
	mlog.Warnf(ctx, "type=request_out_error||errno=%d||err=%v||log_time=%s", error.Code, error, logTime)

	//nolint:errcheck
	_, _ = ctx.JSON(map[string]interface{}{
		"start_time": startTime.Format("15:04:05.0000"),
		"end_time":   now.Format("15:04:05.0000"),
		"log_time":   logTime,
		"ret":        common.RET_ERROR,
		"error": map[string]interface{}{
			"msg":  error.Msg,
			"code": error.Code,
		},
	})

}

func (BaseController) RenderJson(ctx context.Context, data interface{}) {
	startTime := ctx.Values().Get(common.COMMON_START_TIME).(time.Time)
	now := time.Now()
	logTime := now.Sub(startTime).String()
	mlog.Debugf(ctx, "type=request_out_success||log_time=%s||ret=%#v", logTime, data)

	//nolint:errcheck
	_, _ = ctx.JSON(map[string]interface{}{
		"start_time": startTime.Format("15:04:05.0000"),
		"end_time":   now.Format("15:04:05.0000"),
		"log_time":   logTime,
		"ret":        common.RET_SUCCESS,
		"data":       data,
	})

}

func (BaseController) RenderRawJson(ctx context.Context, data map[interface{}]interface{}) {
	startTime := ctx.Values().Get(common.COMMON_START_TIME).(time.Time)
	now := time.Now()
	logTime := now.Sub(startTime).String()

	mlog.Debugf(ctx, "type=request_out_success||log_time=%s||ret=%#v", logTime, data)
	data["start_time"] = startTime.Format("15:04:05.000000")
	data["end_time"] = time.Now().Format("15:04:05.000000")
	data["log_time"] = logTime

	//nolint:errcheck
	_, _ = ctx.JSON(data)
}

// 获取请求数据对象
func (BaseController) GetRequest(ctx context.Context, req interface{}) {
	if err := ctx.ReadJSON(req); err != nil {
		errRet := common.GenError(common.ERROR_REQUEST_PARAMS, "请求参数错误", err)
		panic(errRet)
	}

	v := validator.New()
	errV := v.Struct(req)
	if errV != nil {
		errRet := common.GenError(common.ERROR_REQUEST_PARAMS, "请求参数错误", errV)
		panic(errRet)
	}
}
