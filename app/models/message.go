package models

/**
商品名
社群名
地址
手机
提示
用户id:手机 短信 邮件
订单号
*/
type Message struct {
	GoodsName     string `json:"goods_name"`
	StoreName     string `json:"store_name"`
	AddressDetail string `json:"address_detail"`
	Phone         string `json:"phone"`
	Note          string `json:"note"`
	UserId        string `json:"user_id"`
	OrderNo       string `json:"order_no"`
	Email         string `json:"email"`
}

/**
发送消息管道 顺序的解析
*/
type Sequence struct {
	Sequence  map[int]string `json:"sequence"`
	Sms       Channel
	Subscribe Channel
}

type Channel struct {
	Channel    string `json:"channel"`
	IsRetry    int    `json:"is_retry"`
	Template string `json:"template"`
}
