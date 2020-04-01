package rpcComponent

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/kataras/iris/context"
	"orange_message_service/app/common"
	"orange_message_service/app/components/eventdispatcher"
	"orange_message_service/app/components/http"
	"orange_message_service/app/components/mlog"
	"orange_message_service/app/events"
)

type BaseRpc struct {
	BaseAddr string
}

func (r BaseRpc) GetUrl(index RPC_REQUEST_INDEX) string {
	rpcConfig := RpcConfigs[index]
	return r.BaseAddr + rpcConfig.Uri
}

func (r BaseRpc) Request(ctx context.Context, rpcIndex RPC_REQUEST_INDEX, req interface{}, res interface{}) (common.ErrorCode, []byte) {
	rpcConfig := RpcConfigs[rpcIndex]
	url := r.BaseAddr + rpcConfig.Uri
	return r.sendRequest(ctx, url, req, res, rpcConfig)
}

func (BaseRpc) sendRequest(ctx context.Context, url string, req interface{}, res interface{}, config RpcConfig) (common.ErrorCode, []byte) {
	// 序列化 post_body
	reqBody, errMarshal := jsoniter.Marshal(req)
	if errMarshal != nil {
		mlog.Warnf(ctx, "request marshal error||url=%s||url=%s||err=%v", url, errMarshal.Error())
		return common.ERROR_MARSHAL, nil
	}

	response, errReq := http.PostJson(ctx, url, reqBody, config.Timeout)
	if errReq != nil {
		mlog.Warnf(ctx, "request unknow error||url=%s||params=%s||err=%v", url, string(reqBody), errReq.Error())
		return common.ERROR_REQUEST_UNKNOW, nil
	}
	if response.StatusCode == 500 {
		return common.ERROR_REQUEST_INTERNAL_ERROR_500, response.Body
	}
	// debug 日志
	mlog.Debugf(ctx, "rpc request %s.||params=%s||start_time=%s||end_time=%s||cost_time=%s||response=%s",
		url,
		string(reqBody),
		response.RequestTime.Format(common.FORMAT_YYYY_MM_DD_HH_II_SS_DETAIL),
		response.ReceivedAt.Format(common.FORMAT_YYYY_MM_DD_HH_II_SS_DETAIL),
		response.Time,
		string(response.Body),
	)
	// warning 日志
	if response.Time > config.WaringDuration {
		event := events.NewRpcCallWarningEvent(
			mlog.GetComFields(ctx),
			url,
			common.ERROR_REQUEST_WARNING_DURATION,
			string(reqBody),
			response.RequestTime.Format(common.FORMAT_YYYY_MM_DD_HH_II_SS_DETAIL),
			response.ReceivedAt.Format(common.FORMAT_YYYY_MM_DD_HH_II_SS_DETAIL),
			response.Time.String(),
			string(response.Body),
		)
		_ = eventdispatcher.Dispatcher(event) //ignore error handle.
	}

	// 反序列化 response_body
	errUnMarshal := jsoniter.Unmarshal(response.Body, res)
	if errUnMarshal != nil {
		mlog.Warnf(ctx, "unmarshal rpc response error||url=%s||params=%s||err=%v", url, string(reqBody), errUnMarshal.Error())
		return common.ERROR_UNMARSHAL, response.Body
	}
	return common.ERROR_OK, response.Body
}
