package mongo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
这里不做具体公共方法编写 如需要请自行实现
mongo配置信息可以写在config目录下再获取不用这么写死
*/

func MongoClient(table string) (*mongo.Collection, *mongo.Client) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		fmt.Println("connect error")
	}
	// 获取数据库和集合
	collection := client.Database("test").Collection(table)
	return collection, client
}

/**
插入数据
*/
func Insert(datas map[string]interface{}) {

}

/**
删除
*/
func Delete(datas map[string]interface{}) {

}

/**
更新数据
*/
func Update(datas map[string]interface{}) {

}

/**
查询单个
*/
func SelectOne(datas map[string]interface{}) {

}

/**
批量查询
*/
func Select(datas map[string]interface{}) {

}
