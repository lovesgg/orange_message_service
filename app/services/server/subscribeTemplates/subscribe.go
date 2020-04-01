package subscribeTemplates

import (
	"github.com/kataras/iris/context"
	"orange_message_service/app/common"
	models "orange_message_service/app/models/request"
)

/**
测试模板数据组装 微信订阅消息参考腾讯开发文档 选择对应的模板即可
*/
func SendDataTest(ctx context.Context, params models.ServerReq) map[string]interface{} {
	return map[string]interface{}{
		"touser":      params.Body.UserId,
		"template_id": common.TEMPLATE_VIRTUAL_ORDER,
		"page":        "pages/index/index",
		"data": map[string]interface{}{
			"amount1": map[string]interface{}{
				"value": "20",
			},
			"thing2": map[string]interface{}{
				"value": "订单",
			},
			"date3": map[string]interface{}{
				"value": "2020-01-02 21:00:00",
			},
			"thing4": map[string]interface{}{
				"value": "222",
			},
			"thing5": map[string]interface{}{
				"value": "33",
			},
		},
	}
}

func OtherFunc(ctx context.Context, params models.ServerReq) map[string]interface{} {
	return map[string]interface{}{
		"touser":      params.Body.UserId,
		"template_id": common.TEMPLATE_VIRTUAL_ORDER,
		"page":        "pages/index/index",
		"data": map[string]interface{}{
			"phrase2": map[string]interface{}{
				"value": "",
			},
			"thing7": map[string]interface{}{
				"value": "",
			},
			"phone_number9": map[string]interface{}{
				"value": "",
			},
		},
	}
}
