package redis

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/magiconair/properties/assert"
	config2 "orange_message_service/app/components/config"
	"testing"
	"time"
)

func init() {
	config2.Init()
	Init()
}

func TestMGet(t *testing.T) {
	redisClient := GetProductClient()
	values, err := redisClient.MGet([]string{"name", "name2", "age"})

	spew.Dump(values, err)
}

func TestIncrBy(t *testing.T) {
	redisClient := GetProductClient()
	err := redisClient.Set("age", 1000, 100*time.Second)
	fmt.Println(err)
	fmt.Println(redisClient.GetInt("age"))
	fmt.Println(redisClient.IncrBy("age", 1, 10*time.Second))
	fmt.Println(redisClient.TTL("age"))
	fmt.Println(redisClient.GetInt("age"))
	fmt.Println(redisClient.Del("age"))
}

func TestTTL(t *testing.T) {
	redisClient := GetProductClient()
	_ = redisClient.Set("age", 1000, 100*time.Second)
	d, _ := redisClient.TTL("age")
	assert.Equal(t, 100*time.Second, d)
}
