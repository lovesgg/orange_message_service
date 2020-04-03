package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

//这里是conf/目录下的文件名
var CONFIG_FILES = [10]string{
	"app",
	"redis",
	"rpc",
	"log",
	"message",
	"mysql",
	"wechat",
	"email",
	"common",
	"filter_words",
}

var config *viper.Viper

func Init() {
	var err error
	config = viper.New()
	config.SetConfigType("json")

	serviceEnv := os.Getenv("env_orange_message_service")
	if serviceEnv == "" {
		serviceEnv = "prod"
	}
	fmt.Println("当前的环境是：" + serviceEnv)
	for _, name := range CONFIG_FILES {
		config.SetConfigName(name)
		if serviceEnv == "dev" {
			config.AddConfigPath("./conf")
		} else {
			config.AddConfigPath("/data/www/orange_message_service/conf")
		}
		err = config.MergeInConfig()
		if err != nil {
			panic(fmt.Sprintf("error on parsing configuration file.||name=%s||err=%#v\n", name, err))
		}
	}
}

func GetConfig() *viper.Viper {
	return config
}

func GetString(key string) string {
	return config.GetString(key)
}

func GetStringSlice(key string) []string {
	return config.GetStringSlice(key)
}
