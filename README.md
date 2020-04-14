# Orange 消息推送系统
  Orange是由golang、iris、mysql、redis技术框架组成的消息推送服务。
  支持微信订阅消息、短信消息、邮件消息的统一可配置管理发送。支持不同策略的发送。有效管理模板资源，敏感词过滤，自定义追加配置和模板策略，记录消息体，可以为多个业务方向提供消息发送服务。测试后可线上使用。
  
  为什么需要这样的服务？
  当公司的业务变多，服务不断增加，需要为用户发送的消息越来越多。需要为用户发送譬如下单通知、到货通知、评价消息、接单成功通知、注册消息、其他推广信息。而众所周知利用第三方发送接口如微信订阅消息、阿里云短信、邮件发送等接口的耗时比较高，对于需要发送大量消息的服务来说，如果使用不当会对资源造成严重浪费甚至会影响核心业务逻辑。另一方面，如果每个服务都自己维护一套发送消息的逻辑实在是太麻烦。其中很多坑，比如access_token的维护，接口维护，n多个模板维护，敏感词维护，消息状态记录。如果一家公司n多个服务都维护这么一套逻辑显得很冗余和浪费人力物力。
  
  这两张图可以更加直观的看出我为什么要写这么一套小系统。还是有些优势的。
  
### 普通的发消息方式
  ![](https://github.com/lovesgg/orange_message_service/blob/master/.doc/orange.png)
  
  
#### 橘子消息系统
  ![](https://github.com/lovesgg/orange_message_service/blob/master/.doc/pic.png)
  ![](https://github.com/lovesgg/orange_message_service/blob/master/.doc/orange-new.png)
  
#### 现在开始您的安装之旅吧。

### 依赖
```bash
1.golang
2.iris
3.redis
4.mysql
5.阿里云短信
6.微信订阅消息
7.邮件发送类库
8.mq平台

```

### 安装
```bash
go (>= 1.13.1)

1.git clone https://github.com/lovesgg/orange_message_service.git
2.cd orange_message_service
3.执行命令开启mod支持: export GO111MODULE=on
4.更换代理: export GOPROXY=https://goproxy.cn
5.下载依赖: go mod vendor
6.cp conf/app.json.example conf/app.json
7.加配置信息如redis mysql 端口等
  1.app.json 环境+端口
  2.log.json 日志目录
  3.message.json 消息配置模板 可自定义通道和模板
  4.mysql.json mysql配置信息 需要自行定义数据库和表。根据自己需要的字段创建和对应代码中的字段即可。可直接使用.doc目录下的message.sql文件去建表
  5.redis.json redis配置信息
  6.sms.json 阿里云短信配置
  7.wechat.json 微信订阅消息的appid配置
  
  (如果需要追加.json文件，请在app/components/config/config.go中修改)
7.配置完信息可以拷贝配置文件到另一个目录啦
  mkdir /data/www/orange_message_service/conf
  cp -r orange_message_service/conf/* /data/www/orange_message_service/conf
  
  (这步骤是因为当您使用 rizla main.go时读的是这位置的配置)
8.创建日志目录
  mkdir /data/logs/orange_message_service
9.mq订阅发布
  请自行安装当前流行的mq平台。
  /client/send 负责发送mq消息
  /server/send 负责消费消息mq
  
  (备注:/client/send 接口里边最后有说明 //执行mq发送 由server端来消费 这部分由您根据实际需要自行添加。如有问题可微信联系。)
  
10.到这里可以认为您的环境已经没问题
  到orange_message_service根目录下执行 rizla main.go即可运行啦。当然您也可以go run main.go
  如果运行报错请先自行排查环境是否都已经安装完毕，或者先自行百度。不清楚的可直接微信联系。
  
```

### 参数可配置说明
```bash
  推荐message.json配置：可自定义
  {
    "1000":{
      "sequence":["subscribe"],
      "subscribe":{
        "channel":"微信订阅消息",
        "template":"SendDataTest",
        "is_retry":0
      },
      "sms":{
        "channel":"阿里云",
        "template":"SendDataTest",
        "is_retry":0
      }
    },
    "1001":{
      "sequence":["email"],
      "email":{
        "channel":"邮件",
        "template":"SendDataTest",
        "is_retry":0
      }
    }
  }
  
  1.key值唯一，上游服务根据这唯一通道+模板参数组成。一次调用只需要关注这id和必要参数即可。
  2.sequence:通道数组。循环发送。比如["sms","email","subscribe"] 此消息依次发送短信、邮件、订阅消息。
  3.email|sms|subscribe 里的参数:
    template对应的是一个方法，每个key对应的次方法唯一。
    is_retry:如果发送失败是否需要重试。
  详情参考以下接口调用。
```


### 项目结构
```bash
├── README.md
├── app
│   ├── common
│   │   ├── aliyun.go
│   │   ├── constant.go
│   │   ├── enum
│   │   │   ├── event_name.go
│   │   │   ├── redis_expire.go
│   │   │   ├── redis_key.go
│   │   │   └── topic.go
│   │   ├── error.go
│   │   ├── structure.go
│   │   ├── templates.go
│   │   └── wechat.go
│   ├── components
│   │   ├── config
│   │   │   └── config.go
│   │   ├── eventdispatcher
│   │   │   ├── event.go
│   │   │   ├── eventdispatcher.go
│   │   │   ├── eventdispatcher_test.go
│   │   │   ├── listener.go
│   │   │   └── log.go
│   │   ├── go-email
│   │   │   └── go_email.go
│   │   ├── http
│   │   │   ├── client.go
│   │   │   └── request.go
│   │   ├── init.go
│   │   ├── mlog
│   │   │   └── log.go
│   │   ├── mysql
│   │   │   └── client.go
│   │   ├── redis
│   │   │   ├── client.go
│   │   │   ├── client_test.go
│   │   │   ├── hash.go
│   │   │   ├── hash_test.go
│   │   │   ├── string.go
│   │   │   └── string_test.go
│   │   ├── rpc
│   │   │   ├── base_rpc.go
│   │   │   └── config.go
│   │   ├── sms
│   │   │   ├── aliyun.go
│   │   │   └── aliyun_test.go
│   │   └── test
│   │       └── helper.go
│   ├── console
│   │   ├── commands
│   │   │   ├── hello.go
│   │   │   ├── init.go
│   │   │   └── server_send.go
│   │   └── root.go
│   ├── controllers
│   │   ├── base_controller.go
│   │   ├── client
│   │   │   ├── send_action.go
│   │   │   └── send_controller.go
│   │   └── server
│   │       ├── send_action.go
│   │       └── send_controller.go
│   ├── events
│   │   └── rpc_call_warning.go
│   ├── middleware
│   │   ├── recover_panic.go
│   │   ├── request_bootstrap.go
│   │   └── request_logger.go
│   ├── models
│   │   ├── message.go
│   │   └── request
│   │       └── request.go
│   ├── render
│   │   └── sms.go
│   ├── router
│   │   └── routes.go
│   ├── services
│   │   ├── client
│   │   │   └── client.go
│   │   ├── db.go
│   │   ├── repository
│   │   ├── rpc
│   │   ├── server
│   │   │   ├── email.go
│   │   │   ├── emailTemplates
│   │   │   │   └── email.go
│   │   │   ├── sms.go
│   │   │   ├── smsTemplates
│   │   │   │   └── sms.go
│   │   │   ├── subscribe.go
│   │   │   ├── subscribeTemplates
│   │   │   │   └── subscribe.go
│   │   │   └── word.go
│   │   └── wechat.go
│   └── utils
│       ├── compress
│       │   └── zlib.go
│       ├── helper.go
│       └── typeconv
│           └── type_trans.go
├── bin
│   ├── orange_message_service.supervisor.conf
│   ├── pre-commit
│   ├── start.sh
│   ├── stop.sh
│   └── supervisord.conf
├── build.sh
├── cmd
│   └── main.go
├── conf
│   ├── app.json
│   ├── app.json.example
│   ├── common.json
│   ├── email.json
│   ├── filter_words.json
│   ├── log.json
│   ├── message.json
│   ├── mysql.json
│   ├── redis.json
│   ├── rpc.json
│   ├── sms.json
│   └── wechat.json
├── docker-compose.yml
├── dockerfile
├── go.mod
├── go.sum
├── main.go
├── nursery
├── orange_message_service
├── output
│   └── server
└── runtime


说明:
1.入口main文件:main.go
2.可执行文件:orange_message_service
3.mod管理文件:go.mod/go.sum
4.路由文件:app/router/routers.go
5.项目配置文件:conf/
6.控制器路径:app/controllers
7.公共常量枚举:app/common
8.公共组件:app/components
9.启动文件:build.sh  (rizla main.go 也行)

```

### 接口使用
 序号  |  接口  | 入参   |  备注
 ---- | ----- | ------  | -------
 1  | /health/check | 无  | 返回正常信息即可验证服务正常启动 
 2  | /client/send | 参考以下 |   客户端接收
 3  | /server/send | 参考以下 |   服务端消费发送
 4  | /client/send-batch | 参考以下 |   批量送
 5  | /client/send-by-sync | 参考以下 |   协程批量发送(推荐使用这方法)
 6  | /client/send-by-users | 参考以下 |   协程批量发送-批量用户发送同一模板消息(推荐使用这方法)

```bash
  1. /client/send 
    {
      "msg_key":1000,
      "source_id":1,
      "body":[{
        "goods_name":"苹果",
        "store_name":"wg",
        "address_detail":"wgrg",
        "phone":"1881000000",
        "note":"你好",
        "user_id":"",
        "order_no":"543646"
      }]
    }
    
  2. /server/send
    {
      "msg_key":1000,
      "source_id":1,
      "body":{
        "goods_name":"苹果",
        "store_name":"wg",
        "address_detail":"wgrg",
        "phone":"1881000000",
        "note":"你好",
        "user_id":"",
        "order_no":"543646"
      }
    }
    
  3./client/send-by-users
    {
      "msg_key":1003,
      "source_id":1,
      "users":[
        "99328515@qq.com",
        "773668125@qq.com"
      ]
    }
```
  


### 消息系统概览图
```bash
  包含一些说明图片
  见orange_message_service/.doc 目录下的图片
```

### 温馨提示
```bash
  此服务适合有经验的人。因为涉及一些第三方接口和几个环境的部署。建议用在发送消息比较多的场景。如果只是发很简单的消息不推荐用这个。
  其中可能包含部分不足的地方。陆续优化中。如果用在生产环境，需要掌握熟悉了才部署。
  欢迎好友来一起交流探讨。
```


### 关于我
```bash
  微信: jdk010 克里斯苏
  需要请喝奶茶可直接加
```
