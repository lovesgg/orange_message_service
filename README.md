# Orange 消息推送系统
  Orange是由golang、iris、mysql、redis技术框架组成的消息推送服务。
  支持微信订阅消息、短信消息、邮件消息的统一可配置管理发送。支持不同策略的发送。有效管理模板资源，敏感词过滤，自定义追加配置和模板策略，记录消息体，可以为多个业务方向提供消息发送服务。测试后可线上使用。
  
  为什么需要这样的服务？
  当公司的业务变多，服务不断增加，需要为用户发送的消息越来越多。需要为用户发送譬如下单通知、到货通知、评价消息、接单成功通知、注册消息、其他推广信息。而众所周知利用第三方发送接口如微信订阅消息、阿里云短信、邮件发送等接口的耗时比较高，对于需要发送大量消息的服务来说，如果使用不当会对资源造成严重浪费甚至会影响核心业务逻辑。另一方面，如果每个服务都自己维护一套发送消息的逻辑实在是太麻烦。其中很多坑，比如access_token的维护，接口维护，n多个模板维护，敏感词维护，消息状态记录。如果一家公司n多个服务都维护这么一套逻辑显得很冗余和浪费人力物力。
  
  #使用过程中有问题请联系微信:jdk010

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
9.

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
  1.
  2.
  3.
  4.
  5.
  6.
7.配置完信息可以拷贝配置文件到另一个目录啦
  mkdir /data/www/orange_message_service/conf
  cp -r orange_message_service/conf/* /data/www/orange_message_service/conf
8.创建日志目录
  mkdir /data/logs/orange_message_service
  
  
  
```
  

### 消息系统概览图





