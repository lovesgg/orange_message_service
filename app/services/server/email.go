package server

import (
	"github.com/kataras/iris/context"
	"orange_message_service/app/components/go-email"
	models2 "orange_message_service/app/models"
	models "orange_message_service/app/models/request"
	"orange_message_service/app/services"
	"orange_message_service/app/services/server/emailTemplates"
)

func EmailSend(ctx context.Context, params models.ServerReq, channel models2.Channel) bool {
	var sendData map[string]interface{}
	switch channel.Template {
	case "SendDataTest":
		sendData = emailTemplates.SendDataTest(ctx, params)
		break
	case "OtherFunc":
		sendData = emailTemplates.OtherFunc(ctx, params)
		break
	}
	ret := sendEmail(ctx, params.Body.Email, sendData)

	if ret {
		//已发送的可将数据入库
		services.InsertDataMysql(params)
		return true
	} else {
		return false
	}
}

func sendEmail(ctx context.Context, email string, sendData map[string]interface{}) bool {
	var sendTo []string
	sendTo = append(sendTo, email)
	ret := go_email.SendEmail(sendTo, sendData["title"].(string), sendData["json"].(string))
	if ret {
		return true
	} else {
		return false
	}
}
