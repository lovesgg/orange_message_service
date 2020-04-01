package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"orange_message_service/app/components/config"
	"orange_message_service/app/components/mlog"
	"time"
)

type Client struct {
	baseClient *redis.Client
	prefix     string
}

var clients map[string]*Client

var REDIS_NAMES = []string{
	"notice_redis",
}

func GetCommonClient() *Client {
	return getClient("notice_redis")
}

func getClient(redisName string) *Client {
	return clients[redisName]
}

func getOpt(redisName string) *redis.FailoverOptions {
	config := config.GetConfig()

	maxRetries := config.GetInt(redisName + ".max_retries")
	if maxRetries < 0 || maxRetries > 3 {
		maxRetries = 0
	}
	return &redis.FailoverOptions{
		MasterName:         config.GetString(redisName + ".master_name"),
		SentinelAddrs:      config.GetStringSlice(redisName + ".addrs"),
		DialTimeout:        time.Duration(config.GetInt(redisName+".dial_timeout")) * time.Second,  //建立 redis conn 超时时间
		ReadTimeout:        time.Duration(config.GetInt(redisName+".read_timeout")) * time.Second,  //写入数据超时时间
		WriteTimeout:       time.Duration(config.GetInt(redisName+".write_timeout")) * time.Second, //读取数据超时时间
		PoolTimeout:        time.Duration(config.GetInt(redisName+".pool_timeout")) * time.Second,  //当连接池所有的连接都被使用，client等待的超时时间
		MaxRetries:         maxRetries,                                                             //redis command执行重试次数
		PoolSize:           config.GetInt(redisName + ".pool_size"),                                //连接池大小，默认10*core
		IdleTimeout:        5 * time.Minute,                                                        //空闲连接时间，超时后会释放
		IdleCheckFrequency: 5 * time.Minute,                                                        //5分钟check一次连接池
		//下面两个参数暂时不起作用
		MinIdleConns: config.GetInt(redisName + ".min_idle_conns"), //连接池最小空闲连接数
		MaxConnAge:   1 * time.Hour,                                //连接从创建开始，最大的存活时间
	}
}

func Init() {
	clients = make(map[string]*Client)
	redis.SetLogger(zap.NewStdLog(mlog.GetLogger()))
	for _, redisName := range REDIS_NAMES {
		// 使用本地 redis 进行操作
		//baseClient := redis.NewClient(&redis.Options{
		//	Addr:               "127.0.0.1:6379",
		//	Password:           "root",
		//})
		baseClient := redis.NewFailoverClient(getOpt(redisName))
		baseClient.Options().OnConnect = func(conn *redis.Conn) error {
			mlog.ZapInfo("new redis client connection." + fmt.Sprint(&conn))
			return nil
		}

		baseClient.Ping()
		client := &Client{
			baseClient: baseClient,
			prefix:     config.GetString(redisName + ".prefix"),
		}
		clients[redisName] = client
	}
}

// redis key 不存在的时候会返回 redis.Nil
func (Client) ErrIsNil(err error) bool {
	return err == redis.Nil
}

// Expire 设置过期时间
func (c Client) Expire(key string, expiration time.Duration) error {
	return c.baseClient.Expire(c.CacheKey(key), expiration).Err()
}

// Delete keys.
func (c Client) Del(keys ...string) (int64, error) {
	for i, key := range keys {
		keys[i] = c.CacheKey(key)
	}
	return c.baseClient.Del(keys...).Result()
}

func (c Client) TTL(key string) (time.Duration, error) {
	return c.baseClient.TTL(c.CacheKey(key)).Result()
}

func (c Client) Pipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return c.baseClient.Pipelined(fn)
}

func (c Client) GetBaseClient() *redis.Client {
	return c.baseClient
}

func (c Client) CacheKey(key string) string {
	return c.prefix + key
}

func (c Client) Exists(key string) (int64, error) {
	return c.baseClient.Exists(c.CacheKey(key)).Result()
}


// 记录日志，但是这种方式不能获取到上下文
//func wrapDebugLogger(client *redis.Client) {
//	client.WrapProcess(func(oldProcess func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
//		return func(cmd redis.Cmder) error {
//			startTime := time.Now().UnixNano() / 1e6
//			err := oldProcess(cmd)
//			if err != nil {
//				return err
//			}
//			endTime := time.Now().UnixNano() / 1e6
//
//			mlog.GetLogger().Info(fmt.Sprintf("redis: executed.||cost_time=%dms||cmd=%s", endTime-startTime, spew.Sdump(cmd)))
//
//			return nil
//		}
//	})
//}
