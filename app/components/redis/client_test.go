package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

func TestPipeline(t *testing.T) {
	redisClient := GetProductClient()
	baseClient := redisClient.GetBaseClient()
	baseClient.Set("name1", "name1", 1000*time.Second)
	baseClient.Set("name2", "name2", 1000*time.Second)
	baseClient.Set("name3", "name3", 1000*time.Second)
	baseClient.Set("name4", "name4", 1000*time.Second)

	var get1 *redis.StringCmd
	var get2 *redis.StringCmd
	var get3 *redis.StringCmd
	var get4 *redis.StringCmd
	_, _ = baseClient.Pipelined(func(pipe redis.Pipeliner) error {
		get1 = pipe.Get("name1")
		get2 = pipe.Get("name2")
		get3 = pipe.Get("name3")
		get4 = pipe.Get("name4")
		return nil
	})

	fmt.Println(get1.Val(), get2.Val(), get3.Val(), get4.Val())
}
