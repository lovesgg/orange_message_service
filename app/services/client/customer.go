package services

import (
	"encoding/json"
	"github.com/kataras/iris/context"
	"orange_message_service/app/components/acm"
	models "orange_message_service/app/models/request"
	"strings"
)

/**
acm配置信息 举例
{
    "words": [
        {
            "title":"发货",
            "detail":"20日之后发货"
        },
        {
            "title":"物流哈哈哈",
            "detail":"20日之后发运 顺丰"
        }
    ]
}
*/

func CustomerSay(ctx context.Context, req models.CustomerSayReq) []map[string]string {
	var words map[string][]map[string]string
	var returnData []map[string]string

	getConfig := acm.GetConfig("test", "DEFAULT_GROUP")
	_ = json.Unmarshal([]byte(getConfig), &words)
	data := words["words"]

	for _, value := range data {
		//匹配子字符串 有的话数组返回
		index := strings.Index(value["title"], req.Text)
		if index >= 0 {
			word := map[string]string{
				"title":  value["title"],
				"detail": value["detail"],
			}
			returnData = append(returnData, word)
		}
	}

	return returnData
}
