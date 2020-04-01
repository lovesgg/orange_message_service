package services

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/context"
	"io"
	"orange_message_service/app/common"
	"orange_message_service/app/common/enum"
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

		jsonData, _ := json.Marshal(postData)
		fmt.Println(string(jsonData))
		mcqKey := body.UserId + body.OrderNo; //mcq的key
		fmt.Println(mcqKey)
		fmt.Println(enum.MCQ_MESSAGE_CLIENT_CENTER)
		//执行mq发送 由server端来消费
	}

	return true
}
