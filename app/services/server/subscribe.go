package server

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/context"
	"orange_message_service/app/common"
	"orange_message_service/app/components/http"
	models2 "orange_message_service/app/models"
	models "orange_message_service/app/models/request"
	"orange_message_service/app/services"
	"orange_message_service/app/services/server/subscribeTemplates"
)

func SubscribeSend(ctx context.Context, params models.ServerReq, channel models2.Channel) bool {
	var sendData map[string]interface{}

	//获取acces token
	accessToken := services.GetAccessTokenFromCache(ctx)
	fmt.Println("acc", accessToken)
	if accessToken == "" {
		return false
	}

	channel.Template = "SendDataTest"
	switch channel.Template {
	case "SendDataTest":
		sendData = subscribeTemplates.SendDataTest(ctx, params)
		break
	case "OtherFunc":
		sendData = subscribeTemplates.OtherFunc(ctx, params)
		break
	}
	ret := sendWechat(ctx, sendData, accessToken)

	if ret {
		_ = services.InsertDataMysql(params, params.Body.UserId, 1)
		//已发送的可将数据入库
		return true
	} else {
		return false
	}
}

func sendWechat(ctx context.Context, sendData map[string]interface{}, accessToken string) bool {
	var res models.SendRes
	wxres := http.Post(common.WECHAT_SEND_SUBSCRIBE_URL+"?access_token="+accessToken, sendData)
	//解析结果 具体参考微信文档
	_ = json.Unmarshal([]byte(wxres), &res)
	if res.Errcode == 0 {
		return true
	} else {
		return false
	}
}
