package server

import (
	config2 "orange_message_service/app/components/config"
	models2 "orange_message_service/app/models"
	models "orange_message_service/app/models/request"
	"strings"
)

var wordsJson map[string]string

/**
敏感词过滤 对输入内容字段都做遍历过滤
*/
func FilterWords(params models.ServerReq) models.ServerReq {
	var messageBody models2.Message
	messageBody = params.Body
	config := config2.GetConfig()
	wordsJson = config.GetStringMapString("filter_words")

	//依次在这加需要过滤的字段
	messageBody.GoodsName = filterText(messageBody.GoodsName)
	messageBody.AddressDetail = filterText(messageBody.AddressDetail)

	params.Body = messageBody

	return params
}

func filterText(text string) string {
	var str string
	texts := strings.Split(text, "")
	for _, text := range texts {
		_, ok := wordsJson[text]
		//fmt.Println(ok)
		if ok {
			text = wordsJson[text]
		}
		str += text
	}
	return str
}
