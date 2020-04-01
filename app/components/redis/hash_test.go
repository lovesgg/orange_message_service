package redis

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"orange_message_service/app/utils/compress"
	"testing"
)

type TestStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (t TestStruct) MarshalBinary() (data []byte, err error) {
	return jsoniter.Marshal(t)
}

func TestHash(t *testing.T) {
	t1 := TestStruct{
		Name: "lirui",
		Age:  25,
	}
	redisClient := GetProductClient()
	fmt.Println(redisClient.HSet("stus", "t1", t1))
	_, _ = redisClient.Del("stus")
}

func TestGzipCompress(t *testing.T) {
	t1 := TestStruct{
		Name: "lirui",
		Age:  25,
	}
	bts, _ := t1.MarshalBinary()
	redisClient := GetProductClient()
	cps, _ := compress.ZlibCompress(bts)
	fmt.Println(redisClient.HSet("stus", "t1", cps))
	fmt.Println(redisClient.HGetAll("stus"))
	_, _ = redisClient.Del("stus")
}
