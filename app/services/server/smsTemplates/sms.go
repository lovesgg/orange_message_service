package smsTemplates

import (
	"github.com/kataras/iris/context"
	models "orange_message_service/app/models/request"
)

/**
测试模板数据组装 注意有的是必填信息 不传发不出去
*/
func SendDataTest(ctx context.Context, params models.ServerReq) map[string]interface{} {
	return map[string]interface{}{
		"phone": params.Body.Phone,
		"templateCode": "SMS_174986902",
		"server_name": "orange",
		"db": "message",
		"sign_name": "",//不能为空
	}
}

func OtherFunc(ctx context.Context, params models.ServerReq) map[string]interface{} {
	return map[string]interface{}{
		"phone": params.Body.Phone,
		"templateCode": "SMS_186950738",
		"sign_name": "",//不能为空
	}
}
