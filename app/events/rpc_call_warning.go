package events

import (
	"orange_message_service/app/common"
	"orange_message_service/app/common/enum"
	"orange_message_service/app/components/eventdispatcher"
	"orange_message_service/app/components/mlog"
)

type RpcCallWarningEventData struct {
	CtxFields []interface{}
	Url       string
	Errno     common.ErrorCode
	ReqBody   string
	StartTime string
	EndTime   string
	CostTime  string
	Response  string
}

func NewRpcCallWarningEvent(ctxFields []interface{}, url string, errno common.ErrorCode, reqBody string, startTime string, endTime string, costTime string, response string) *RpcCallWarningEvent {
	data := &RpcCallWarningEventData{CtxFields: ctxFields, Url: url, Errno: errno, ReqBody: reqBody, StartTime: startTime, EndTime: endTime, CostTime: costTime, Response: response}
	return &RpcCallWarningEvent{
		Name: enum.EVENT_NAME_RPC_CALL_WARNING,
		Data: data,
	}
}

type RpcCallWarningEvent struct {
	Name string
	Data *RpcCallWarningEventData
}

func (r RpcCallWarningEvent) GetName() string {
	return r.Name
}

func (r RpcCallWarningEvent) GetData() interface{} {
	return r.Data
}

var LogWarningRpcCallListener = eventdispatcher.NewListener(func(event eventdispatcher.EventInterface) error {
	data := event.GetData().(*RpcCallWarningEventData)
	mlog.WarnfWithFields(data.CtxFields, "rpc request %s.||errno=%d||params=%s||start_time=%s||end_time=%s||cost_time=%s||response=%s",
		data.Url,
		data.Errno,
		data.ReqBody,
		data.StartTime,
		data.EndTime,
		data.CostTime,
		data.Response,
	)
	return nil
}, false)
