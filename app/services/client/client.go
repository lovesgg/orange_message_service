package services

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/context"
	"io"
	"orange_message_service/app/common"
	"orange_message_service/app/common/enum"
	"orange_message_service/app/components/http"
	"orange_message_service/app/components/mlog"
	"orange_message_service/app/components/redis"
	models "orange_message_service/app/models/request"
)

func SendMessage(ctx context.Context, req models.SendReq) bool {
	str, _ := json.Marshal(req)
	w := md5.New()
	_, _ = io.WriteString(w, string(str))
	//将str写入到w中
	md5str := fmt.Sprintf("%x", w.Sum(nil))

	//写入缓存
	redisClient := redis.GetCommonClient()
	cacheKey := enum.REDIS_KEY_MJYX_MESSAGE_IDEMPOTENT + md5str
	ret, err := redisClient.Exists(cacheKey)
	if err != nil {
		mlog.Warnf(ctx, "redis send msg error error.||key=%s||errno=%d||err=%#v", cacheKey, common.ERROR_REDIS_CALL_ERROR, err)
	}
	if ret > 0 {
		return false //已经存在
	}
	setErr := redisClient.Set(cacheKey, 1, enum.REDIS_EXPIRE_NINE_SECONDS)
	if setErr != nil {
		mlog.Warnf(ctx, "redis set msg error error.||key=%s||errno=%d||err=%#v", cacheKey, common.ERROR_REDIS_CALL_ERROR, err)
		return false //连接有误
	}

	for _, body := range req.Body {
		postData := make(map[string]interface{})
		postData["msg_key"] = req.MsgKey
		postData["source_id"] = req.SourceId
		postData["body"] = body

		mcqKey := body.UserId + body.OrderNo; //mcq的key
		fmt.Println(mcqKey)
		fmt.Println(enum.MCQ_MESSAGE_CLIENT_CENTER)
		fmt.Println("开始发送mq")

		//执行mq发送 由server端来消费 现在这方法是模拟生产消费同步调用的。需要您自己搭建mq服务发送topic然后/server/send来消费
		//关于这部分过程的调用请看根目录下的.doc目下的图片哈.这么调用也可发送。就是显得很冗余。因为加了层http调用了。你要是觉得调内网这部分耗时可忽略那就这么用吧
		ret := sendMqTest(postData)
		fmt.Println(ret)
		//sendMq(postData) //生产环境先实现这方法
	}

	return true
}

//测试方法 模拟消息队列消费 生产环境不能这么用哈
func sendMqTest(params map[string]interface{}) bool {
	//需要换成您自己的url和端口 (这就是你运行这项目的ip+port)
	ret := http.Post("http://127.0.0.1:2195/server/send", params)
	fmt.Println("server端消息发送结果:" ,ret)
	return true
}

//这里实现mq生产方法  然后验证/server/send是否接收到并能发送就可以了
func sendMq(params map[string]interface{}) bool {
	return true
}