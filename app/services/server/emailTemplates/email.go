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
		"json": jsons,
	}
}

func OtherFunc(ctx context.Context, params models.ServerReq) map[string]interface{} {
	return map[string]interface{}{
		"title": params.Body.GoodsName,
	}
}

