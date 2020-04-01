package services

import (
	"encoding/json"
	"github.com/kataras/iris/context"
	"orange_message_service/app/common"
	"orange_message_service/app/common/enum"
	config2 "orange_message_service/app/components/config"
	"orange_message_service/app/components/http"
	redis2 "orange_message_service/app/components/redis"
	models "orange_message_service/app/models/request"
)

/**
这里最好加一层缓存 因为是需要经常调用的
关于appid根据自己的实际需要 可以写在配置文件中读取 也可以在这写死
对于线上的access_token维护最好是
 */
func GetAccessTokenFromCache(ctx context.Context) string {
	var accessToken models.AccessToken
	var accessStr string

	config := config2.GetConfig()

	grantType := config.GetString("wechat.grantType")
	appId := config.GetString("wechat.appId")
	Secret := config.GetString("wechat.Secret")

	url := common.WECHAT_ACCESS_TOKEN_URL + "?grant_type=" + grantType + "&appid=" + appId + "&secret=" + Secret

	redis := redis2.GetCommonClient()
	cacheKey := enum.REIDS_KEY_WECHAT_ACCESS_TOKEN
	ret, err := redis.GetString(cacheKey)
	if err != nil || ret == "" {
		//没数据
		requestData := http.Post(url, map[string]interface{}{})
		_ = json.Unmarshal([]byte(requestData), &accessToken)
		accessStr = accessToken.AccessToken

		_ = redis.Set(cacheKey, accessStr, enum.REDIS_EXPIRE_NINE_MINUTE)
	}
	if ret != "" {
		accessStr = ret
	}


	return accessStr
}

