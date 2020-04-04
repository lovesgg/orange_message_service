package services

import (
	"encoding/json"
	"fmt"
	"orange_message_service/app/components/mysql"
	models "orange_message_service/app/models/request"
	"time"
)

/**
备注 需要您自己新建一张数据表如message
存储的字段由您根据实际参数决定 建议保留的字段:user_id,c_t,u_t,return_data,send_status
保留信息是为了后续排查发送状态等
您也可以忽略这过程 除非对于发送的 结果不care
 */


/**
放入mysql
*/
func InsertDataMysql(params models.ServerReq, toUser string, sendStatus int) bool {
	fmt.Println("insert mysql")

	//return true //如果不需要保存信息到mysql 请修改这方法。直接返回true即可

	db := mysql.GetDb()
	defer db.Close()

	jsonBody, _ := json.Marshal(params.Body)
	create := time.Now().Unix()

	ret, err := db.Exec("INSERT INTO message (to_user, body, c_t, u_t, send_status, source_id) VALUES (?, ?, ?, ?, ?, ?)", toUser, jsonBody, create, create,sendStatus,params.SourceId)
	if err != nil {
		fmt.Println(ret, err)
		return false
	}
	fmt.Println("insert db ok")
	return true

}

/**
放入mongodb
*/
func InsertDataMongodb() {

}

/**
放入oracle
*/
func InsertDataOracle() {

}
