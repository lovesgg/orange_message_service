package server

import (
	"github.com/kataras/iris/context"
	"orange_message_service/app/components/sms"
	models2 "orange_message_service/app/models"
	models "orange_message_service/app/models/request"
	"orange_message_service/app/services"
	"orange_message_service/app/services/server/smsTemplates"
)

func SmsSend(ctx context.Context, params models.ServerReq, channel models2.Channel) bool {
	var sendData map[string]interface{}
	switch channel.Template {
	case "SendDataTest":
		sendData = smsTemplates.SendDataTest(ctx, params)
		break
	case "OtherFunc":
		sendData = smsTemplates.OtherFunc(ctx, params)
		break
	}
	ret := sendSms(ctx, params.Body.Phone, sendData["templateCode"].(string) ,sendData)

	if ret {
		//已发送的可将数据入库
		services.InsertDataMysql(params, params.Body.Phone, 1)
		return true
	} else {
		return false
	}
}

func sendSms(ctx context.Context, phone string, templateCode string, sendData map[string]interface{}) bool {
	ret := sms.Send(phone, templateCode, sendData)
	return ret
}
