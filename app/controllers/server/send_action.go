package controllers

import (
	"fmt"
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

	fmt.Println("server端开始消费消息")

	//如果不需要过滤敏感词可注释这行 过滤敏感词会耗费部分时间哟
	req = servers.FilterWords(req)

	config := config.GetConfig()
	sequenceData := config.GetStringMap(strconv.Itoa(req.MsgKey)) //整个msg_id  对应的配置
	//遍历发送通道 这里可以写成独立方法 is_retry为1时可以递归调用重复发送
	for _, sequence := range sequenceData["sequence"].([]interface{}) {
		var ret bool
		var channel models.Channel
		channelData := sequenceData[sequence.(string)]
		_ = mapstructure.Decode(channelData, &channel)
		fmt.Println("msg_key:", req.MsgKey)
		fmt.Println("通道:", sequence)
		fmt.Println("模板:", channel.Template)
		switch sequence {
		case "subscribe":
			ret = servers.SubscribeSend(ctx, req, channel)
			break
		case "sms":
			ret = servers.SmsSend(ctx, req, channel)
			break
		case "email":
			ret = servers.EmailSend(ctx, req, channel)
			break
		default:
			ret = servers.SubscribeSend(ctx, req, channel)
		}
		if ret == false {
			c.RenderJson(ctx, "发送有误")
			return
		}
	}
	fmt.Println("server端消费结束。消息发送成功。")

	c.RenderJson(ctx, "send ok")
}
