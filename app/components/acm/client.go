package acm

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	config2 "orange_message_service/app/components/config"
)

type Client struct {
}

func GetClient() config_client.IConfigClient {
	config := config2.GetConfig()

	var endpoint = config.GetString("acm.endpoint")
	var namespaceId = config.GetString("acm.namespaceId")
	var accessKey = config.GetString("acm.accessKey")
	var secretKey = config.GetString("acm.secretKey")

	clientConfig := constant.ClientConfig{
		Endpoint:       endpoint + ":8080",
		NamespaceId:    namespaceId,
		AccessKey:      accessKey,
		SecretKey:      secretKey,
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})

	if err != nil {
		fmt.Println(err)
		return configClient
	}

	return configClient
}

func GetConfig(dataId string, group string) string {
	group = "DEFAULT_GROUP"
	content, err := GetClient().GetConfig(vo.ConfigParam{DataId: dataId, Group: group})
	if err != nil {
		return ""
	}
	return content
}
