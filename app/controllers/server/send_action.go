package controllers

import (
	"github.com/goinggo/mapstructure"
	"github.com/kataras/iris/context"
	"orange_message_service/app/components/config"
	"orange_message_service/app/models"
	modelsReq "orange_message_service/app/models/request"
	servers "orange_message_service/app/services/server"
	"strconv"
)

/**
可以用协程来提高高并发
*/
func (c ServerController) Send(ctx context.Context) {
	var req modelsReq.ServerReq

	c.GetRequest(ctx, &req)

	config := config.GetConfig()
	sequenceData := config.GetStringMap(strconv.Itoa(req.MsgKey)) //整个msg_id  对应的配置
	//遍历发送通道 这里可以写成独立方法 is_retry为1时可以递归调用重复发送
	for _, sequence := range sequenceData["sequence"].([]interface{}) {
		var ret bool
		var channel models.Channel
		channelData := sequenceData[sequence.(string)]
		_ = mapstructure.Decode(channelData, &channel)
		switch sequence {
		case "subscribe":
			ret = servers.SubscribeSend(ctx, req, channel)
			break
		case "sms":
			ret = servers.SmsSend(ctx, req, channel)
			break
		default:
			ret = servers.SubscribeSend(ctx, req, channel)
		}
		if ret == false {
			c.RenderJson(ctx, "发送有误")
			return
		}
	}

	c.RenderJson(ctx, "send ok")
}
