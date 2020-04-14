package emailTemplates

import (
	"encoding/json"
	"github.com/kataras/iris/context"
	models "orange_message_service/app/models/request"
)

/**
测试模板数据组装
*/
func SendDataTest(ctx context.Context, params models.ServerReq) map[string]interface{} {
	jsons, _ := json.Marshal(params)
	return map[string]interface{}{
		"title": params.Body.GoodsName,
		"json":  string(jsons),
	}
}

func OtherFunc(ctx context.Context, params models.ServerReq) map[string]interface{} {
	return map[string]interface{}{
		"title": params.Body.GoodsName,
		"json":  "",
	}
}

func SendByUsers(ctx context.Context, params models.ServerReq) map[string]interface{} {
	return map[string]interface{}{
		"title": "每个用户都不一样",
		"json":  "test",
	}
}
