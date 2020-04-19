package services

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/context"
	"orange_message_service/app/components/acm"
	models "orange_message_service/app/models/request"
)

/**
acm配置信息 举例

{
    "test":{
        "text":"你好"
    },
	"order":{
		"text":"订单不存在"
    }
}
*/

func CustomerSay(ctx context.Context, req models.CustomerSayReq) string {
	var words map[string]map[string]string

	getConfig := acm.GetConfig("test", "DEFAULT_GROUP")
	_ = json.Unmarshal([]byte(getConfig), &words)
	data, ok := words[req.Text]
	fmt.Println(data, ok)

	if ok {
		return data["text"] //取的是配置里的自动回复内容
	} else {
		fmt.Println("##########")
		return "暂时没有回答"
	}
}
