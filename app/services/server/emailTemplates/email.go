package emailTemplates

import (
	"github.com/kataras/iris/context"
	models "orange_message_service/app/models/request"
)

/**
测试模板数据组装
*/
func SendDataTest(ctx context.Context, params models.ServerReq) map[string]interface{} {
	return map[string]interface{}{
		"phone": params.Body.Phone,
	}
}

func OtherFunc(ctx context.Context, params models.ServerReq) map[string]interface{} {
	return map[string]interface{}{
		"phone": params.Body.Phone,
	}
}

